package models

import (
	"backend/utils"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/syndtr/goleveldb/leveldb"
	"golang.org/x/crypto/bcrypt"
)

type Account struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type NewAccount struct {
	Email           string `json:"email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirmPassword"`
}

// Validate an account send to bevalid
func (account *Account) Validate() (bool, string) {
	if account.Email == "" {
		return false, "Email cannot be empty"
	}
	if account.Password == "" {
		return false, "Password cannot be empty"
	}
	if len(account.Password) <= 6 {
		return false, "Password must be more than  6 characters"
	}
	return true, ""
}

// Create a new account
func (account *Account) Create() (bool, string) {
	// Let's find if the user name already exists in the database:
	db, _ := leveldb.OpenFile("./db", nil)
	defer db.Close()
	_, err := db.Get([]byte(account.Email), nil)
	// Getting it must be EMPTY!
	if err != leveldb.ErrNotFound {
		// Some other error occured
		return false, "This account already exists"
	}
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(account.Password), bcrypt.DefaultCost)
	db.Put([]byte(account.Email), hashedPassword, nil)
	// Create a new jwt token as well
	tk := &utils.Token{Email: account.Email, Exp: getExpiryDate()}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))
	return true, tokenString
}

// Login using the account credentials
func (account *Account) Login() (bool, string) {
	// Check fi account exists
	db, _ := leveldb.OpenFile("./db", nil)
	defer db.Close()
	hashedPassword, err := db.Get([]byte(account.Email), nil)
	if err != nil {
		// Some error occured
		return false, "No such account"
	}
	err = bcrypt.CompareHashAndPassword(hashedPassword, []byte(account.Password))
	if err != nil {
		return false, err.Error()
	}
	// Generate a jwt token for this new login
	tk := &utils.Token{Email: account.Email, Exp: getExpiryDate()}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))
	return true, tokenString
}

func getExpiryDate() int64 {
	start := time.Now()
	end := start.Add(time.Second * 60)
	return end.Unix()
}
