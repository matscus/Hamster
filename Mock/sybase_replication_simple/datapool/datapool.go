package datapool

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"time"
)

var (
	//Pool - slice for all operations, init at the start application
	Pool []Datapool
)

func init() {
	readJSON()
}

//Datapool - struct for operation
type Datapool struct {
	Name   string `json:"Name"`
	Insert string `json:"Insert"`
	Update []struct {
		Update struct {
			Str string `json:"Str"`
		} `json:"Update"`
	} `json:"Update"`
	Delete string `json:"Delete"`
}

func readJSON() {
	files, err := ioutil.ReadDir("./json")
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range files {
		pool := Datapool{}
		file, err := os.Open("./json/" + f.Name())
		if err != nil {
			log.Fatalln("error read file: ", err)
		}
		err = json.NewDecoder(file).Decode(&pool)
		if err != nil {
			log.Fatalln("error decode file: ", err)
		}
		Pool = append(Pool, pool)
	}
}

//Reader - func for read slice Pool and send str(for exec to db) to chan
func Reader(c chan string) {
	lenpool := len(Pool)
	for {
		for i := 0; i < lenpool; i++ {
			pool := Pool[i]
			lenupdate := len(pool.Update)
			c <- pool.Insert
			for i := 0; i < lenupdate; i++ {
				c <- pool.Update[i].Update.Str
			}
			c <- pool.Delete
		}
	}
}

//Writer - func read chan and exec operation to db
func Writer(c chan string, timeout time.Duration) {
	iter := 0
	for {
		str := <-c
		_, err := DB.Exec(str)
		if err != nil {
			log.Panic("operation fail: ", err)
		}
		iter++
		time.Sleep(timeout)
	}
}
