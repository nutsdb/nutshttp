package nutshttp

import (
	"github.com/gin-gonic/gin"
)

type ReqLRange struct {
	Start int `form:"start"`
	End   int `form:"end"`
}

func (s *NutsHTTPServer) LRange(c *gin.Context) {

	bucket := c.Param("bucket")
	key := c.Param("key")

	var req ReqLRange
	if err := c.ShouldBindQuery(&req); err != nil {
		WriteError(c, ErrBadRequest)
		return
	}

	items, err := s.core.LRange(bucket, []byte(key), req.Start, req.End)
	if err != nil {
		WriteError(c, ErrInternalServerError)
		return
	}

	WriteSucc(c, items)
}
