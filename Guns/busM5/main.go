package main

import (
	"flag"
	"log"
	"strconv"
	"time"

	"github.com/matscus/Hamster/Package/Clients/client"

	"./asserts"
	"./mqops"
	"./pool"
)

var (
	duration                                                  int
	runID, testName, testType, scenario, datapool, csv, dburl string
)

func main() {
	//defer profile.Start().Stop()
	//defer profile.Start(profile.MemProfile).Stop()
	// f, err := os.Create("trace.out")
	// if err != nil {
	// 	panic(err)
	// }
	// defer f.Close()

	// err = trace.Start(f)
	// if err != nil {
	// 	panic(err)
	// }
	// defer trace.Stop()
	flag.IntVar(&duration, "duration", 5, "duration test")
	flag.StringVar(&scenario, "scenario", "mqconfig.json", "set scenario")
	flag.StringVar(&datapool, "datapool", "db", "set datapool")
	flag.StringVar(&csv, "csv", "users.csv", "name csv file")
	flag.StringVar(&testName, "testName", "superPuperTest", "name test")
	flag.StringVar(&testType, "testType", "mq_stage_2", "type test")
	flag.Parse()
	go mqops.InitConfigs("./scenario/" + scenario)
	pgclient := client.PGClient{}.New()
	runID, err := pgclient.GetNewRunID()
	if err != nil {
		log.Println("Get New run id: ", err)
	}
	if runID == 0 {
		runID = 1
	}
	switch datapool {
	case "db":
		go pool.InitPoolChFromDB()
	case "csv":
		go pool.InitPoolChFromCSV("./pool/" + csv)
	}
	time.Sleep(5 * time.Second)
	err = pgclient.SetStartTest(strconv.FormatInt(runID, 10), testName, testType)
	if err != nil {
		log.Println("error set start test", err)
	}
	lops := len(mqops.MQConf.Data.Operation)
	lmqops := len(mqops.MQOps)
	for i := 0; i < lops; i++ {
		opername := mqops.MQConf.Data.Operation[i].Data.Name
		for i := 0; i < lmqops; i++ {
			if opername == mqops.MQOps[i].Name {
				var action mqops.Action
				action = mqops.MQOps[i]
				action.InitMQOps()
				go action.RunConsumer()
				go action.RunProduser()
				break
			}
		}
	}
	time.Sleep(convertDuration(duration) * time.Minute)
	pgclient.SetStopTest(strconv.FormatInt(runID, 10))
	res := asserts.CheckTestResult()
	if !res {
		log.Println("Test fail")
	} else {
		log.Println("Test pass")
	}
}
func convertDuration(i int) (d time.Duration) {
	return time.Duration(i) * time.Nanosecond
}
