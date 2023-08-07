package jwt

import (
	"os"

	"github.com/dgrijalva/jwt-go"
)

type Gateway struct{}

func (jwtT *Gateway) ValidateToken(token string) (bool, map[string]interface{}, error) {
	claims := jwt.MapClaims{}
	if token == "" || token == "null" {
		return false, nil, nil
	}
	decodedToken, err := jwtT.decodeToken(token, claims)
	if err != nil {
		// log.Println(err)
		return false, nil, err
	}

	return decodedToken.Valid, claims, nil
}

func (jwtT *Gateway) decodeToken(token string, claims jwt.MapClaims) (*jwt.Token, error) {
	decodedToken, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_TOKEN_SECRET")), nil
	})
	return decodedToken, err
}
