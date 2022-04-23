package nutshttp

import (
	"github.com/gin-gonic/gin"
	"github.com/xujiajun/nutsdb/ds/zset"
)

var (
	ErrBadUri       = APIMessage{Code: 4101, Message: "wrong path specification"}
	ErrBadJsonParam = APIMessage{Code: 4102, Message: "wrong json specification"}
)

type Node struct {
	Key   string `json:"key"`
	Value []byte `json:"value"`
}

func (s *NutsHTTPServer) ZAdd(c *gin.Context) {
	type Params struct {
		Score float64 `json:"score" binding:"required"`
		Value string  `json:"value" binding:"required"`
	}
	var (
		err     error
		baseUri BaseUri
		params  Params
	)

	if err = c.ShouldBindUri(&baseUri); err != nil {
		WriteError(c, APIMessage{Message: err.Error()})
		return
	}
	if err = c.ShouldBindJSON(&params); err != nil {
		WriteError(c, APIMessage{Message: err.Error()})
		return
	}
	if err = s.core.zPut(baseUri.Bucket, []byte(baseUri.Key), params.Score, []byte(params.Value)); err != nil {
		WriteError(c, APIMessage{Message: err.Error()})
		return
	}
	WriteSucc(c, APIOK)
}

func (s *NutsHTTPServer) ZCard(c *gin.Context) {
	var (
		err     error
		num     int
		baseUri BaseUri
	)

	if err = c.ShouldBindUri(&baseUri); err != nil {
		WriteError(c, APIMessage{Message: err.Error()})
		return
	}

	if num, err = s.core.zCard(baseUri.Bucket); err != nil {
		WriteError(c, APIMessage{Message: err.Error()})
		return
	}
	WriteSucc(c, num)
}

func (s *NutsHTTPServer) ZCount(c *gin.Context) {
	type Params struct {
		Start        float64 `json:"start" binding:"required"`
		End          float64 `json:"end" binding:"required"`
		Limit        int     `json:"limit,omitempty"`
		ExcludeStart bool    `json:"exclude_start,omitempty"`
		ExcludeEnd   bool    `json:"exclude_end,omitempty"`
	}

	type Response struct {
		Count int `json:"count"`
	}

	var (
		count   int
		err     error
		baseUri BaseUri
		p       Params
	)

	if err = c.ShouldBindUri(&baseUri); err != nil {
		WriteError(c, APIMessage{Message: err.Error()})
		return
	}

	if err = c.ShouldBindJSON(&p); err != nil {
		WriteError(c, APIMessage{Message: err.Error()})
		return
	}

	if count, err = s.core.zCount(baseUri.Bucket, p.Start, p.End, &zset.GetByScoreRangeOptions{
		Limit: p.Limit, ExcludeStart: p.ExcludeStart, ExcludeEnd: p.ExcludeEnd}); err != nil {
		WriteError(c, APIMessage{Message: err.Error()})
		return
	}
	WriteSucc(c, Response{Count: count})

}

func (s *NutsHTTPServer) ZGetByKey(c *gin.Context) {
	type Response struct {
		Node Node `json:"node"`
	}
	var (
		node    *zset.SortedSetNode
		err     error
		baseUri BaseUri
	)

	if err = c.ShouldBindUri(&baseUri); err != nil {
		WriteError(c, APIMessage{Message: err.Error()})
		return
	}

	if node, err = s.core.zGetByKey(baseUri.Bucket, baseUri.Key); err != nil {
		WriteError(c, APIMessage{Message: err.Error()})
		return
	}

	WriteSucc(c, Response{Node{Key: node.Key(), Value: node.Value}})

}
func (s *NutsHTTPServer) ZMembers(c *gin.Context) {

	type Response struct {
		Nodes []Node `json:"nodes"`
	}

	var (
		err     error
		nodes   []*zset.SortedSetNode
		baseUri BaseUri
		res     Response
	)

	if err = c.ShouldBindUri(&baseUri); err != nil {
		WriteError(c, APIMessage{Message: err.Error()})
		return
	}

	if nodes, err = s.core.zMembers(baseUri.Bucket); err != nil {
		WriteError(c, APIMessage{Message: err.Error()})
		return
	}
	for _, node := range nodes {
		res.Nodes = append(res.Nodes, Node{node.Key(), node.Value})
	}
	WriteSucc(c, res)
}
func (s *NutsHTTPServer) ZPeekMax(c *gin.Context) {
	type Response struct {
		Node Node `json:"node"`
	}
	var (
		node    *zset.SortedSetNode
		err     error
		baseUri BaseUri
	)

	if err = c.ShouldBindUri(&baseUri); err != nil {
		WriteError(c, APIMessage{Message: err.Error()})
		return
	}

	if node, err = s.core.zPeekMax(baseUri.Bucket); err != nil {
		WriteError(c, APIMessage{Message: err.Error()})
		return
	}

	WriteSucc(c, Response{Node{node.Key(), node.Value}})

}

