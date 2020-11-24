package utils

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var mySigningKey = []byte("AllYourBase")

// MyCustomClaims 定制
type MyCustomClaims struct {
	UserID uint `json:"user_id"`
	jwt.StandardClaims
}

// SignToken 签发 token
func SignToken(userID uint) (string, error) {

	// Create the Claims
	claims := &MyCustomClaims{
		userID,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24 * 2).Unix(),
			Issuer:    "test",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(mySigningKey)
}

// ParseToken 解析 token
func ParseToken(tokenString string) (uint, error) {

	token, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})

	if claims, ok := token.Claims.(*MyCustomClaims); ok && token.Valid {
		fmt.Printf("%v %v", claims.UserID, claims.StandardClaims.ExpiresAt)
		return claims.UserID, nil
	}
	return 0, err

}
