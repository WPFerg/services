package http_handlers

import "net/http"
import "log"

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
		log.Panicln("Method not allowed", r.Method)
		w.WriteHeader(405)
		w.Write([]byte("Method not allowed"))
		break
	}
}
