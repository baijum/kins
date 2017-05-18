package project

import (
	"fmt"
	"testing"

	"github.com/baijum/kins/db"
	"github.com/boltdb/bolt"
)

func TestProjectCreate(t *testing.T) {
	db.DB.Update(func(tx *bolt.Tx) error {
		err := tx.DeleteBucket([]byte(db.ProjectBucket))
		if err != nil {
			return fmt.Errorf("delete bucket: %s", err)
		}
		_, err = tx.CreateBucket([]byte(db.ProjectBucket))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		return nil
	})

	s := Schema{Name: "somename", Description: "Some description"}
	err := s.create()
	if err != nil {
		t.Error(err)
	}
}
