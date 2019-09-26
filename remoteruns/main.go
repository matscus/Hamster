package main

import (
	"encoding/json"
	"flag"
	"log"
	"os"
	"sync"

	"github.com/matscus/Hamster/Package/Clients/client"
)

var str string

type RemoteAddr []struct {
	Host string `json:"Host"`
}

func main() {
	flag.StringVar(&str, "str ", "", "string for run to remote host")
	flag.Parse()
	if str == "" {
		log.Panic("[ERROR] %s", "param str in nil")
	}
	file, err := os.Open("config.json")
	if err != nil {
		log.Printf("[ERROR] %s", err)
	}
	cnfg := RemoteAddr{}
	err = json.NewDecoder(file).Decode(&cnfg)
	if err != nil {
		log.Printf("[ERROR] %s", err)
	}
	l := len(cnfg)
	var wg sync.WaitGroup
	for i := 0; i < l; i++ {
		wg.Add(1)
		s := "nohup " + str + " &> /dev/null"
		go run(cnfg[i].Host, s, &wg)
	}
	wg.Wait()
	log.Printf("[INFO] %s", "run is complined")
}

func run(host string, str string, wg *sync.WaitGroup) error {
	client, err := client.SSHClient{}.New()
	if err != nil {
		return err
	}
	err = client.Run(host, str)
	if err != nil {
		return err
	}
	wg.Done()
	return nil
}
