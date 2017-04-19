package httpHandlers

import (
	"net/http"

	"github.com/wpferg/services/httpHandlers/httpUtils"
	"github.com/wpferg/services/storage"
)

func List(w http.ResponseWriter, r *http.Request) {
	httpUtils.HandleSuccess(&w, storage.Get())
}
