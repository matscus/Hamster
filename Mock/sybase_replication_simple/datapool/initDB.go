package datapool

import (
	"database/sql"
	"log"

	//_ driver for tds
	_ "github.com/thda/tds"
)

var (
	//DB - pointer for DB connect pool
	DB *sql.DB
)

func init() {
	var err error
	cnxStr := "tds://sa:password@localhost:5000/master?charset=utf8"
	DB, err = sql.Open("tds", cnxStr)
	if err != nil {
		log.Fatalln("Init connection error ", err)
	}
}
