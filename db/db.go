package db

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/baijum/kins/config"
	"github.com/baijum/logger"
	"github.com/boltdb/bolt"
)

var (
	// DB is the database connection
	DB *bolt.DB
)

// common buckets
const (
	ProjectBucket = "project"
	ItemBucket    = "item"
)

func createBucket(db *bolt.DB, name string) error {
	err := db.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists([]byte(name))
		if err != nil {
			return err
		}
		if logger.Level <= logger.DEBUG {
			log.Println("Bucket created:", bucket)
		}
		return nil
	})
	return err
}

func init() {
	var err error
	cwd, _ := os.Getwd()
	fmt.Println(cwd)
	file := config.Config.DatabaseFile()
	DB, err = bolt.Open(file, 0600, &bolt.Options{Timeout: 10 * time.Second})
	if err != nil {
		log.Fatal(err.Error())
	}

	err = createBucket(DB, ProjectBucket)
	if err != nil {
		log.Fatalf("Unable to create '%s' bucket: %s", ProjectBucket, err)
	}
	err = createBucket(DB, ItemBucket)
	if err != nil {
		log.Fatalf("Unable to create '%s' bucket: %s", ItemBucket, err)
	}
}
