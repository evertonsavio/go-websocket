package main

import (
	"fmt"

	"github.com/dgrijalva/jwt-go"
	"github.com/havyx/golang-websocket/src/evertonsavio/application/services"
	"log"
	"net/http"
)

func main()  {

	mux := routes()

	log.Println("Starting channel listener")
	go services.ListenToWsChannel()

	log.Println("Starting on port 8080")

// 	token, err := getToken("123123")
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	fmt.Println(token)

// 	tc, err := verifyToken(token)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	fmt.Println(tc)

	_ = http.ListenAndServe(":8080", mux)

}

func getToken(name string) (string, error) {
	signingKey := []byte("PADOPADOPADOPADO")
	fmt.Println(signingKey)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": name,
	})
	tokenString, err := token.SignedString(signingKey)
	return tokenString, err
}

func verifyToken(tokenString string) (jwt.Claims, error) {
	signingKey := []byte("PADOPADOPADOPADO")
	fmt.Println(signingKey)
	testToken := "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJzdWIiOiIxMjMxMjMifQ.eoqrniuoQeG9hEwV9bJGBT5bFE5vn8CyPgenVZqwYDY"

	token, err := jwt.Parse(testToken, func(token *jwt.Token) (interface{}, error) { return signingKey, nil })
	if err != nil {
		return nil, err
	}
	return token.Claims, err
}