func (s *NutsHTTPServer) ZPeekMin(c *gin.Context) {
	type Response struct {
		Node Node `json:"node"`
	}
	var (
		node    *zset.SortedSetNode
		err     error
		baseUri BaseUri
	)

	if err = c.ShouldBindUri(&baseUri); err != nil {
		WriteError(c, APIMessage{Message: err.Error()})
		return
	}

	if node, err = s.core.zPeekMin(baseUri.Bucket); err != nil {
		WriteError(c, APIMessage{Message: err.Error()})
		return
	}

	WriteSucc(c, Response{Node{node.Key(), node.Value}})
}

func (s *NutsHTTPServer) ZPopMax(c *gin.Context) {
	type Response struct {
		Node Node `json:"node"`
	}
	var (
		node    *zset.SortedSetNode
		err     error
		baseUri BaseUri
	)

	if err = c.ShouldBindUri(&baseUri); err != nil {
		WriteError(c, APIMessage{Message: err.Error()})
		return
	}

	if node, err = s.core.zPopMax(baseUri.Bucket); err != nil {
		WriteError(c, APIMessage{Message: err.Error()})
		return
	}

	WriteSucc(c, Response{Node{node.Key(), node.Value}})
}
func (s *NutsHTTPServer) ZPopMin(c *gin.Context) {
	type Response struct {
		Nodes Node `json:"nodes"`
	}
	var (
		node    *zset.SortedSetNode
		err     error
		baseUri BaseUri
	)

	if err = c.ShouldBindUri(&baseUri); err != nil {
		WriteError(c, APIMessage{Message: err.Error()})
		return
	}

	if node, err = s.core.zPopMin(baseUri.Bucket); err != nil {
		WriteError(c, APIMessage{Message: err.Error()})
		return
	}

	WriteSucc(c, Response{Node{node.Key(), node.Value}})
}

func (s *NutsHTTPServer) ZRangeByRank(c *gin.Context) {
	type Response struct {
		Nodes []Node `json:"nodes"`
	}
	type Params struct {
		Start int `json:"start"`
		End   int `json:"end"`
	}
	var (
		err     error
		params  Params
		reps    Response
		nodes   []*zset.SortedSetNode
		baseUri BaseUri
	)

	if err = c.ShouldBindUri(&baseUri); err != nil {
		WriteError(c, APIMessage{Message: err.Error()})
		return
	}
	if err = c.ShouldBindJSON(&params); err != nil {
		WriteError(c, APIMessage{Message: err.Error()})
		return
	}

	if nodes, err = s.core.zRangeByRank(baseUri.Bucket, params.Start, params.End); err != nil {
		WriteError(c, APIMessage{Message: err.Error()})
		return
	}
	for _, node := range nodes {
		reps.Nodes = append(reps.Nodes, Node{node.Key(), node.Value})
	}

	WriteSucc(c, reps)
}

