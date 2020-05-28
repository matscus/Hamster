package client

import (
	"database/sql"

	"github.com/matscus/Hamster/Package/Clients/client/postgres"
	"github.com/matscus/Hamster/Package/Clients/client/sshimpl"
	"github.com/matscus/Hamster/Package/Clients/subset/pgclient"
	"golang.org/x/crypto/ssh"
)

//PGClient  struct for postgres client
type PGClient struct {
	DB *sql.DB
}

//New funct return client
func New(clientType string, config interface{}) interface{} {
	switch clientType {
	case "postgres":
		c := config.(postgres.Config)
		db, err := sql.Open(c.Driver, "user="+c.User+" password="+c.Password+" dbname="+c.DataBase+" sslmode="+c.SSLMode)
		if err != nil {
			return err
		}
		var client pgclient.PGClient
		client = postgres.PGClient{DB: db}
		return client
	case "ssh":
		client := sshimpl.SSHClient{
			SHHConfig: config.(*ssh.ClientConfig)}
		return client
	}
	return nil
}
