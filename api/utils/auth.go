package utils

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

var Secret string

func init() {
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		panic("Something went wrong!")
	}
	Secret = hex.EncodeToString(b)
}

func CreateToken(username string, remember bool) (string, error) {
	expiry := time.Now().Add(time.Hour * 24).Unix()
	if remember {
		expiry = time.Now().Add(time.Hour * 168).Unix()
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    username,
		ExpiresAt: expiry,
	})

	token, err := claims.SignedString([]byte(Secret))
	if err != nil {
		return "", err
	}

	return token, err
}

func CreateResetToken(username string) (string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    username,
		ExpiresAt: time.Now().Add(time.Minute * 10).Unix(),
	})

	token, err := claims.SignedString([]byte(Secret))
	if err != nil {
		return "", err
	}

	return token, err
}

func CheckToken(usertoken string) (interface{}, error) {
	token, err := jwt.Parse(usertoken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(Secret), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims["iss"], err
	} else {
		return nil, err
	}
}