func (s *NutsHTTPServer) ZRangeByScore(c *gin.Context) {
	type Params struct {
		Start        float64 `json:"start" binding:"required"`
		End          float64 `json:"end" binding:"required"`
		Limit        int     `json:"limit,omitempty"`
		ExcludeStart bool    `json:"exclude_start,omitempty"`
		ExcludeEnd   bool    `json:"exclude_end,omitempty"`
	}

	type Response struct {
		Nodes []Node `json:"nodes"`
	}

	var (
		reps    Response
		err     error
		nodes   []*zset.SortedSetNode
		baseUri BaseUri
		p       Params
	)

	if err = c.ShouldBindUri(&baseUri); err != nil {
		WriteError(c, ErrBadUri)
		return
	}

	if err = c.ShouldBindJSON(&p); err != nil {
		WriteError(c, ErrBadJsonParam)
		return
	}

	if nodes, err = s.core.zRangeByScore(baseUri.Bucket, p.Start, p.End, &zset.GetByScoreRangeOptions{
		Limit: p.Limit, ExcludeStart: p.ExcludeStart, ExcludeEnd: p.ExcludeEnd}); err != nil {
		WriteError(c, APIMessage{4000, err.Error()})
		return
	}
	for _, node := range nodes {
		reps.Nodes = append(reps.Nodes, Node{node.Key(), node.Value})
	}

	WriteSucc(c, reps)

}
func (s *NutsHTTPServer) ZRank(c *gin.Context) {
	type Response struct {
		Rank int `json:"rank"`
	}
	var (
		rank    int
		err     error
		baseUri BaseUri
	)

	if err = c.ShouldBindUri(&baseUri); err != nil {
		WriteError(c, APIMessage{Message: err.Error()})
		return
	}

	if rank, err = s.core.zRank(baseUri.Bucket, baseUri.Key); err != nil {
		WriteError(c, APIMessage{Message: err.Error()})
		return
	}
	WriteSucc(c, Response{Rank: rank})

}
func (s *NutsHTTPServer) ZRevRank(c *gin.Context) {
	type Response struct {
		Rank int `json:"rank"`
	}
	var (
		rank    int
		err     error
		baseUri BaseUri
	)

	if err = c.ShouldBindUri(&baseUri); err != nil {
		WriteError(c, APIMessage{Message: err.Error()})
		return
	}

	if rank, err = s.core.zRevRank(baseUri.Bucket, baseUri.Key); err != nil {
		WriteError(c, APIMessage{Message: err.Error()})
		return
	}
	WriteSucc(c, Response{Rank: rank})
}
func (s *NutsHTTPServer) ZRem(c *gin.Context) {
	var (
		err     error
		baseUri BaseUri
	)

	if err = c.ShouldBindUri(&baseUri); err != nil {
		WriteError(c, APIMessage{Message: err.Error()})
		return
	}

	if err = s.core.zRem(baseUri.Bucket, baseUri.Key); err != nil {
		WriteError(c, APIMessage{Message: err.Error()})
		return
	}
	WriteSucc(c, APIOK)
}
func (s *NutsHTTPServer) ZRemRangeByRank(c *gin.Context) {
	type Params struct {
		Start int `json:"start"`
		End   int `json:"end"`
	}
	var (
		err     error
		params  Params
		baseUri BaseUri
	)

	if err = c.ShouldBindUri(&baseUri); err != nil {
		WriteError(c, APIMessage{Message: err.Error()})
		return
	}

	if err = c.ShouldBindJSON(&params); err != nil {
		WriteError(c, APIMessage{Message: err.Error()})
		return
	}
	if err = s.core.ZRemRangeByRank(baseUri.Bucket, params.Start, params.End); err != nil {
		WriteError(c, APIMessage{Message: err.Error()})
		return
	}
	WriteSucc(c, APIOK)
}
func (s *NutsHTTPServer) ZScore(c *gin.Context) {
	type Response struct {
		Score float64 `json:"score"`
	}
	var (
		score   float64
		err     error
		baseUri BaseUri
	)

	if err = c.ShouldBindUri(&baseUri); err != nil {
		WriteError(c, APIMessage{Message: err.Error()})
		return
	}

	if score, err = s.core.zScore(baseUri.Bucket, baseUri.Key); err != nil {
		WriteError(c, APIMessage{Message: err.Error()})
		return
	}
	WriteSucc(c, Response{Score: score})
}
