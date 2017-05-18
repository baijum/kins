package project_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/boltdb/bolt"

	"github.com/baijum/kins/db"
	. "github.com/baijum/kins/project"
	"github.com/baijum/kins/route"
)

func TestProjectCreateHandler(t *testing.T) {
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

	ts := httptest.NewServer(route.RT)
	defer ts.Close()
	n := []byte(`{
  "data": {
    "type": "projects",
    "attributes": {
      "name": "somename",
      "description": "Some description"
    }
  }
}`)
	reqPayload := make(map[string]Data)
	decoder1 := json.NewDecoder(bytes.NewReader(n))
	err := decoder1.Decode(&reqPayload)
	if err != nil {
		log.Println("Unable to decode body: ", err)
		return
	}

	req, _ := http.NewRequest("POST", ts.URL+"/api/v1/projects", bytes.NewReader(n))
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	respPayload := make(map[string]Data)
	decoder2 := json.NewDecoder(resp.Body)
	err = decoder2.Decode(&respPayload)
	if err != nil {
		t.Fatal("Unable to decode body: ", err)
	}
	respData := respPayload["data"]
	reqData := reqPayload["data"]
	reqData.ID = respData.ID
	if !reflect.DeepEqual(reqData, respData) {
		t.Errorf("Data not matching. \nOriginal: %#v\nNew Data: %#v", reqData, respData)
	}
}
