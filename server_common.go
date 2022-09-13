package nutshttp

import (
	"github.com/gin-gonic/gin"
)

func (s *NutsHTTPServer) GetAllBuckets(c *gin.Context) {

	type BaseUrl struct {
		DS  string `uri:"ds" binding:"required"`
		Reg string `uri:"reg" binding:"required"`
	}

	base := &BaseUrl{}

	if err := c.ShouldBindUri(base); err != nil {
		WriteError(c, APIMessage{
			Message: err.Error(),
		})
		return
	}

	buckets, err := s.core.GetAllBuckets(base.DS, base.Reg)

	if err != nil {
		WriteError(c, APIMessage{Message: err.Error()})
		return
	}

	WriteSucc(c, buckets)
	return
}
