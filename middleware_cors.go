package nutshttp

import "github.com/gin-contrib/cors"

func (s *NutsHTTPServer) initCorsMiddleware() {
	config := cors.DefaultConfig()
	config.AllowHeaders = append(config.AllowHeaders, "Authorization")
	config.AllowAllOrigins = true
	s.r.Use(cors.New(config))
}
