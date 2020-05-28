package handlers

import (
	"time"

	"github.com/matscus/Hamster/Package/Clients/client/postgres"

	"github.com/matscus/Hamster/Package/ScriptCache/scriptcache"
)

var (
	cache    = scriptcache.New(1*time.Minute, 5*time.Minute)
	PgClient *postgres.PGClient
)
