package jwt

import (
	"time"
	"videocall/internal/app/auth/schema"

	"github.com/golang-jwt/jwt/v4"
)

var secretKey = []byte("your-256-bit-secret")

func CreateToken(userData *schema.TokenRequest) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":  userData.Id,
		"username": userData.Username,
		"role":     userData.Role,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ValidateToken(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil {
		return err
	}

	if !token.Valid {
		return err
	}

	return nil
}
