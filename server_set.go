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
		Value []string `json:"value" binding:"required"`
	}

	var (
		err           error
		baseUri       BaseUri
		addSetRequest AddSetRequest
	)

	if err = c.ShouldBindUri(&baseUri); err != nil {
		WriteError(c, APIMessage{
			Message: err.Error(),
		})
		return
	}

	if err = c.ShouldBindJSON(&addSetRequest); err != nil {
		WriteError(c, APIMessage{
			Message: err.Error(),
		})
		return
	}

	if err = s.core.addSet(baseUri.Bucket, baseUri.Key, addSetRequest.Value...); err != nil {
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

	if err = c.ShouldBindUri(&baseUri); err != nil {
		WriteError(c, APIMessage{
			Message: err.Error(),
		})
		return
	}

	if resp.Items, err = s.core.listSet(baseUri.Bucket, baseUri.Key); err != nil {
		WriteError(c, APIMessage{
			Message: err.Error(),
		})
		return
	}

	WriteSucc(c, resp)
}

func (s *NutsHTTPServer) SAreMembers(c *gin.Context) {

	type (
		SAreMembersRequest struct {
			Value []string `json:"value" binding:"required"`
		}

		SAreMembersResp struct {
			IsExist bool `json:"is_exist"`
		}
	)

	var (
		ok               bool
		err              error
		baseUri          BaseUri
		sAreMembersquest SAreMembersRequest
	)

	if err = c.ShouldBindUri(&baseUri); err != nil {
		WriteError(c, APIMessage{
			Message: err.Error(),
		})
		return
	}

	if err = c.ShouldBindJSON(&sAreMembersquest); err != nil {
		WriteError(c, APIMessage{
			Message: err.Error(),
		})
		return
	}

	if ok, err = s.core.sAreMembers(baseUri.Bucket, baseUri.Key, sAreMembersquest.Value...); err != nil {
		WriteError(c, APIMessage{
			Message: err.Error(),
		})
		return
	}

	WriteSucc(c, SAreMembersResp{
		IsExist: ok,
	})
}

func (s *NutsHTTPServer) SIsMember(c *gin.Context) {

	type (
		SIsMemberRequest struct {
			Value string `json:"value" binding:"required"`
		}

		SIsMemberResp struct {
			IsExist bool `json:"is_exist"`
		}
	)

	var (
		ok                bool
		err               error
		baseUri           BaseUri
		sIsMembersRequest SIsMemberRequest
	)

	if err = c.ShouldBindUri(&baseUri); err != nil {
		WriteError(c, APIMessage{
			Message: err.Error(),
		})

		return
	}

	if err = c.ShouldBindJSON(&sIsMembersRequest); err != nil {
		WriteError(c, APIMessage{
			Message: err.Error(),
		})

		return
	}

	if ok, err = s.core.sIsMember(baseUri.Bucket, baseUri.Key, sIsMembersRequest.Value); err != nil {
		WriteError(c, APIMessage{
			Message: err.Error(),
		})

		return
	}

	WriteSucc(c, SIsMemberResp{
		IsExist: ok,
	})
}

func (s *NutsHTTPServer) SCard(c *gin.Context) {

	type ScardResp struct {
		Num int `json:"num"`
	}

	var (
		err     error
		baseUri BaseUri
		resp    ScardResp
	)

	if err = c.ShouldBindUri(&baseUri); err != nil {
		WriteError(c, APIMessage{
			Message: err.Error(),
		})
		return
	}

	if resp.Num, err = s.core.sCard(baseUri.Bucket, baseUri.Key); err != nil {
		WriteError(c, APIMessage{
			Message: err.Error(),
		})
		return
	}

	WriteSucc(c, resp)
}
