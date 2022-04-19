package nutshttp

import (
	"github.com/gin-gonic/gin"
	"github.com/xujiajun/nutsdb"
)

func (s *NutsHTTPServer) SGet(c *gin.Context) {
	var (
		err     error
		baseUri BaseUri
	)

	if err = c.ShouldBindUri(&baseUri); err != nil {
		WriteError(c, APIMessage{
			Message: err.Error(),
		})
		return
	}

	value, err := s.core.SGet(baseUri.Bucket, baseUri.Key)
	if err != nil {
		// if key not exist, return the not-found err msg
		if err == nutsdb.ErrNotFoundKey {
			WriteError(c, ErrKeyNotFoundInBucket)
			return
		}
		WriteError(c, ErrInternalServerError)
		return
	}

	WriteSucc(c, value)

}

func (s *NutsHTTPServer) SUpdate(c *gin.Context) {
	type UpdateStringRequest struct {
		Value string `json:"value" binding:"required"`
		Ttl   uint32 `json:"ttl"`
	}
	var (
		err                 error
		baseUri             BaseUri
		updateStringRequest UpdateStringRequest
	)

	if err = c.ShouldBindUri(&baseUri); err != nil {
		WriteError(c, APIMessage{
			Message: err.Error(),
		})
		return
	}

	if err = c.ShouldBindJSON(&updateStringRequest); err != nil {
		WriteError(c, APIMessage{
			Message: err.Error(),
		})
		return
	}

	err = s.core.SUpdate(baseUri.Bucket, baseUri.Key, updateStringRequest.Value, updateStringRequest.Ttl)
	if err != nil {
		WriteError(c, ErrInternalServerError)
		return
	}
	WriteSucc(c, struct{}{})
}

func (s *NutsHTTPServer) SDelete(c *gin.Context) {
	var (
		err     error
		baseUri BaseUri
	)

	if err = c.ShouldBindUri(&baseUri); err != nil {
		WriteError(c, APIMessage{
			Message: err.Error(),
		})
		return
	}

	_, err = s.core.SGet(baseUri.Bucket, baseUri.Key)
	if err != nil {
		// if key not exist, return the not-found err
		if err == nutsdb.ErrNotFoundKey {
			WriteError(c, ErrKeyNotFoundInBucket)
			return
		}
		WriteError(c, ErrInternalServerError)
		return
	}

	err = s.core.SDelete(baseUri.Bucket, baseUri.Key)

	if err != nil {
		WriteError(c, ErrInternalServerError)
		return
	}
	WriteSucc(c, struct{}{})
}
