package nutshttp

func (s *NutsHTTPServer) initStringRouter() {
	sr := s.r.Group("/string")

	sr.GET("/sget/:bucket/:key", s.SGet)

	sr.POST("sset/:bucket/:key", s.SUpdate)

	sr.POST("supdate/:bucket/:key", s.SUpdate)

	sr.POST("sdelete/:bucket/:key", s.SDelete)
}
