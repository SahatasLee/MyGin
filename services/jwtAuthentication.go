package services

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTService interface {
	GenerateToken(userId uint) string
	ValidateToken(token string) (*jwt.Token, error)
}

type authCustomClaims struct {
	UserId string `json:"userId"`
	jwt.MapClaims
}

type jwtServices struct {
	secretKey string
	issure    string
}

func JWTAuthService() JWTService {
	return &jwtServices{
		secretKey: "secret",
		issure:    "Bikask",
	}
}

func (service *jwtServices) GenerateToken(userId uint) string {
	claims := &authCustomClaims{
		string(rune(userId)),
		jwt.MapClaims{
			"exp":      time.Now().Add(time.Hour * 24).Unix(),
			"Issure":   service.issure,
			"IssureAt": time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte("secret"))

	if err != nil {
		panic(err)
	}

	return t
}

func (service *jwtServices) ValidateToken(encodedToken string) (*jwt.Token, error) {
	return jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("secret"), nil
	})
}
