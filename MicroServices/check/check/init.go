package check

import (
	"os"

	"github.com/matscus/Hamster/Package/Clients/client"
	"github.com/matscus/Hamster/Package/Clients/client/postgres"
)

var (
	pgClient *postgres.PGClient
)

func init() {
	pgConf := postgres.Config{Driver: "postgres",
		User:     os.Getenv("POSTGRESUSER"),
		Password: os.Getenv("POSTGRESPASSWORD"),
		DataBase: os.Getenv("POSTGRESDB"),
		SSLMode:  "disable"}
	client := client.New("postgres", pgConf).(postgres.PGClient)
	pgClient = &client
}
