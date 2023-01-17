package nutshttp

import (
	"github.com/gin-gonic/gin"
	"github.com/xujiajun/nutsdb"
)

func (s *NutsHTTPServer) Get(c *gin.Context) {
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

	value, err := s.core.Get(baseUri.Bucket, baseUri.Key)
	if err != nil {
		switch err {
		case nutsdb.ErrNotFoundKey:
			WriteError(c, ErrKeyNotFoundInBucket)
		default:
			WriteError(c, ErrUnknown)
		}
		return
	}

	WriteSucc(c, value)

}

func (s *NutsHTTPServer) Update(c *gin.Context) {
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

	err = s.core.Update(baseUri.Bucket, baseUri.Key, updateStringRequest.Value, updateStringRequest.Ttl)
	if err != nil {
		switch err {
		case nutsdb.ErrNotFoundKey:
			WriteError(c, ErrKeyNotFoundInBucket)
		default:
			WriteError(c, ErrUnknown)
		}
		return
	}
	WriteSucc(c, struct{}{})
}

func (s *NutsHTTPServer) MulDelete(context *gin.Context) {
	type MulDeleteRequest struct {
		Keys []string `json:"keys" binding:"required"`
	}
	var (
		err              error
		mulDeleteRequest MulDeleteRequest
	)
	//get bucket from uri
	param := context.Param("bucket")
	if param == "" {
		WriteError(context, ErrBucketEmpty)
		return
	}
	bucket := param
	if err = context.ShouldBindJSON(&mulDeleteRequest); err != nil {
		WriteError(context, APIMessage{
			Message: err.Error(),
		})
		return
	}

	for i := range mulDeleteRequest.Keys {
		err = s.core.Delete(bucket, mulDeleteRequest.Keys[i])
		//stop delete when error
		if err != nil {
			switch err {
			case nutsdb.ErrKeyEmpty:
				WriteError(context, ErrKeyNotFoundInBucket)
			default:
				WriteError(context, ErrUnknown)
			}
			return
		}
	}

	WriteSucc(context, struct{}{})

}

func (s *NutsHTTPServer) Delete(c *gin.Context) {
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

	_, err = s.core.Get(baseUri.Bucket, baseUri.Key)
	if err != nil {
		switch err {
		case nutsdb.ErrNotFoundKey:
			WriteError(c, ErrKeyNotFoundInBucket)
		default:
			WriteError(c, ErrUnknown)
		}
		return
	}

	err = s.core.Delete(baseUri.Bucket, baseUri.Key)

	if err != nil {
		switch err {
		case nutsdb.ErrKeyEmpty:
			WriteError(c, ErrKeyNotFoundInBucket)
		default:
			WriteError(c, ErrUnknown)
		}
		return
	}
	WriteSucc(c, struct{}{})
}

func (s *NutsHTTPServer) Scan(c *gin.Context) {
	const (
		PrefixScan       = "prefixScan"
		PrefixSearchScan = "prefixSearchScan"
		RangeScan        = "rangeScan"
		GetAll           = "getAll"
	)

	type ScanParam struct {
		Bucket   string `uri:"bucket" binding:"required"`
		ScanType string `uri:"scanType" binding:"required"`
	}

	var (
		err       error
		entries   nutsdb.Entries
		scanParam ScanParam
	)

	if err = c.ShouldBindUri(&scanParam); err != nil {
		WriteError(c, APIMessage{
			Message: err.Error(),
		})
		return
	}

	switch scanParam.ScanType {
	case PrefixScan:
		type ScanRequest struct {
			OffSet   *int    `form:"offSet"  binding:"required"`
			LimitNum *int    `form:"limitNum"  binding:"required"`
			Prefix   *string `form:"prefix" binding:"required"`
		}

		var scanReq ScanRequest
		if err = c.ShouldBindQuery(&scanReq); err != nil {
			WriteError(c, APIMessage{
				Message: err.Error(),
			})
			return
		}
		entries, err = s.core.PrefixScan(scanParam.Bucket, *scanReq.Prefix, *scanReq.OffSet, *scanReq.LimitNum)
		if err != nil {
			switch err {
			case nutsdb.ErrPrefixScan:
				WriteError(c, ErrPrefixScan)
			default:
				WriteError(c, ErrUnknown)
			}
			return
		}
		var res = map[string]string{}
		for _, e := range entries {
			res[string(e.Key)] = string(e.Value)
		}
		WriteSucc(c, res)
	case PrefixSearchScan:
		type ScanSearchReq struct {
			OffSet   *int    `form:"offSet"  binding:"required"`
			LimitNum *int    `form:"limitNum"  binding:"required"`
			Prefix   *string `form:"prefix" binding:"required"`
			Reg      *string `form:"reg" binding:"required"`
		}
		var scanSearchReq ScanSearchReq
		if err = c.ShouldBindQuery(&scanSearchReq); err != nil {
			WriteError(c, APIMessage{
				Message: err.Error(),
			})
			return
		}
		entries, err = s.core.PrefixSearchScan(scanParam.Bucket, *scanSearchReq.Prefix, *scanSearchReq.Reg, *scanSearchReq.OffSet, *scanSearchReq.LimitNum)
		if err != nil {
			switch err {
			case nutsdb.ErrPrefixSearchScan:
				WriteError(c, ErrPrefixSearchScan)
			default:
				WriteError(c, ErrUnknown)
			}
			return
		}
		var res = map[string]string{}
		for _, e := range entries {
			res[string(e.Key)] = string(e.Value)
		}
		WriteSucc(c, res)
	case RangeScan:
		type RangeScanReq struct {
			Start *string `form:"start" binding:"required"`
			End   *string `form:"end" binding:"required"`
		}
		var rangeScanReq RangeScanReq

		if err = c.ShouldBindQuery(&rangeScanReq); err != nil {
			WriteError(c, APIMessage{
				Message: err.Error(),
			})
			return
		}

		entries, err = s.core.RangeScan(scanParam.Bucket, *rangeScanReq.Start, *rangeScanReq.End)
		if err != nil {
			switch err {
			case nutsdb.ErrRangeScan:
				WriteError(c, ErrRangeScan)
			default:
				WriteError(c, ErrUnknown)
			}
			return
		}
		var res = map[string]string{}
		for _, e := range entries {
			res[string(e.Key)] = string(e.Value)
		}
		WriteSucc(c, res)
	case GetAll:
		entries, err = s.core.GetAll(scanParam.Bucket)
		if err != nil {
			switch err {
			case nutsdb.ErrBucketEmpty:
				WriteError(c, ErrBucketEmpty)
			default:
				WriteError(c, ErrUnknown)
			}
			return
		}
		var res = map[string]string{}
		for _, e := range entries {
			res[string(e.Key)] = string(e.Value)
		}
		WriteSucc(c, res)
	}

	return
}
