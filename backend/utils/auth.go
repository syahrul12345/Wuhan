package utils

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

type Token struct {
	UserID   uint
	UserName string
	Exp      int64
	jwt.StandardClaims
}

// Auth wiill authenticate and return a token
func Auth(writer http.ResponseWriter, request *http.Request) bool {
	requestPath := request.URL.Path
	noAuth := []string{"/", "/create", "/api/createAccount", "login"}
	//check if response does not require authenthication
	for _, value := range noAuth {
		if value == requestPath || strings.Contains(requestPath, ".chunk.js") || strings.Contains(requestPath, ".chunk.csss") {
			return true
		}
	}
	//other wise it requires authentication
	response := make(map[string]interface{})
	tokenHeader := request.Header.Get("Authorization")

	if tokenHeader == "" {
		response = Message(false, "Missing auth token")
		writer.WriteHeader(http.StatusForbidden)
		writer.Header().Add("Content-Type", "application/json")
		Respond(writer, response)
		return false
	}
	splitted := strings.Split(tokenHeader, " ") //The token normally comes in format `Bearer {token-body}`, we check if the retrieved token matched this requirement
	if len(splitted) != 2 {
		response = Message(false, "Invalid/Malformed auth token")
		writer.WriteHeader(http.StatusForbidden)
		writer.Header().Add("Content-Type", "application/json")
		Respond(writer, response)
		return false
	}
	tokenPart := splitted[1] // the information that we're interested in
	tk := Token{}

	token, err := jwt.ParseWithClaims(tokenPart, tk, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("token_password")), nil
	})

	//malformed token, return 403
	if err != nil {
		fmt.Println(err)
		response = Message(false, "Malformed auth token")
		writer.WriteHeader(http.StatusForbidden)
		writer.Header().Add("Content-Type", "application/json")
		Respond(writer, response)
		return false
	}
	//token is invalid
	if !token.Valid {
		fmt.Println(token.Valid)
		response = Message(false, "Token is invalid")
		writer.WriteHeader(http.StatusForbidden)
		writer.Header().Add("Content-Type", "application/json")
		Respond(writer, response)
		return false
	}
	return false
}
