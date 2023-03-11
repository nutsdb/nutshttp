package nutshttp

func (s *NutsHTTPServer) initStringRouter() {
	sr := s.r.Group("/string").Use(JWT())

	sr.GET("get/:bucket/:key", s.Get)

	sr.POST("update/:bucket/:key", s.Update)

	sr.DELETE("delete/:bucket/:key", s.Delete)

	sr.DELETE("muldelete/:bucket", s.MulDelete)

	sr.GET("scan/:bucket/:scanType", s.Scan)
}
