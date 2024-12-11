package helpers

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	"fmt"
)

func VerifyToken(tokenStr string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return []byte("mysecretpassword"), nil
	})

	if err != nil {
		return token, err
	}

	return token, err
}

func DecodeToken(tokenString string) (interface{}, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return "", nil
		}
		return []byte("mysecretpassword"), nil
	})
	
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println(claims["email"], "LETSSSSSS")
		return (claims["email"]), nil
	} else {
		fmt.Println(err)
	}

	fmt.Println("My Claims token", token)

	return "", nil
}

func GetEnv(key string) string {
	envFile, _ := godotenv.Read("../.env")
	return envFile[key]
}
