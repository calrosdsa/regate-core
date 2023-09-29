package util

import (
	// "encoding/json"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"

	// "log"
	// "net/http"
	"time"
)

type Claims struct {
	UserId  int `json:"user_id"`
	Uuid string `json:"uuid"`
	Email string `json:"email"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

type ClaimsAdmin struct {
	Id int64 `json:"id"`
	UserId string `json:"admin_id"`
	Email string `json:"email"`
	Username string `json:"username"`
	Rol int8 `json:"rol"`
	EmpresaId int `json:"empresa_id"`
	jwt.RegisteredClaims
}

var sampleSecretKey = []byte(viper.GetString("JWT_SECRET"))

func GetToken(token string) string{
	return strings.TrimSpace(strings.Split(token, "Bearer")[1])
}

func GenerateJWT(userId int,uuid string, email string,username string) (string, error) {
	claims := &Claims{
		UserId:  userId,
		Uuid:uuid,
		Email: email,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(100 * time.Hour)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(sampleSecretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ExtractClaims(tokenString string) (*Claims, error) {
	claims := &Claims{}
	_, err := jwt.ParseWithClaims(tokenString, claims, func(tokenKey *jwt.Token) (interface{}, error) {
		return sampleSecretKey, nil
	})
	if err != nil {
		return claims, err
	}
	return claims, err
}

func GenerateShortAdminJWT(id int64,userId string, email string,username string,rol int8,empresaId int) (string, error) {
	claims := &ClaimsAdmin{
		Id:id,
		UserId:  userId,
		Email: email,
		Username: username,
		Rol: rol,
		EmpresaId: empresaId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Hour)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(sampleSecretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}


func GenerateAdminJWT(id int64,userId string, email string,username string,rol int8,empresaId int) (string, error) {
	claims := &ClaimsAdmin{
		UserId:  userId,
		Id: id,
		Email: email,
		Username: username,
		Rol: rol,
		EmpresaId: empresaId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(100 * time.Hour)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(sampleSecretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ExtractClaimsAdmin(tokenString string) (*ClaimsAdmin, error) {
	claims := &ClaimsAdmin{}
	_, err := jwt.ParseWithClaims(tokenString, claims, func(tokenKey *jwt.Token) (interface{}, error) {
		return sampleSecretKey, nil
	})
	if err != nil {
		return claims, err
	}
	return claims, err
}
