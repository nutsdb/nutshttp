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
	s.initSetRouter()

	s.initListRouter()
}

func (s *NutsHTTPServer) initSetRouter() {
	sr := s.r.Group("/set")

	sr.GET("/:bucket/:key", s.ListSet)
}

func (s *NutsHTTPServer) initListRouter() {
	sr := s.r.Group("/list")

	sr.GET("/:bucket/:key", s.LRange)
}
