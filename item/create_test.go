package item

import (
	"fmt"
	"log"
	"testing"

	"github.com/baijum/kins/db"
	"github.com/boltdb/bolt"
)

func TestItemCreate(t *testing.T) {
	err := db.DB.Update(func(tx *bolt.Tx) error {
		err := tx.DeleteBucket([]byte(db.ItemBucket))
		if err != nil {
			return fmt.Errorf("delete bucket: %s", err)
		}
		_, err = tx.CreateBucket([]byte(db.ItemBucket))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		return nil
	})
	if err != nil {
		log.Println("Unable to update: ", err)
		return
	}

	s := Schema{Title: "sometitle", Description: "Some description"}
	id, err := s.Create()
	if err != nil {
		t.Error(err)
	}
	if id <= 0 {
		t.Errorf("Data not inserted. ID: %#v", id)
	}
}
