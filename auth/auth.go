package auth

import (
	//"os"
	"time"

	"github.com/MosesOnuh/airline-api/models"
	"github.com/dgrijalva/jwt-go"

)
type TokenHandler interface {
	CreateToken(userId string) (string, error)
	ValidToken(jwtToken string) (*models.Claims, error)
}

type tokenHandler struct {
	jwtSecret string
}

// validate interface implementation
var _ TokenHandler = &tokenHandler{}

func New(secret string) TokenHandler {
	return &tokenHandler{
		jwtSecret: secret,
	}
}

func (t *tokenHandler) CreateToken(userId string) (string, error) {
	claims := &models.Claims{
		UserId: userId,
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(time.Hour * 1).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jwtTokenString, err := token.SignedString([]byte(t.jwtSecret))
	if err != nil {
		return "", err
	}
	return jwtTokenString, nil

}

func  (t *tokenHandler) ValidToken(jwtToken string) (*models.Claims, error) {
	claims := &models.Claims{}
	keyFunc := func(token *jwt.Token) (i interface{}, e error) {
		return []byte(t.jwtSecret), nil
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
