package auth

import (
	"os"

	"github.com/golang-jwt/jwt/v5"
	"gitub.com/umardev500/gopos/pkg/constant"
)

func GenerateJWT(claims jwt.MapClaims) (string, error) {
	secret := os.Getenv("JWT_SECRET")

	// Create the token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(secret))
}

func VerifyJWT(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}

		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, constant.ErrInvalidToken
}
