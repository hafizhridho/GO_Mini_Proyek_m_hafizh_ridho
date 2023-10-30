package middleware

import (
	"time"
	"github.com/golang-jwt/jwt/v5"
)

type JwtCustomClaims struct {
	ID   uint   `json:"id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}


func GenerateJWT (id uint, username string) string {
	var payLoad JwtCustomClaims
	payLoad.ID = id
	payLoad.Username = username
	payLoad.ExpiresAt = jwt.NewNumericDate(time.Now().Add(time.Hour * 72))

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payLoad)
	t, _ := token.SignedString([]byte("12345"))
	return t
}
