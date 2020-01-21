package postgres

import "database/sql"

//PGClient  struct for postgres client
type PGClient struct {
	DB *sql.DB
}

//Config for client postgres DB
type Config struct {
	Driver   string
	User     string
	Password string
	DataBase string
	SSLMode  string
}
