package main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
)

func ExtractClaims(tokenStr string) {//(jwt.MapClaims, bool) {
	hmacSecretString := "n9c9823pct2793gtp@#&9" // Value
	hmacSecret := []byte(hmacSecretString)
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		 // check token signing method etc
		 if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		 }
		 
		 return hmacSecret, nil
	})
	// token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
	// 	return []byte(hmacSecret), nil
	// })
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(token)
	fmt.Println(token.Header["alg"])

	claims := token.Claims.(jwt.MapClaims)
	fmt.Println(claims["sub"])


}
