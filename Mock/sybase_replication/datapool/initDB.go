package datapool

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"

	//_ driver for tds
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	_ "github.com/thda/tds"
)

var (
	// 	//imdg
	// 	//Xerjdcrbq10
	//DB - pointer for DB connect pool
	DB             *sql.DB
	Cnfg           Config
	JsonPool       []Operation
	ConnectionPool sync.Map
	//TempIDSyncMap  sync.Map
	metrics = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "sybase_operation_count",
			Help: "operation count",
		},
		[]string{"node", "table", "operation_type"},
	)
	errmetrics = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "sybase_operation_error_count",
			Help: "operation error count",
		},
		[]string{"node", "table", "operation_type", "error"},
	)
)

type Instance struct {
	Host       string
	DB         *sql.DB
	Duration   time.Duration
	C          chan TypeOrSTR
	Operations []Operation
}

type Config []struct {
	Host   string `json:"Host"`
	Port   string `json:"Port"`
	DBName string `json:"DBName"`
	Tables []struct {
		TableName   string  `json:"TableName"`
		InsertCount float64 `json:"InsertCount"`
		UpdateCount float64 `json:"UpdateCount"`
	} `json:"Tables"`
}
type Operation struct {
	Tablename string `json:"tablename"`
	Insert    string `json:"insert"`
	Update    []struct {
		Str string `json:"str"`
	} `json:"update"`
	Delete string `json:"delete"`
}

//TypeOrSTR - temp struct for exec to database
type TypeOrSTR struct {
	TableName     string
	OperationType string
	STR           string
}

func init() {
	err := readConfig()
	if err != nil {
		log.Panicf("[Panic] read config errors: %s", err)
	}
	err = readJSON()
	if err != nil {
		log.Panicf("[Panic] read json files errors: %s", err)
	}
	l := len(Cnfg)
	for i := 0; i < l; i++ {
		cnxStr := "tds://sa:password@" + Cnfg[i].Host + ":" + Cnfg[i].Port + "/" + Cnfg[i].DBName + "?charset=utf8"
		db, err := sql.Open("tds", cnxStr)
		if err != nil {
			log.Fatalln("Init connection error ", err)
		}
		ConnectionPool.Store(Cnfg[i].Host, db)
	}
	prometheus.MustRegister(metrics)
	prometheus.MustRegister(errmetrics)
	http.Handle("/metrics", promhttp.Handler())
	log.Printf("[INFO] " + " Listenport: 9990")
	go http.ListenAndServe(":9990", nil)
}

func (instance Instance) RunInstance(wg *sync.WaitGroup) {
	var tempIDSyncMap sync.Map
	lastIDSlice, err := GetLastIDFromTable(instance.DB)
	if err != nil {
		log.Printf("[ERROR] getlastid error: %s", err)
	}
	var localwg sync.WaitGroup
	localwg.Add(1)
	go Writer(instance.Host, instance.DB, instance.C, instance.Duration, &localwg, &tempIDSyncMap)
	l := len(instance.Operations)
	for i := 0; i < l; i++ {
		tablename := instance.Operations[i].Tablename
		for _, val := range Cnfg {
			l := len(val.Tables)
			for i := 0; i < l; i++ {
				if val.Tables[i].TableName == tablename {
					localwg.Add(1)
					go instance.Operations[i].Reader(instance.C, val.Tables[i].InsertCount, val.Tables[i].UpdateCount, instance.Duration, &localwg)
					break
				}
			}
		}
	}
	localwg.Wait()
	clearDataBaseAfterTest(instance.DB, &lastIDSlice)
	wg.Done()
}

