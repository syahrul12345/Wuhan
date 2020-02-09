package utils

import (
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Token struct {
	Email string
	Exp   int64
	jwt.StandardClaims
}

// Auth wiill authenticate and return a token
func Auth(writer http.ResponseWriter, request *http.Request) bool {
	requestPath := request.URL.Path
	noAuth := []string{"/", "/create", "/api/createAccount", "/login"}
	//check if response does not require authenthication
	for _, value := range noAuth {
		if value == requestPath || strings.Contains(requestPath, ".chunk.js") || strings.Contains(requestPath, ".chunk.css") || strings.Contains(requestPath, "manifest.json") || strings.Contains(requestPath, "favicon.ico") || strings.Contains(requestPath, ".png") {
			return true
		}
	}
	//other wise it requires authentication
	response := make(map[string]interface{})
	// tokenHeader := request.Header.Get("Authorization")
	tokenCookie, err := request.Cookie("x-wuhan-cookie")
	tokenHeader := tokenCookie.Value

	if err != nil {
		response = Message(false, "Missing auth cookie")
		writer.WriteHeader(http.StatusForbidden)
		writer.Header().Add("Content-Type", "application/json")
		Respond(writer, response)
		return false
	}
	if tokenHeader == "" {
		response = Message(false, "Missing auth token")
		writer.WriteHeader(http.StatusForbidden)
		writer.Header().Add("Content-Type", "application/json")
		Respond(writer, response)
		return false
	}
	splitted := strings.Split(tokenHeader, "%20") //The token normally comes in format `Bearer {token-body}`, we check if the retrieved token matched this requirement
	if len(splitted) != 2 {
		response = Message(false, "Invalid/Malformed auth token")
		writer.WriteHeader(http.StatusForbidden)
		writer.Header().Add("Content-Type", "application/json")
		Respond(writer, response)
		return false
	}
	tokenPart := splitted[1] // the information that we're interested in
	// Validate the token
	tk := &Token{}
	token, err := jwt.ParseWithClaims(tokenPart, tk, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("token_password")), nil
	})
	//malformed token, return 403
	if err != nil {
		response = Message(false, "Malformed auth token")
		writer.WriteHeader(http.StatusForbidden)
		writer.Header().Add("Content-Type", "application/json")
		Respond(writer, response)
		return false
	}
	//token is invalid
	if !token.Valid {
		response = Message(false, "Token is invalid")
		writer.WriteHeader(http.StatusForbidden)
		writer.Header().Add("Content-Type", "application/json")
		Respond(writer, response)
		return false
	}
	// Expired token
	if tk.Exp < time.Now().Unix() {
		response = Message(false, "Token has expired")
		writer.WriteHeader(http.StatusForbidden)
		writer.Header().Add("Content-Type", "application/json")
		Respond(writer, response)
		return false
	}
	return true
}
