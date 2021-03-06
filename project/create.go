package project

import (
	"encoding/json"
	"log"
	"net/http"
)

func createHandler(w http.ResponseWriter, r *http.Request) {
	payload := make(map[string]Data)
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&payload)
	if err != nil {
		log.Println("Unable to decode body: ", err)
		return
	}
	s := New(payload["data"])
	err = s.create()
	if err != nil {
		log.Println("Unable save data: ", err)
		return
	}
	tmpData := payload["data"]
	tmpData.ID = s.Name
	payload["data"] = tmpData
	b, err := json.Marshal(payload)
	if err != nil {
		log.Println("Unable marshal data: ", err)
		return
	}
	w.Write(b)
}
