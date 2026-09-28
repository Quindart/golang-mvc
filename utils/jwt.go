package utils

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecretKey = []byte("super-secret-key-change-me")

type Claim struct {
	UserID   int64  `json:"user_id"`
	FullName string `json:"full_name"`
	jwt.RegisteredClaims
}

func GenerateToken(userID int64, fullName string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)

	claims := &Claim{
		UserID:   userID,
		FullName: fullName,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "abc?",
		},
	}
	// Tạo token với thuật toán HS256 và claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtSecretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func ValidateToken(tokenStr string) (*Claim, error) {
	claim := &Claim{}
	token, err := jwt.ParseWithClaims(tokenStr, claim, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtSecretKey, nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, errors.New("invalid token")
	}
	return claim, nil
}
