package httpHandlers

import "net/http"
import "log"
import "github.com/wpferg/services/httpHandlers/httpUtils"

func HandleRequest(w http.ResponseWriter, r *http.Request) {
	log.Println("Incoming Request:", r.Method)
	switch r.Method {
	case http.MethodGet:
		List(w, r)
		break
	case http.MethodPost:
		Add(w, r)
		break
	case http.MethodDelete:
		Remove(w, r)
		break
	default:
		httpUtils.HandleError(&w, 405, "Method not allowed", "Method not allowed", nil)
		break
	}
}
