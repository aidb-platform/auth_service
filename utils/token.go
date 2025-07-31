package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte("your-secret-key") // Replace with env secret

type Claims struct {
	UserID string `json:"user_id"`
	OrgID  string `json:"org_id"`
	jwt.RegisteredClaims
}

type JWTClaim struct {
	UserID string `json:"user_id"`
	jwt.RegisteredClaims
}
func GenerateToken(userID, orgID string) (string, error) {
	expiration := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		UserID: userID,
		OrgID:  orgID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiration),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}
