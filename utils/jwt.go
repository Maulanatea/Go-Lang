package utils

import (
	"fmt"

	"github.com/dgrijalva/jwt-go/v4"
)

var SecretKey = "SECRET_TOKEN"

func GenerateToken(claims *jwt.MapClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	webtoken, err := token.SignedString([]byte(SecretKey))
	if err != nil {
		return "", err
	}

	return webtoken, nil
}

func VerifyToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) { //kita lakukan parsing
		if _, isValid := t.Method.(*jwt.SigningMethodHMAC); !isValid {
			return nil, fmt.Errorf("unexpekted signing methd: %v", t.Header["alg"])
		}
		return []byte(SecretKey), nil
	})
	if err != nil { //apakah proses parsing ini ada masalah
		return nil, err //jika ada masalah
	}
	return token, nil //jika tidak error, kita kembalikan token nya

}

func DecodeToken(tokenS string) (jwt.MapClaims, error) {
	token, err := VerifyToken(tokenS)
	if err != nil {
		return nil, err
	}

	claims, isOk := token.Claims.(jwt.MapClaims)
	if isOk && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("invalide token")
}
