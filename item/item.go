package item

import (
	"encoding/json"

	"github.com/baijum/kins/db"
	"github.com/baijum/kins/route"
	"github.com/boltdb/bolt"
)

// Data represents a item payload
type Data struct {
	Type       string            `json:"type"`
	ID         string            `json:"id"`
	Attributes map[string]string `json:"attributes"`
}

// Schema represents a database schema
type Schema struct {
	Title       string
	Description string
}

// Create creates a new item
func (obj *Schema) Create() (int, error) {
	var err error
	var id int

	objJSON, _ := json.Marshal(obj)

	err = db.DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(db.ItemBucket))
		id, _ := b.NextSequence()
		err := b.Put([]byte(string(id)), []byte(objJSON))
		return err
	})

	if err != nil {
		return -1, err
	}
	return id, nil
}

// New returns a schema
func New(d Data) *Schema {
	s := &Schema{}
	s.Title = d.Attributes["title"]
	s.Description = d.Attributes["description"]
	return s
}

func init() {
	route.RT.HandleFunc("/api/v1/items", createHandler).Methods("POST")
}
