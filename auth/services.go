package auth

import (
	"errors"

	"github.com/dgrijalva/jwt-go"
)

type Service interface {
	GenerateToken(userID int) (string, error)
	ValidateToken(encodeToken string) (*jwt.Token, error)
}

type jwtService struct {
	SecretKeyJWT string
}

func NewService(SecretKeyJWT string) *jwtService {
	return &jwtService{SecretKeyJWT}
}

func (s *jwtService) GenerateToken(userID int) (string, error) {
	// claims is payload or data where is jwt bring
	claim := jwt.MapClaims{}
	claim["user_id"] = userID

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	signedToken, err := token.SignedString([]byte(s.SecretKeyJWT))
	if err != nil {
		return signedToken, err
	}

	return signedToken, nil
}

func (s *jwtService) ValidateToken(encodeToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(encodeToken, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("Invalid token")
		}
		return []byte(s.SecretKeyJWT), nil
	})

	if err != nil {
		return token, err
	}

	return token, nil
}
