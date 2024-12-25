package jwt

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)



type CustomClaims struct {
	UserEmail string `json:"userEmail"`
	jwt.RegisteredClaims
}

func GenerateToken(userEmail string) (string, error) {
	claims := CustomClaims{
		UserEmail: userEmail,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)), 
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	secretKey := getSecretyKey()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretKey)
}

func ValidateToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("método de assinatura inválido: %v", token.Header["alg"])
		}
		secretKey := getSecretyKey()
		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("token inválido")
}

func getSecretyKey() []byte{
	err2 := godotenv.Load()
    if err2 != nil {
		fmt.Println("Erro ao carregar o arquivo .env")
    }
	PASSWORD := os.Getenv("JWT_KEY")
	jwtSecret := []byte(PASSWORD)

	return jwtSecret
}