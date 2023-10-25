package router

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

type TokenInfo struct {
	UserID uuid.UUID `json:"userid"`
}

type TokenClaims struct {
	UserID uuid.UUID `json:"userid"`
	NBF    time.Time `json:"nbf"`
	EXP    time.Time `json:"exp"`
	jwt.RegisteredClaims
}

// Create creates accessToken and refreshToken

func CreateToken(info TokenInfo, secret []byte) (string, string, error) {

	// 10 minutes & one day

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userid": info.UserID,
		"nbf":    time.Now(),
		"exp":    time.Now().Add(time.Minute * 10),
	})

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userid": info.UserID,
		"nbf":    time.Now(),
		"exp":    time.Now().AddDate(0, 0, 1),
	})

	// sign with secret
	// and get encoded token

	var err error
	accessTokenString, err := accessToken.SignedString(secret)
	if err != nil {
		log.Println(err)
		return "", "", errors.New("Could not create access token")
	}

	refreshTokenString, err := refreshToken.SignedString(secret)
	if err != nil {
		log.Println(err)
		return "", "", errors.New("Could not create refresh token")
	}

	return accessTokenString, refreshTokenString, nil
}

// Validate verifies token signature and expiration date
// returns nil if valid, err otherwise

func Validate(tokenString string, secret []byte) (error, *TokenClaims) {

	token, err := jwt.ParseWithClaims(
		tokenString, &TokenClaims{},
		func(token *jwt.Token) (interface{}, error) {

			// validate algorithm
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}

			return secret, nil
		})

	if err != nil {
		return err, nil
	}

	if claims, ok := token.Claims.(*TokenClaims); ok && token.Valid {
		if time.Now().Before(claims.EXP) {
			return nil, claims
		}
	}

	return errors.New("Could not create refresh token"), nil
}
