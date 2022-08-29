package nutshttp

func (s *NutsHTTPServer) initSetRouter() {
	sr := s.r.Group("/set").Use(JWT())

	sr.POST("/sadd/:bucket/:key", s.SAdd)

	sr.POST("/saremembers/:bucket/:key", s.SAreMembers)

	sr.POST("/sismember/:bucket/:key", s.SIsMember)

	sr.GET("/smembers/:bucket/:key", s.SMembers)

	sr.GET("/scard/:bucket/:key", s.SCard)
}
