package handlers

import (
	"time"

	"github.com/matscus/Hamster/Package/ScriptCache/scriptcache"
)

var (
	cache = scriptcache.New(1*time.Minute, 5*time.Minute)
)
