package nutshttp

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/xujiajun/nutsdb"
)

func Enable(db *nutsdb.DB) error {
	r := mux.NewRouter()

	r.HandleFunc("/", IndexHandler)

	registerSetHandler(r, db)

	// r.PathPrefix("/set/").Handler(makeSetHandler(db))
	// r.Handle("/", makeSetHandler(db))

	// r.PathPrefix("/set").Handler(makeSetHandler(db))
	// r.Handle("/set", makeSetHandler(db))

	return http.ListenAndServe(":8080", r)
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "hello world")
}

func registerSetHandler(r *mux.Router, db *nutsdb.DB) {
	sr := r.PathPrefix("/set/").Subrouter()

	sr.Handle("/123", listSet(db))
}
