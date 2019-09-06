package main

import (
	"flag"
	"log"
	"os"
	"runtime/trace"
	"sync"
	"time"

	"github.com/pkg/profile"

	"./datapool"
)

var (
	duration int
)

func init() {
	datapool.ReadJSON()
}
func main() {
	defer profile.Start(profile.MemProfile).Stop()
	f, err := os.Create("trace.out")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	err = trace.Start(f)
	if err != nil {
		panic(err)
	}
	defer trace.Stop()
	flag.IntVar(&duration, "duration ", 10, "duration work")
	flag.Parse()
	l := len(datapool.Datapool)
	var wg sync.WaitGroup
	for i := 0; i < l; i++ {
		wg.Add(1)
		if datapool.Datapool[i].InsertDeleteCount == 0 {
			log.Printf("Values insert and delete equally 0, delete operation %s excluded from startup script", datapool.Datapool[i].Tablename)
		} else {
			go datapool.Datapool[i].Run(convertDuration(duration), &wg)
		}

	}
	wg.Wait()
	os.Exit(0)
}

func convertDuration(i int) (d time.Duration) {
	return time.Duration(i) * time.Minute
}
