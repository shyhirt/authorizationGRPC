package jwtUtil

import (
	"crypto/rsa"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"log"
	"os"
)

type JwtUtil struct {
	PrivateKey *rsa.PrivateKey
	PublicKey  *rsa.PublicKey
}

var errValidate = errors.New("incorrect user id")

func NewJwtUtil(key, pub string) *JwtUtil {
	pkey, err := os.ReadFile(key)
	if err != nil {
		log.Fatalln("Private key load", err)
	}
	signKey, err := jwt.ParseRSAPrivateKeyFromPEM(pkey)
	if err != nil {
		log.Fatalln("Private key", err)
	}
	pubKey, err := os.ReadFile(pub)
	if err != nil {
		log.Fatalln("Public key load", err)
	}
	verifyKey, err := jwt.ParseRSAPublicKeyFromPEM(pubKey)
	if err != nil {
		log.Fatalln("Public key", err)
	}
	return &JwtUtil{
		PrivateKey: signKey,
		PublicKey:  verifyKey,
	}
}

func (j *JwtUtil) GetJwtToken(iat, seconds, userId int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["userId"] = userId
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	token.Claims = claims
	return token.SignedString(j.PrivateKey)
}

func (j *JwtUtil) Validate(tokenString string) (int64, bool, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return j.PublicKey, nil
	})
	if err != nil {
		return 0, false, err
	}
	if !token.Valid {
		return 0, false, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	var userId int64

	switch i := claims["userId"].(type) {
	case int64:
		userId = i
	case float64:
		userId = int64(i)
	}
	if userId == 0 {
		err = errValidate
	}
	return userId, ok, err
}
