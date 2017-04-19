package httpHandlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"log"

	"github.com/wpferg/services/httpHandlers/httpUtils"
	"github.com/wpferg/services/storage"
	"github.com/wpferg/services/structs"
)

func Add(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	byteData, err := ioutil.ReadAll(r.Body)

	if err != nil {
		httpUtils.HandleError(&w, 500, "Internal Server Error", "Error reading data from body", err)
		return
	}

	var message structs.Message

	err = json.Unmarshal(byteData, &message)

	if err != nil {
		httpUtils.HandleError(&w, 400, "Bad Request", "Error unmarhsalling JSON", err)
		return
	}

	id := storage.Add(message)

	log.Println("Added message:", message)

	jsonID, err := json.Marshal(structs.ID{ID: id})

	if err != nil {
		httpUtils.HandleError(&w, 500, "Internal Server Error", "Error marshalling response", err)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(jsonID)
}
