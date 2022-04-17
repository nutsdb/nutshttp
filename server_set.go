package nutshttp

import (
	"github.com/gin-gonic/gin"
)

type BaseUri struct {
	Bucket string `uri:"bucket" binding:"required"`
	Key    string `uri:"key" binding:"required"`
}

func (s *NutsHTTPServer) SAdd(c *gin.Context) {

	type AddSetRequest struct {
		Value []string `form:"value" binding:"required"`
	}

	var (
		err           error
		baseUri       BaseUri
		addSetRequest AddSetRequest
	)

	err = c.ShouldBindUri(&baseUri)
	if err != nil {
		WriteError(c, APIMessage{
			Message: err.Error(),
		})
		return
	}

	err = c.ShouldBindJSON(&addSetRequest)
	if err != nil {
		WriteError(c, APIMessage{
			Message: err.Error(),
		})
		return
	}

	err = s.core.addSet(baseUri.Bucket, baseUri.Key, addSetRequest.Value...)
	if err != nil {
		WriteError(c, APIMessage{
			Message: err.Error(),
		})
		return
	}

	WriteSucc(c, struct{}{})
}

func (s *NutsHTTPServer) SMembers(c *gin.Context) {

	type listSetResp struct {
		Items []string `json:"items"`
	}

	var (
		err     error
		baseUri BaseUri
		resp    listSetResp
	)

	err = c.ShouldBindUri(&baseUri)
	if err != nil {
		WriteError(c, APIMessage{
			Message: err.Error(),
		})
		return
	}

	resp.Items, err = s.core.listSet(baseUri.Bucket, baseUri.Key)
	if err != nil {
		WriteError(c, APIMessage{
			Message: err.Error(),
		})
		return
	}

	WriteSucc(c, resp)
}
