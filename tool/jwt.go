package tools

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type myCustomClaims struct {
	jwt.StandardClaims
}

var issue string = "gShop"
var subject string = "token"
var audienc string = "gShop"

// JWTIssue issue jwt
func JWTIssue(jwtID string) (string, error) {
	// set key
	mySigningKey := []byte(EnvConfig.JWT.Key)
	// Calculate expiration time
	nowTime := time.Now()
	expireTime := nowTime.Add(time.Duration(EnvConfig.JWT.Expiration) * time.Second)
	// Create the Claims
	claims := myCustomClaims{
		// https://en.wikipedia.org/wiki/JSON_Web_Token
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    issue,
			IssuedAt:  nowTime.Unix(),
			Subject:   subject,
			Audience:  audienc,
			NotBefore: nowTime.Unix(),
			Id:        jwtID,
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
func JWTDecrypt(tokenString string) (tokenID *string, err error) {
	token, err := jwt.ParseWithClaims(tokenString, &myCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(EnvConfig.JWT.Key), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*myCustomClaims); ok && token.Valid {
		// success
		return &claims.Id, nil
	}
	return nil, err
}
