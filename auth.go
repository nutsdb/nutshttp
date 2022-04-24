package nutshttp

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"time"
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
		SigningMethod:  jwt.SigningMethodES512,
		checkFunc: func(s string) bool {
			return true
		},
	}
}

type TokenClaims struct {
	Cert string `json:"cert"`
	jwt.StandardClaims
}

var secret = []byte("78")

func GenerateToken(cert string) (string, error) {
	claims := TokenClaims{
		cert,
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
	err := c.ShouldBind(&cert)
	if err != nil {
		WriteError(c, ErrMissingParam)
		return
	}
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
