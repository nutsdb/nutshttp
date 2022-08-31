package nutshttp

func (s *NutsHTTPServer) initZSetRouter() {
	sr := s.r.Group("/zset")

	sr.POST("/zadd/:bucket/:key", s.ZAdd)

	sr.GET("/zcard/:bucket/:key", s.ZCard)

	sr.GET("/zcount/:bucket/:key", s.ZCount)

	sr.GET("/zgetbykey/:bucket/:key", s.ZGetByKey)

	sr.GET("/zmembers/:bucket/:key", s.ZMembers)

	sr.GET("/zpeekmax/:bucket/:key", s.ZPeekMax)

	sr.GET("/zpeekmin/:bucket/:key", s.ZPeekMin)

	sr.DELETE("/zpopmax/:bucket/:key", s.ZPopMax)

	sr.DELETE("/zpopmin/:bucket/:key", s.ZPopMin)

	sr.GET("/zrangebyrank/:bucket/:key", s.ZRangeByRank)

	sr.GET("/zrangebyscore/:bucket/:key", s.ZRangeByScore)

	sr.GET("/zrank/:bucket/:key", s.ZRank)

	sr.GET("/zrevrank/:bucket/:key", s.ZRevRank)

	sr.DELETE("/zrem/:bucket/:key", s.ZRem)

	sr.DELETE("/zremrangebyrank/:bucket/:key", s.ZRemRangeByRank)

	sr.GET("/zscore/:bucket/:key", s.ZScore)

}
