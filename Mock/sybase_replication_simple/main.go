package main

import (
	"flag"
	"os"
	"time"

	"./datapool"
)

var (
	duration int
	tps      int
)

func main() {
	flag.IntVar(&duration, "duration ", 1, "duration work")
	flag.IntVar(&tps, "tps", 10, "transaction per second")
	flag.Parse()
	c := make(chan string, 10000)
	go datapool.Reader(c)
	time.Sleep(1 * time.Second)
	go datapool.Writer(c, convertTPS(tps))
	func() {
		time.Sleep(convertDuration(duration))
		os.Exit(0)
	}()
}

func convertDuration(i int) (d time.Duration) {
	return time.Duration(i) * time.Minute
}
func convertTPS(tps int) (d time.Duration) {
	return time.Duration(1000/tps) * time.Millisecond
}
