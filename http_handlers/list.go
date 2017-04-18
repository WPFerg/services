package http_handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/wpferg/services/storage"
)

func List(w http.ResponseWriter, r *http.Request) {
	var data, err = json.Marshal(storage.Get())

	if err != nil {
		log.Panicln("Error marshalling JSON")
		log.Panicln(err.Error())
		w.WriteHeader(500)
		w.Write([]byte("Internal Server Error"))
	}

	log.Println("Successfully returned data")
	w.Header().Add("Content-Type", "application/json")
	w.Write(data)
}
