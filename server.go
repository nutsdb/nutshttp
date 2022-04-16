package nutshttp

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/xujiajun/nutsdb"
)

type NutsHTTPServer struct {
	// db *nutsdb.DB
	core *core

	r *mux.Router
}

func NewNutsHTTPServer(db *nutsdb.DB) *NutsHTTPServer {
	c := &core{db}

	s := &NutsHTTPServer{
		core: c,
		r:    mux.NewRouter(),
	}

	s.initRouter()

	return s
}

func (s *NutsHTTPServer) Run(addr string) error {
	return http.ListenAndServe(addr, s.r)
}

func (s *NutsHTTPServer) initRouter() {
	setRouter := s.r.PathPrefix("/set/").Subrouter()

	setRouter.HandleFunc("/{bucket}/{setname}", s.ListSet)
}
