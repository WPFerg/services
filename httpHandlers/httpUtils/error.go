package httpUtils

import (
	"log"
	"net/http"
)

func HandleError(w *http.ResponseWriter, code int, responseText string, logMessage string, err error) {
	errorMessage := ""

	if err != nil {
		errorMessage = err.Error()
	}

	log.Panicln(logMessage, errorMessage)
	(*w).WriteHeader(code)
	(*w).Write([]byte(responseText))
}
