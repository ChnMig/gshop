package tools

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

// UserData jwt user info
type UserData struct {
	ID int `json:"id" db:"user_id"`
}

type myCustomClaims struct {
	Data UserData `json:"data"`
	jwt.StandardClaims
}

// JWTIssue issue jwt
func JWTIssue(d UserData) (string, error) {
	// set key
	mySigningKey := []byte(EnvConfig.JWT.Key)

	// Calculate expiration time
	nowTime := time.Now()
	expireTime := nowTime.Add(time.Duration(EnvConfig.JWT.Expiration) * time.Second)

	// Create the Claims
	claims := myCustomClaims{
		d,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "remoteAdmin",
		},
	}

	// issue
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	st, err := t.SignedString(mySigningKey)
	if err != nil {
		return "", err
	}
	return st, nil
}

// JWTDecrypt string token to data
func JWTDecrypt(st string) (*UserData, error) {
	token, err := jwt.ParseWithClaims(st, &myCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(EnvConfig.JWT.Key), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*myCustomClaims); ok && token.Valid {
		// success
		return &claims.Data, nil
	}
	return nil, err
}
