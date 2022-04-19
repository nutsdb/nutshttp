package nutshttp

func (s *NutsHTTPServer) initStringRouter() {
	sr := s.r.Group("/string")

	sr.GET("/get/:bucket/:key", s.SGet)

	sr.POST("update/:bucket/:key", s.SUpdate)

	sr.POST("delete/:bucket/:key", s.SDelete)
}
