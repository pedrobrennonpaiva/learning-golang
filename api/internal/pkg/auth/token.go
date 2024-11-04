package auth

import (
	"errors"
	"fmt"
	"golang-api/internal/config"
	"golang-api/internal/models/responses"
	"net/http"
	"strconv"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

func CreateToken(userId uint64) (responses.AuthResponse, error) {

	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["exp"] = time.Now().Add(time.Hour * 6).Unix()
	claims["userId"] = userId

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(config.GetConfig().SecretKey))
	if err != nil {
		return responses.AuthResponse{}, err
	}

	return responses.AuthResponse{ID: strconv.FormatUint(userId, 10), Token: tokenString, ExpiresAt: claims["exp"].(int64)}, nil
}

func ValidateToken(r *http.Request) error {
	tokenString := extractToken(r)

	token, err := jwt.Parse(tokenString, returnVerifyKey)
	if err != nil {
		return err
	}

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return nil
	}

	return errors.New("invalid token")
}

func extractToken(r *http.Request) string {
	token := r.Header.Get("Authorization")

	if len(strings.Split(token, " ")) != 2 {
		return ""
	}

	return strings.Split(token, " ")[1]
}

func returnVerifyKey(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("signing method unexpected! %v", token.Header["alg"])
	}

	return []byte(config.GetConfig().SecretKey), nil
}

func ExtractUserId(r *http.Request) (uint64, error) {
	tokenString := extractToken(r)
	token, err := jwt.Parse(tokenString, returnVerifyKey)
	if err != nil {
		return 0, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userId, ok := claims["userId"].(float64)
		if !ok {
			return 0, errors.New("invalid user ID in token")
		}
		return uint64(userId), nil
	}

	return 0, errors.New("invalid token")
}
