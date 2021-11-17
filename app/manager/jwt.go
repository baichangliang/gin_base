package manager

import (
	"gin_test/app/models"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var jwtKey = []byte("a1b2c3d4e5f6g7h8i9j0k1l2m3n")

type Claims struct {
	ManagerId uint
	jwt.StandardClaims
}

func ReleaseToken(m models.Manager, jwtKey []byte) (string, error) {
	expirationTime := time.Now().Add(7 * 24 * time.Hour)
	claims := &Claims{
		ManagerId: m.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "master",
			Subject:   "bcl",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)

	return tokenString, err
}

func ParseToken(tokenString string) (*jwt.Token, *Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (i interface{}, err error) {
		return jwtKey, nil
	})

	return token, claims, err
}
