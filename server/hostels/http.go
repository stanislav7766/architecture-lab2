package hostels

import (
	"encoding/json"
	"log"
	"net/http"
	"github.com/stanislav7766/architecture-lab2/server/tools"
)

// Channels HTTP handler.
type HttpHandlerFunc http.HandlerFunc

// HttpHandler creates a new instance of channels HTTP handler.
func HttpHandler(store *Store) HttpHandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		if r.URL.String() == "/getBestHostel" {
			handleGetBestHostel(r, rw, store)
		} else if r.URL.String() == "/sendStudent" {
			handleSendStudent(r, rw, store)
		} else {
			rw.WriteHeader(http.StatusMethodNotAllowed)
		}
	}
}

func handleSendStudent(r *http.Request, rw http.ResponseWriter, store *Store) {
	var sStudent SendStudent
	if err := json.NewDecoder(r.Body).Decode(&sStudent); err != nil {
		log.Printf("Error decoding channel input: %s", err)
		tools.WriteJsonBadRequest(rw, "bad JSON payload")
		return
	}
	res,err := store.setStudent(&sStudent)
	if err != nil {
		log.Printf("Error inserting record: %s", err)
		tools.WriteJsonCustomError(rw,err)
	} else {
		json.NewEncoder(rw).Encode(res)
	}
}

func handleGetBestHostel(r *http.Request, rw http.ResponseWriter, store *Store) {
	var hostel Hostel
	if err := json.NewDecoder(r.Body).Decode(&hostel); err != nil {
		log.Printf("Error decoding channel input: %s", err)
		tools.WriteJsonBadRequest(rw, "bad JSON payload")
		return
	}
	res, err := store.getBestHostel(&hostel)
	if err != nil {
		log.Printf("Error making query to the db: %s", err)
		tools.WriteJsonCustomError(rw,err)
		return
	}
	json.NewEncoder(rw).Encode(res)
}
