package main

import (
	"database/sql"
	"flag"
	"log"
	"os"
	"sync"
	"time"

	"github.com/matscus/Hamster/Mock/sybase_replication/datapool"
	//"./datapool"
)

var (
	duration int
)

func main() {
	// defer profile.Start(profile.MemProfile).Stop()
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
	log.Printf("[INFO] mock sybase replication is start")
	flag.IntVar(&duration, "duration ", 5, "duration work")
	flag.Parse()
	var wg sync.WaitGroup
	l := len(datapool.Cnfg)
	for i := 0; i < l; i++ {
		val, _ := datapool.ConnectionPool.Load(datapool.Cnfg[i].Host)
		c := make(chan datapool.TypeOrSTR, 1000)
		wg.Add(1)
		inctance := datapool.Instance{datapool.Cnfg[i].Host, val.(*sql.DB), convertDuration(duration), c, datapool.JsonPool}
		go inctance.RunInstance(&wg)
	}
	wg.Wait()
	log.Printf("[INFO] mock sybase replication is complite")
	os.Exit(0)
}

func convertDuration(i int) (d time.Duration) {
	return time.Duration(i) * time.Minute
}
