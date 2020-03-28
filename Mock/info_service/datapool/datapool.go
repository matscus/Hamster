package datapool

import (
	"bufio"
	"encoding/csv"
	"io"
	"log"
	"os"
)

var (
	FilePath string
	GUIDPool = make(map[string]Datapool)
)

type Datapool struct {
	GUID        string
	UserID      string
	DealID      string
	AccNum      string
	Phone       string
	ContractNum string
}

func IntitDataPool() error {
	log.Println("init start")
	file, err := os.Open(FilePath)
	if err != nil {
		return err
	}
	defer file.Close()
	reader := csv.NewReader(bufio.NewReader(file))
	for {
		r, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		GUIDPool[r[0]] = Datapool{GUID: r[0], UserID: r[1], DealID: r[2], AccNum: r[3], Phone: r[4], ContractNum: r[5]}
	}
	log.Println("init done")
	return nil
}
