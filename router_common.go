package nutshttp

//common router
func (s *NutsHTTPServer) initCommonRouter() {
	sr := s.r.Group("/common").Use(JWT())

	sr.GET("/getAll/:ds/:reg", s.GetAllBuckets)

}