func Writer(host string, db *sql.DB, c chan TypeOrSTR, d time.Duration, wg *sync.WaitGroup, tempIDSyncMap *sync.Map) {
	start := time.Now()
	var typeOrSTR TypeOrSTR
	for {
		typeOrSTR = <-c
		switch typeOrSTR.OperationType {
		case "insert":
			i, err := db.Exec(typeOrSTR.STR)
			if err != nil {
				log.Println("Exec insert: ", err)
				errmetrics.With(prometheus.Labels{"node": host, "table": typeOrSTR.TableName, "operation_type": "Insert", "error": err.Error()}).Inc()
			} else {
				metrics.With(prometheus.Labels{"node": host, "table": typeOrSTR.TableName, "operation_type": "Insert"}).Inc()
				id, err := i.LastInsertId()
				if err != nil {
					log.Println("Exec insert: ", err)
				} else {
					pt, ok := tempIDSyncMap.Load(typeOrSTR.TableName)
					if ok {
						temp := append(pt.([]int64), id)
						tempIDSyncMap.Store(typeOrSTR.TableName, temp)
					} else {
						tempIDSyncMap.Store(typeOrSTR.TableName, []int64{id})
					}
				}
			}
		case "update":
			pt, ok := tempIDSyncMap.Load(typeOrSTR.TableName)
			if ok {
				l := len(pt.([]int64))
				updates := pt.([]int64)
				_, err := db.Exec(typeOrSTR.STR, updates[rand.Intn(l)])
				if err != nil {
					errmetrics.With(prometheus.Labels{"node": host, "table": typeOrSTR.TableName, "operation_type": "Update", "error": err.Error()}).Inc()

				} else {
					metrics.With(prometheus.Labels{"node": host, "table": typeOrSTR.TableName, "operation_type": "Update"}).Inc()
				}
			}
		case "delete":
			pt, ok := tempIDSyncMap.Load(typeOrSTR.TableName)
			if ok {
				deletes := pt.([]int64)
				randID := rand.Intn(len(pt.([]int64)))
				id := deletes[randID]
				_, err := db.Exec(typeOrSTR.STR, id)
				if err != nil {
					errmetrics.With(prometheus.Labels{"node": host, "table": typeOrSTR.TableName, "operation_type": "Delete", "error": err.Error()}).Inc()

				} else {
					metrics.With(prometheus.Labels{"node": host, "table": typeOrSTR.TableName, "operation_type": "Delete"}).Inc()
					tempIDSyncMap.Store(typeOrSTR.TableName, append(deletes[:randID], deletes[randID+1:]...))
				}
			}
		}
		if time.Now().Sub(start) >= d {
			wg.Done()
			break
		}
	}
}

func GetLastIDFromTable(db *sql.DB) (map[string]uint64, error) {
	lastIDSlice := make(map[string]uint64)
	l := len(JsonPool)
	for i := 0; i < l; i++ {
		var id uint64
		rows, err := db.Query("select max(" + getNameTableID(JsonPool[i].Tablename) + ") from " + JsonPool[i].Tablename)
		if err != nil {
			return nil, err
		}
		for rows.Next() {
			rows.Scan(&id)
		}
		lastIDSlice[JsonPool[i].Tablename] = id
	}
	return lastIDSlice, nil
}

func (o Operation) Reader(c chan TypeOrSTR, insertCount float64, updateCount float64, d time.Duration, wg *sync.WaitGroup) {
	start := time.Now()
	pacing := time.Duration(math.Round(3600000000/insertCount)) * time.Microsecond
	var rate float64
	if insertCount < updateCount {
		rate = math.Round(updateCount / insertCount)
	} else {
		rate = 1
	}
	insDelTic := time.NewTicker(pacing)
	updateTic := time.NewTicker(pacing / time.Duration(rate))
	deleteflag := true
	for {
		c <- TypeOrSTR{o.Tablename, "insert", o.Insert}
		l := len(o.Update)
		for i := 0; i < int(rate); i++ {
			c <- TypeOrSTR{o.Tablename, "update", o.Update[rand.Intn(l)].Str}
			<-updateTic.C
		}
		if deleteflag {
			c <- TypeOrSTR{o.Tablename, "delete", o.Delete}
			deleteflag = false
		} else {
			deleteflag = true
		}
		if time.Now().Sub(start) >= d {
			wg.Done()
			break
		}
		<-insDelTic.C
	}
}
func readConfig() (err error) {
	file, err := os.Open("./datapool/config.json")
	if err != nil {
		return err
	}
	err = json.NewDecoder(file).Decode(&Cnfg)
	if err != nil {
		return err
	}
	return nil
}
func readJSON() (err error) {
	files, err := ioutil.ReadDir("./json")
	if err != nil {
		return err
	}
	for _, f := range files {
		oper := Operation{}
		file, err := os.Open("./json/" + f.Name())
		if err != nil {
			return err
		}
		err = json.NewDecoder(file).Decode(&oper)
		if err != nil {
			return err
		}
		JsonPool = append(JsonPool, oper)
	}
	return nil
}
func getNameTableID(tablename string) (res string) {
	switch tablename {
	case "tCard":
		res = "CardID"
	case "tCardProduct":
		res = "CardProductID"
	case "tContract":
		res = "ContractID"
	case "tContractCredit":
		res = "ContractCreditID"
	case "tInstitution":
		res = "InstitutionID"
	case "tNode":
		res = "NodeID"
	case "tResource":
		res = "ResourceID"
	case "tSecurity":
		res = "SecurityID"
	}
	return res
}
func remove(slice []int64, s int) []int64 {
	return append(slice[:s], slice[s+1:]...)
}
func clearDataBaseAfterTest(db *sql.DB, lastIDSlice *map[string]uint64) (err error) {
	for key, value := range *lastIDSlice {
		_, err = db.Exec("delete from " + key + " where " + getNameTableID(key) + " > " + strconv.FormatUint(value, 10))
		if err != nil {
			fmt.Printf("[ERROR] delete err %s", err)
			//return nil
		}
	}
	return nil
}
