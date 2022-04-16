package nutshttp

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func (s *NutsHTTPServer) ListSet(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	bucket := vars["bucket"]
	key := vars["setname"]

	items, err := s.core.listSet(bucket, key)
	if err != nil {
		WriteError(w, ErrInternalServerError)
		return
	}

	WriteSucc(w, items)
}

func (s *NutsHTTPServer) SMembers(w http.ResponseWriter, r *http.Request) {
	// NOTE(zy): to do

	log.Println("hello world")
}
