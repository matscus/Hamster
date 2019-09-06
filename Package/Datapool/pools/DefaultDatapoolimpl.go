package pools

import (
	"database/sql"
	"errors"
	"math/rand"
	"os"

	//_ mask PG driver
	_ "github.com/lib/pq"

	"github.com/matscus/Hamster/Package/Datapool/structs"
)

var (
	selectDefaultPool = "select Surname,Name,Patronymic,DocumentsType,DocumentFullValue,Birthdate,Phone,Mail,Snils,Inn,AccountNumber,ContractNumber,CardNumber,SystemCode,ContractID,SystemID,RawID,Hid,Pan from defaultdatapool"
)

func (p Datapool) GetDefaultDatapool(lenpool int, c chan<- structs.DefaultDatapool) error {
	datapool := make([]structs.DefaultDatapool, 0, lenpool)
	db, err := sql.Open("postgres", "user="+os.Getenv("POSTGRESUSER")+" password="+os.Getenv("POSTGRESPASSWORD")+" dbname="+os.Getenv("POSTGRESDB")+" sslmode=disable")
	if err != nil {
		err = errors.New(err.Error())
		return err
	}
	defer db.Close()
	stmt, err := db.Prepare(selectDefaultPool)
	if err != nil {
		return err
	}
	defer stmt.Close()
	var rows *sql.Rows
	rows, err = stmt.Query()
	if err != nil {
		return err
	}
	i := 0
	for rows.Next() {
		if i < lenpool {
			pool := structs.DefaultDatapool{}
			rows.Scan(&pool.Surname, &pool.Name, &pool.Patronymic, &pool.DocumentsType, &pool.DocumentFullValue, &pool.Birthdate, &pool.Phone, &pool.Mail, &pool.Snils, &pool.Inn, &pool.AccountNumber, &pool.ContractNumber, &pool.CardNumber, &pool.SystemCode, &pool.ContractID, &pool.SystemID, &pool.RawID, &pool.Hid, &pool.Pan)
			datapool = append(datapool, pool)
			i++
		} else {
			break
		}
	}
	rows.Close()
	for {
		p := rand.Intn(lenpool)
		c <- datapool[p]
	}
}
