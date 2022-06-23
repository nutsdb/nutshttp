package nutshttp

import (
	"errors"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type CheckFunc func(string) bool

type Option struct {
	ExpireDuration time.Duration
	Issuer         string
	SigningMethod  jwt.SigningMethod
	checkFunc      CheckFunc
}

//var checkFunc CheckFunc
var option Option

func defaultOptAndCheckFunc() Option {
	return Option{
		ExpireDuration: time.Hour * 4,
		Issuer:         "nutsdb.nutsdb-http",
		SigningMethod:  jwt.SigningMethodHS256,
		checkFunc: func(s string) bool {
			return true
		},
	}
}

type TokenClaims struct {
	Cert []byte `json:"cert"`
	jwt.StandardClaims
}

var secret []byte

func GenerateToken(cert string) (string, error) {
	claims := TokenClaims{
		[]byte(cert),
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(option.ExpireDuration).Unix(),
			Issuer:    option.Issuer,
		}}
	token := jwt.NewWithClaims(option.SigningMethod, claims)
	signedToken, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

func CheckToken(tokenStr string) (*TokenClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*TokenClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}

func handler(c *gin.Context) {
	if &option == nil {
		panic(errors.New("invalid auth option"))
	}
	var cert string
	cert = c.Param("cert")
	b := option.checkFunc(cert)
	if b {
		token, err := GenerateToken(cert)
		if err != nil {
			WriteError(c, ErrInternalServerError)
			return
		}
		WriteSucc(c, token)
		return
	}
	WriteError(c, ErrRefuseIssueToken)
}

func (s *NutsHTTPServer) DefaultInitAuth() {
	option = defaultOptAndCheckFunc()
	sr := s.r.Group("/auth")
	sr.GET("/:cert", handler)
	s.r.Use(JWTAuthMiddleware())
}

func SetSecret(s string) {
	secret = []byte(s)
}

const Bearer = "Bearer"

func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if len(authHeader) == 0 {
			WriteError(c, ErrAuthInvalid)
			c.Abort()
			return
		}
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == Bearer) {
			WriteError(c, ErrAuthInvalid)
			c.Abort()
			return
		}
		mc, err := CheckToken(parts[1])
		if err != nil {
			WriteError(c, ErrAuthInvalid)
			c.Abort()
			return
		}
		c.Set("cert", mc.Cert)
		c.Next()
	}
}
