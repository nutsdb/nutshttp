package nutshttp

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type PushRequest struct {
	Value string `json:"value" binding:"required"`
}

func (s *NutsHTTPServer) LRange(c *gin.Context) {
	var (
		err     error
		baseUri BaseUri
	)

	type RangeReq struct {
		Start *int `form:"start" binding:"required"`
		End   *int `form:"end" binding:"required"`
	}

	var rangeReq RangeReq

	if err = c.ShouldBindUri(&baseUri); err != nil {
		WriteError(c, APIMessage{
			Message: err.Error(),
		})
		return
	}

	if err = c.ShouldBindQuery(&rangeReq); err != nil {
		WriteError(c, APIMessage{
			Message: err.Error(),
		})
		return
	}

	fmt.Println(*rangeReq.Start, " ", *rangeReq.End)

	items, err := s.core.LRange(baseUri.Bucket, baseUri.Key, *rangeReq.Start, *rangeReq.End)
	if err != nil {
		switch err {
		default:
			WriteError(c, ErrUnknown)
		}
		return
	}

	WriteSucc(c, items)
}

func (s *NutsHTTPServer) RPush(c *gin.Context) {
	var (
		err      error
		baseUri  BaseUri
		rPushReq PushRequest
	)

	if err = c.ShouldBindUri(&baseUri); err != nil {
		WriteError(c, APIMessage{
			Message: err.Error(),
		})
		return
	}

	if err = c.ShouldBindJSON(&rPushReq); err != nil {
		WriteError(c, APIMessage{
			Message: err.Error(),
		})
		return
	}

	err = s.core.RPush(baseUri.Bucket, baseUri.Key, rPushReq.Value)
	if err != nil {
		switch err {
		default:
			WriteError(c, ErrUnknown)
		}
	}

	WriteSucc(c, struct{}{})

}

func (s *NutsHTTPServer) LPush(c *gin.Context) {
	var (
		err      error
		baseUri  BaseUri
		rPushReq PushRequest
	)

	if err = c.ShouldBindUri(&baseUri); err != nil {
		WriteError(c, APIMessage{
			Message: err.Error(),
		})
		return
	}

	if err = c.ShouldBindJSON(&rPushReq); err != nil {
		WriteError(c, APIMessage{
			Message: err.Error(),
		})
		return
	}

	err = s.core.LPush(baseUri.Bucket, baseUri.Key, rPushReq.Value)
	if err != nil {
		switch err {
		default:
			WriteError(c, ErrUnknown)
		}
	}

	WriteSucc(c, struct{}{})
}

func (s *NutsHTTPServer) RPop(c *gin.Context) {
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

	pop, err := s.core.RPop(baseUri.Bucket, baseUri.Key)
	if err != nil {
		switch err {
		default:
			WriteError(c, ErrUnknown)
		}
	}
	WriteSucc(c, pop)
}

func (s *NutsHTTPServer) LPop(c *gin.Context) {
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

	pop, err := s.core.LPop(baseUri.Bucket, baseUri.Key)
	if err != nil {
		switch err {
		default:
			WriteError(c, ErrUnknown)
		}
	}
	WriteSucc(c, pop)
}

func (s *NutsHTTPServer) RPeek(c *gin.Context) {
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

	peek, err := s.core.RPeek(baseUri.Bucket, baseUri.Key)
	if err != nil {
		switch err {
		default:
			WriteError(c, ErrUnknown)
		}
	}
	WriteSucc(c, peek)
}

func (s *NutsHTTPServer) LPeek(c *gin.Context) {
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

	peek, err := s.core.RPeek(baseUri.Bucket, baseUri.Key)
	if err != nil {
		switch err {
		default:
			WriteError(c, ErrUnknown)
		}
	}
	WriteSucc(c, peek)
}

func (s *NutsHTTPServer) LRem(c *gin.Context) {
	var (
		err     error
		baseUri BaseUri
	)

	type LRemReq struct {
		Value *string `json:"value" binding:"required"`
		Count *int    `json:"count" binding:"required"`
	}

	if err = c.ShouldBindUri(&baseUri); err != nil {
		WriteError(c, APIMessage{
			Message: err.Error(),
		})
		return
	}

	var lremReq LRemReq
	if err = c.ShouldBindJSON(&lremReq); err != nil {
		WriteError(c, APIMessage{
			Message: err.Error(),
		})
		return
	}

	num, err := s.core.LRem(baseUri.Bucket, baseUri.Key, *lremReq.Value, *lremReq.Count)
	if err != nil {
		switch err {
		default:
			WriteError(c, ErrUnknown)
		}
		return
	}
	WriteSucc(c, num)
}

func (s *NutsHTTPServer) LSet(c *gin.Context) {
	var (
		err     error
		baseUri BaseUri
	)

	type SetReq struct {
		Value *string `json:"value" binding:"required"`
		Index *int    `json:"index" binding:"required"`
	}

	var setReq SetReq

	if err = c.ShouldBindUri(&baseUri); err != nil {
		WriteError(c, APIMessage{
			Message: err.Error(),
		})
		return
	}

	if err = c.ShouldBindJSON(&setReq); err != nil {
		WriteError(c, APIMessage{
			Message: err.Error(),
		})
		return
	}

	err = s.core.LSet(baseUri.Bucket, baseUri.Key, *setReq.Value, *setReq.Index)
	if err != nil {
		switch err {
		default:
			WriteError(c, ErrUnknown)
		}
		return
	}
	WriteSucc(c, struct{}{})

}

func (s *NutsHTTPServer) LTrim(c *gin.Context) {
	var (
		err     error
		baseUri BaseUri
	)

	type TrimReq struct {
		Start *int `json:"start" binding:"required"`
		End   *int `json:"end" binding:"required"`
	}

	var trimReq TrimReq

	if err = c.ShouldBindUri(&baseUri); err != nil {
		WriteError(c, APIMessage{
			Message: err.Error(),
		})
		return
	}

	if err = c.ShouldBindJSON(&trimReq); err != nil {
		WriteError(c, APIMessage{
			Message: err.Error(),
		})
		return
	}

	err = s.core.LTrim(baseUri.Bucket, baseUri.Key, *trimReq.Start, *trimReq.End)
	if err != nil {
		switch err {
		default:
			WriteError(c, ErrUnknown)
		}
		return
	}
	WriteSucc(c, struct{}{})
}

func (s *NutsHTTPServer) LSize(c *gin.Context) {
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
	size, err := s.core.LSize(baseUri.Bucket, baseUri.Key)
	if err != nil {
		switch err {
		default:
			WriteError(c, ErrUnknown)
		}
		return
	}

	WriteSucc(c, size)

}
