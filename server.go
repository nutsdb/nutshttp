package nutshttp

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xujiajun/nutsdb"
)

type NutsHTTPServer struct {
	core *core
	r    *gin.Engine
}

func NewNutsHTTPServer(db *nutsdb.DB) *NutsHTTPServer {
	c := &core{db}

	r := gin.Default()

	s := &NutsHTTPServer{
		core: c,
		r:    r,
	}

	s.initRouter()

	return s
}

func (s *NutsHTTPServer) Run(addr string) error {
	return http.ListenAndServe(addr, s.r)
}

func (s *NutsHTTPServer) initRouter() {
	if EnableAuth {
		s.DefaultInitAuth()
	}
	s.initSetRouter()

	s.initListRouter()

	s.initStringRouter()

	s.initZSetRouter()
}
