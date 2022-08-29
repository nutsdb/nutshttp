package nutshttp

func (s *NutsHTTPServer) initListRouter() {
	sr := s.r.Group("/list").Use(JWT())

	sr.GET("/lrange/:bucket/:key", s.LRange)

	sr.POST("/rpush/:bucket/:key", s.RPush)

	sr.POST("/lpush/:bucket/:key", s.LPush)

	sr.GET("/rpop/:bucket/:key", s.RPop)

	sr.GET("/lpop/:bucket/:key", s.LPop)

	sr.GET("/rpeek/:bucket/:key", s.RPeek)

	sr.GET("/lpeek/:bucket/:key", s.LPeek)

	sr.POST("/lrem/:bucket/:key", s.LRem)

	sr.POST("/lset/:bucket/:key", s.LSet)

	sr.POST("/ltrim/:bucket/:key", s.LTrim)

	sr.GET("/lsize/:bucket/:key", s.LSize)

}
