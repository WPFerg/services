package http_handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"log"

	"github.com/wpferg/services/storage"
	"github.com/wpferg/services/structs"
)

func Add(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	byteData, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Panicln("Error reading data from body", err.Error())
		w.WriteHeader(500)
		w.Write([]byte("Internal Server Error"))
		return
	}

	var message structs.Message

	err = json.Unmarshal(byteData, &message)

	if err != nil {
		log.Panicln("Error unmarshalling", err.Error())
		w.WriteHeader(500)
		w.Write([]byte("Internal Server Error"))
		return
	}

	id := storage.Add(message)

	log.Println("Added message:", message)

	jsonID, err := json.Marshal(structs.ID{ID: id})

	if err != nil {
		log.Panicln("Error marshalling response", err.Error())
		w.WriteHeader(500)
		w.Write([]byte("Internal Server Error"))
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(jsonID)
}
