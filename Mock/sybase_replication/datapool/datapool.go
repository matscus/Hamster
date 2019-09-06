package datapool

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"math"
	"math/rand"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	Datapool []Operation
	metrics  = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "sybase_operation_count",
			Help: "operation count",
		},
		[]string{"table", "operation_type"},
	)
	errmetrics = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "sybase_operation_error_count",
			Help: "operation error count",
		},
		[]string{"table", "operation_type", "error"},
	)
)

func init() {
	prometheus.MustRegister(metrics)
	prometheus.MustRegister(errmetrics)
	http.Handle("/metrics", promhttp.Handler())
	log.Println("Listenport: 9990")
	go http.ListenAndServe(":9990", nil)
}

type Operation struct {
	Tablename         string  `json:"Tablename"`
	InsertDeleteCount float64 `json:"InsertDeleteCount"`
	UpdateCount       float64 `json:"UpdateCount"`
	Insert            string  `json:"Insert"`
	Update            []struct {
		Str string `json:"Str"`
	} `json:"Update"`
	Delete string `json:"Delete"`
}

func (o Operation) Run(d time.Duration, wg *sync.WaitGroup) {
	start := time.Now()
	pacing := time.Duration(math.Round(3600000000/o.InsertDeleteCount)) * time.Microsecond
	var rate float64
	if o.InsertDeleteCount < o.UpdateCount {
		rate = math.Round(o.UpdateCount / o.InsertDeleteCount)
	} else {
		rate = 1
	}
	insDelTic := time.NewTicker(pacing)
	updateTic := time.NewTicker(pacing / time.Duration(rate))
	for {
		i, err := DB.Exec(o.Insert)
		if err != nil {
			log.Println("Exec insert: ", err)
			errmetrics.With(prometheus.Labels{"table": o.Tablename, "operation_type": "Insert", "error": err.Error()}).Inc()
			wg.Done()
			break
		} else {
			metrics.With(prometheus.Labels{"table": o.Tablename, "operation_type": "Insert"}).Inc()
			id, err := i.LastInsertId()
			if err != nil {
				log.Println("Get Last ID err of table ", o.Tablename, err)
				wg.Done()
				break
			}
			l := len(o.Update)
			for i := 0; i < int(rate); i++ {
				_, err = DB.Exec(o.Update[rand.Intn(l)].Str, id)
				if err != nil {
					metrics.With(prometheus.Labels{"table": o.Tablename, "operation_type": "Update"}).Inc()
				} else {
					log.Println("Exec update: ", err)
					errmetrics.With(prometheus.Labels{"table": o.Tablename, "operation_type": "Update", "error": err.Error()}).Inc()
				}
				<-updateTic.C
			}
			_, err = DB.Exec(o.Delete, id)
			if err != nil {
				metrics.With(prometheus.Labels{"table": o.Tablename, "operation_type": "Delete"}).Inc()
			} else {
				log.Println("Exec Delete: ", err)
				errmetrics.With(prometheus.Labels{"table": o.Tablename, "operation_type": "Delete", "error": err.Error()}).Inc()
			}
			if time.Now().Sub(start) >= d {
				wg.Done()
				break
			}
		}
		<-insDelTic.C
	}
}

func ReadJSON() {
	files, err := ioutil.ReadDir("./json")
	checkError("Read dir: ", err)
	for _, f := range files {
		oper := Operation{}
		file, err := os.Open("./json/" + f.Name())
		checkError("Read file: ", err)
		err = json.NewDecoder(file).Decode(&oper)
		checkError("Decode file: ", err)
		Datapool = append(Datapool, oper)
	}
}

func checkError(str string, err error) {
	if err != nil {
		log.Println(str, err)
	}
}
