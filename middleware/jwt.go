package middleware

import (
	"os"
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


	secretKey := os.Getenv("JWT_SECRET_KEY")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payLoad)
    t, err := token.SignedString([]byte(secretKey))
    if err != nil {
        
        panic(err)
    }

    return t
}
