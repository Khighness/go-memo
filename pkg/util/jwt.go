package util

import (
	"github.com/dgrijalva/jwt-go"
	"os"
	"time"
)

// @Author KHighness
// @Update 2022-02-13

const issuer = "memo"

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

type Claims struct {
	jwt.StandardClaims
	Id        uint   `json:"id"`
	Username  string `json:"username"`
	Authority int    `json:"authority"`
}

// 签发Token
func GenerateToken(id uint, username string, authority int) (string, error)  {
	currentTime := time.Now()
	expiredTime := currentTime.Add(time.Hour * 24)
	claims := Claims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiredTime.Unix(),
			Issuer:    issuer,
		},
		Id:        id,
		Username:  username,
		Authority: authority,
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)
	return token, err
}

// 验证Token
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
