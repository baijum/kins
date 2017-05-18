package project

import (
	"encoding/json"
	"errors"

	"github.com/baijum/kins/db"
	"github.com/baijum/kins/route"
	"github.com/boltdb/bolt"
)

// Data represents a project payload
type Data struct {
	Type       string            `json:"type"`
	ID         string            `json:"id"`
	Attributes map[string]string `json:"attributes"`
}

// Schema represents a database schema
type Schema struct {
	Name        string
	Description string
}

func (obj *Schema) create() error {
	var err error

	err = db.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(db.ProjectBucket))
		v := b.Get([]byte(obj.Name))
		if v != nil {
			return errors.New("Project already exists: " + obj.Name)
		}
		return nil
	})
	if err != nil {
		return err
	}

	objJSON, _ := json.Marshal(obj)

	err = db.DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(db.ProjectBucket))
		err := b.Put([]byte(obj.Name), []byte(objJSON))
		return err
	})

	if err != nil {
		return err
	}
	return nil
}

// New returns a schema
func New(d Data) *Schema {
	s := &Schema{}
	s.Name = d.Attributes["name"]
	s.Description = d.Attributes["description"]
	return s
}

func init() {
	route.RT.HandleFunc("/api/v1/projects", createHandler).Methods("POST")
}
