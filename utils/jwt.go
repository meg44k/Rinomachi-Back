package utils

import (
	"context"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte(os.Getenv("JWT_SECRET_KEY"))

type JWTClaims struct {
	UID  string `json:"user_id"`
	Role string `json:"role"`
	jwt.RegisteredClaims
}

func GenerateJWT(UID, role string) (string, error) {
	claims := jwt.RegisteredClaims{
		Subject:   UID,
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)), // 有効期限: 24時間
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		Issuer:    "renomachi", // 任意の発行者名
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretKey)
}

func JWTAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			http.Error(w, "Missing token", http.StatusUnauthorized)
		}

		claims := &JWTClaims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return secretKey, nil
		})

		if err != nil || !token.Valid {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}
		ctx := context.WithValue(r.Context(), "userClaims", claims)
		next.ServeHTTP(w, r.WithContext(ctx))

	})
}
