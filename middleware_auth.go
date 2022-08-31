package nutshttp

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

var Secret = []byte("nutshttp-secret")

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// GenerateToken generate a jwt based on claims
func GenerateToken(claims Claims) (string, error) {
	originalToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := originalToken.SignedString(Secret)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func parseToken(tokenString string) (*Claims, error) {
	var claims Claims
	token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		return Secret, nil
	})
	if err != nil || !token.Valid {
		jwtError := handleJWTError(err)
		return nil, jwtError
	} else {
		return &claims, nil
	}
}

func handleJWTError(err error) error {
	if validationError, ok := err.(*jwt.ValidationError); ok {
		return errors.New("Invalid Token")
	} else if validationError.Errors & (jwt.ValidationErrorExpired | jwt.ValidationErrorNotValidYet) != 0 {
		return errors.New("Token Expired")
	} else {
		return errors.New("Invalid Token")
	}
}

func (s *NutsHTTPServer) initAuthMiddleware() {
	s.r.Use(func(c *gin.Context) {
		// skip login api
		if c.Request.RequestURI == "/login" {
			c.Next()
			return
		}

		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			WriteError(c, AuthFail)
			c.Abort()
			return
		}

		_, err := parseToken(tokenString)
		if err != nil {
			WriteError(c, APIMessage{
				Message: err.Error(),
			})
			c.Abort()
			return
		}

		c.Next()
	})
}
