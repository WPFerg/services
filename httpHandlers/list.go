package httpHandlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/wpferg/services/httpHandlers/httpUtils"
	"github.com/wpferg/services/storage"
)

func List(w http.ResponseWriter, r *http.Request) {
	var data, err = json.Marshal(storage.Get())

	if err != nil {
		httpUtils.HandleError(&w, 500, "Internal Server Error", "Error marshalling response", err)
		return
	}

	log.Println("Successfully returned data")
	w.Header().Add("Content-Type", "application/json")
	w.Write(data)
}
