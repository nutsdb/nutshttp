package nutshttp

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type user struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
}

func (s *NutsHTTPServer) initLoginRouter() {

	s.r.POST("/login", func(c *gin.Context) {
		u := &user{}
		err := c.ShouldBindJSON(u)
		if err != nil {
			WriteError(c, APIMessage{
				Message: err.Error(),
			})
			return
		}
		//Checking User Status
		users := viper.GetStringMapString("users")

		if password, ok := users[u.UserName]; !ok || password != u.Password {
			WriteError(c, AuthFail)
		} else {
			//Verification successful, write to JWT
			token, err := GenerateToken(u.UserName)
			if err != nil {
				WriteError(c, AuthFail)
			}
			WriteSucc(c, token)
		}

	})

}
