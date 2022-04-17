package nutshttp

func (s *NutsHTTPServer) initSetRouter() {
	sr := s.r.Group("/set")

	sr.POST("/sadd/:bucket/:key", s.SAdd)

	sr.GET("/smembers/:bucket/:key", s.SMembers)
}
