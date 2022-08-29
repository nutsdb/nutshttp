package nutshttp

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"time"
)

type JWTSetting struct {
	Secret string
	Issuer string
	Expire time.Duration
}

var jwtSetting JWTSetting

// Claims Defines the basic properties of the JWT
type Claims struct {
	Username string `json:"username,omitempty"`
	jwt.StandardClaims
}

// GetJWTSecret Get the JWT key
func GetJWTSecret() []byte {
	return []byte(jwtSetting.Secret)
}

// GenerateToken To generate the token
func GenerateToken(username string) (string, error) {
	nowTime := time.Now()
	expire := nowTime.Add(jwtSetting.Expire)
	claims := Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expire.Unix(),
			Issuer:    jwtSetting.Issuer,
		},
	}
	// padding the first field information, such as the method, type, etc
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedString, err := tokenClaims.SignedString(GetJWTSecret())
	return signedString, err
}

// ParseToken parse token
func ParseToken(token string) (claims *Claims, err error) {
	result, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return GetJWTSecret(), nil
	})
	if result != nil {
		claims, ok := result.Claims.(*Claims)
		if ok && result.Valid {
			return claims, nil
		}
	}
	return nil, err
}

//JWT middleware
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		//Look for it in the Query path, and if you can't find it, look for it in header
		token, ok := c.GetQuery("token")
		var err error
		if !ok {
			token = c.GetHeader("token")
			if token == "" {
				token, err = c.Cookie("token")
				if err != nil {
					WriteError(c, AuthFail)
					c.Abort()
				}
			}
		}
		//not find jwt
		if token == "" {
			WriteError(c, APIMessage{
				Code:    401,
				Message: "Token is not found",
			})
			c.Abort()
		} else {
			//parse jwt with error
			_, err := ParseToken(token)
			if err != nil {
				WriteError(c, APIMessage{
					Code:    401,
					Message: "JWT error with :" + err.Error(),
				})
				c.Abort()
			}
		}
		c.Next()
	}
}
