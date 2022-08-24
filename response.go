package nutshttp

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type APIMessage struct {
	Code    int
	Message string
}

var (
	APIOK = APIMessage{Code: 200, Message: "OK"}

	ErrBadRequest          = APIMessage{Code: 400, Message: "Bad Request"}
	AuthFail               = APIMessage{401, "Auth Fail"}
	ErrNotFound            = APIMessage{404, "Not Found"}
	ErrInternalServerError = APIMessage{500, "Internal Server Error"}
	ErrKeyNotFoundInBucket = APIMessage{40001, "Key Not Found In Bucket"}
	ErrPrefixScan          = APIMessage{40002, "Prefix Scans Not Found"}
	ErrPrefixSearchScan    = APIMessage{40003, "Prefix Search Scans Not Found"}
	ErrRangeScan           = APIMessage{40004, "Range Scans Not Found"}
	ErrBucketEmpty         = APIMessage{40005, "Bucket Empty"}
	ErrUnknown             = APIMessage{40404, "UnKnown Err"}
)

type Response struct {
	Data   interface{}         `json:"data"`
	Header map[string][]string `json:"header,omitempty"`

	Code  int    `json:"code"`
	Error string `json:"error,omitempty"`
}

func WriteSucc(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code: APIOK.Code,
		Data: data,
	})
}

func WriteError(c *gin.Context, msg APIMessage) {
	c.JSON(http.StatusOK, Response{
		Code:  msg.Code,
		Error: msg.Message,
	})
}
