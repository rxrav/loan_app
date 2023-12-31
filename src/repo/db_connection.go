package repo

import (
	"fmt"
	"sync"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog/log"
)

var db *sqlx.DB
var err error
var dbInitLock = &sync.Mutex{}

func GetDbInstance(connStr string) *sqlx.DB {
	if db == nil {
		dbInitLock.Lock()
		defer dbInitLock.Unlock()
		db, err = sqlx.Open("postgres", connStr)
		if err != nil {
			log.Error().Msg(fmt.Sprintf("Db connection not acquired: %v", err))
		}
		db.SetConnMaxLifetime(time.Minute * 5)
		db.SetMaxOpenConns(5)
		db.SetMaxIdleConns(5)
	}
	return db
}
