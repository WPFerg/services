package httpHandlers

import (
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"encoding/json"

	"github.com/wpferg/services/httpHandlers/httpUtils"
	"github.com/wpferg/services/storage"
	"github.com/wpferg/services/structs"
)

func Remove(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	requestBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		httpUtils.HandleError(&w, 500, "Internal Server Error", "Error reading data from body", err)
		return
	}

	var id structs.ID

	err = json.Unmarshal(requestBody, &id)

	if err != nil {
		httpUtils.HandleError(&w, 400, "Bad Request", "Error unmarshalling", err)
		return
	}

	if storage.Remove(id.ID) {
		log.Println("Removed ID:", id.ID)

		w.WriteHeader(200)
		w.Write([]byte(strconv.Itoa(id.ID)))
	} else {
		httpUtils.HandleError(&w, 400, "Bad Request", "ID not found", nil)
	}
}
