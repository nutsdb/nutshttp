package nutshttp

func (s *NutsHTTPServer) initListRouter() {
	sr := s.r.Group("/list")

	sr.GET("/range/:bucket/:key", s.Range)

	sr.POST("/rpush/:bucket/:key", s.RPush)

	sr.POST("/lpush/:bucket/:key", s.LPush)

	sr.GET("/rpop/:bucket/:key", s.RPop)

	sr.GET("/lpop/:bucket/:key", s.LPop)

	sr.GET("/rpeek/:bucket/:key", s.RPeek)

	sr.GET("/lpeek/:bucket/:key", s.LPeek)

	sr.POST("/lem/:bucket/:key", s.Lem)

	sr.POST("/set/:bucket/:key", s.Set)

	sr.POST("/ltrim/:bucket/:key", s.LTrim)

	sr.GET("/size/:bucket/:key", s.Size)

}
