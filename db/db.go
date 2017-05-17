package db

import (
	"log"
	"time"

	"github.com/baijum/kins/config"
	"github.com/boltdb/bolt"
)

var (
	// DB is the database connection
	DB *bolt.DB
)

func init() {
	var err error
	file := config.Config.DatabaseFile()
	DB, err = bolt.Open(file, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		log.Fatal(err.Error())
	}

}
