package auth

import (
	"os"
	"time"

	"github.com/MosesOnuh/airline-api/models"
	"github.com/dgrijalva/jwt-go"
)

func CreateToken(userId string) (string, error) {
	claims := &models.Claims{
		UserId: userId,
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(time.Minute * 20).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jwtTokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err
	}
	return jwtTokenString, nil

}

func ValidToken(jwtToken string) (*models.Claims, error) {
	claims := &models.Claims{}
	keyFunc := func(token *jwt.Token) (i interface{}, e error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	}
	token, err := jwt.ParseWithClaims(jwtToken, claims, keyFunc)
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, err
	}

	return claims, nil

}
