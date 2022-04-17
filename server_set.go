package nutshttp

import (
	"log"

	"github.com/gin-gonic/gin"
)

func (s *NutsHTTPServer) ListSet(c *gin.Context) {

	bucket := c.Param("bucket")
	key := c.Param("key")

	items, err := s.core.listSet(bucket, key)
	if err != nil {
		WriteError(c, ErrInternalServerError)
		return
	}

	WriteSucc(c, items)
}

func (s *NutsHTTPServer) SMembers(c *gin.Context) {
	// NOTE(zy): to do

	log.Println("hello world")
}
