package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", "user="+os.Getenv("POSTGRESUSER")+" password="+os.Getenv("POSTGRESPASSWORD")+" dbname="+os.Getenv("POSTGRESDB")+" sslmode=disable")
	if err != nil {
		log.Println(err)
	}

}
