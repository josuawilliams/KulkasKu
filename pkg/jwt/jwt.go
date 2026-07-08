package jwt

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(id int64, email string, name string, secretKey string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":    id,
		"email": email,
		"name":  name,
		"exp":   time.Now().Add(24 * time.Hour).Unix(),
	})

	key := []byte(secretKey)
	tokenStr, err := token.SignedString(key)
	return tokenStr, err
}

func ValidateToken(tokenStr, secretKey string, withClaimValidation bool) (int64, string, error) {
	var (
		key    = []byte(secretKey)
		claims = jwt.MapClaims{}
		token  *jwt.Token
		err    error
	)
	if withClaimValidation {
		token, err = jwt.ParseWithClaims(
			tokenStr,
			claims,
			func(t *jwt.Token) (interface{}, error) {
				return key, nil
			},
			jwt.WithoutClaimsValidation(),
		)
	} else {
		token, err = jwt.ParseWithClaims(
			tokenStr,
			claims,
			func(t *jwt.Token) (interface{}, error) {
				return key, nil
			},
		)
	}

	if err != nil {
		return 0, "", err
	}

	if token == nil || !token.Valid {
		return 0, "", fmt.Errorf("invalid token at (54)")
	}

	userId := int64(claims["id"].(float64))
	email := claims["email"].(string)

	return userId, email, nil
}
