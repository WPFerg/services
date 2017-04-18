package http_handlers

import (
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"encoding/json"

	"github.com/wpferg/services/storage"
	"github.com/wpferg/services/structs"
)

func Remove(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	requestBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Panicln("Error reading data from body", err.Error())

		w.WriteHeader(500)
		w.Write([]byte("Internal Server Error"))
		return
	}

	var id structs.ID

	err = json.Unmarshal(requestBody, &id)

	if err != nil {
		log.Panicln("Error unmarshalling", err.Error())

		w.WriteHeader(500)
		w.Write([]byte("Internal Server Error"))
		return
	}

	storage.Remove(id.ID)

	log.Println("Removed ID:", id.ID)

	w.WriteHeader(200)
	w.Write([]byte(strconv.Itoa(id.ID)))
}
