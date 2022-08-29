package nutshttp

import (
	"github.com/spf13/viper"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/xujiajun/nutsdb"
)

type NutsHTTPServer struct {
	core *core
	r    *gin.Engine
}

func NewNutsHTTPServer(db *nutsdb.DB) (*NutsHTTPServer, error) {
	c := &core{db}

	r := gin.Default()

	s := &NutsHTTPServer{
		core: c,
		r:    r,
	}

	err := s.InitConfig()
	if err != nil {
		return nil, err
	}

	s.initRouter()

	return s, nil
}

func (s *NutsHTTPServer) InitConfig() error {
	viper.SetConfigType("yaml")
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.SetDefault("port", "8080")

	err := viper.ReadInConfig()
	if err != nil {
		return err
	}

	err = viper.UnmarshalKey("JWT", &jwtSetting)

	jwtSetting.Expire *= time.Minute
	if err != nil {
		return err
	}

	return nil
}

func (s *NutsHTTPServer) Run(addr string) error {
	return http.ListenAndServe(addr, s.r)
}

func (s *NutsHTTPServer) initRouter() {
	s.initSetRouter()

	s.initListRouter()

	s.initStringRouter()

	s.initZSetRouter()

	s.initLoginRouter()

}
