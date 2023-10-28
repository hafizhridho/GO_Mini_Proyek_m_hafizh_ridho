package middleware

import (
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JwtCustomClaims struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	jwt.RegisteredClaims
}


func GenerateJWT (id uint, name string) string {
	var payLoad JwtCustomClaims
	payLoad.ID = id
	payLoad.Name = name
	payLoad.ExpiresAt = jwt.NewNumericDate(time.Now().Add(time.Hour * 72))

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payLoad)

	secretKey := os.Getenv("JWT_SECRET")
	if secretKey == "" {
		log.Println("jwt secret key tidak di pasang")
	}
	t, _ := token.SignedString([]byte(secretKey))
    return t
}
