package service

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

type JwtCustomClaim struct {
	Identity interface{} `json:"identity"`
	jwt.RegisteredClaims
}

type JwtService interface {
	GenerateToken(identity interface{}) string
	ExtractToken(tokenString string) (*jwt.Token, error)
}

type jwtService struct {
	secretKey string
}

func (js *jwtService) GenerateToken(identity interface{}) string {
	issued := time.Now().UTC()
	claim := JwtCustomClaim{
		identity,
		jwt.RegisteredClaims{
			ID:        uuid.NewString(),
			Issuer:    "Sigmatech Test",
			ExpiresAt: jwt.NewNumericDate(issued.Add(time.Hour * 3)),
			IssuedAt:  jwt.NewNumericDate(issued),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	tokenString, _ := token.SignedString([]byte(js.secretKey))
	return tokenString
}

func (js *jwtService) ExtractToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method %v", t.Header["alg"])
		}
		return []byte(js.secretKey), nil
	})
}

func NewJwtService(secretKey string) JwtService {
	return &jwtService{
		secretKey: secretKey,
	}
}
