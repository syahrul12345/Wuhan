package models

import (
	"backend/utils"
	"encoding/binary"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
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

	db := GetAccountDB()
	defer db.Close()

	_, err := db.Get([]byte(account.Email), nil)
	if err != leveldb.ErrNotFound {
		// Some other error occured
		return false, "This account already exists"
	}

	// Search the postgres db
	tempAccount := &Account{}
	err = GetDB().Table("accounts").Where("email = ?", account.Email).First(tempAccount).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, "Error connecting to DB or this acocun"
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(account.Password), bcrypt.DefaultCost)
	account.Password = string(hashedPassword)
	// Put in the cache
	// In the cache, we append the result with 4 bytes representing the timestamp. This is important for the scheduler to know when to clear the cahce
	expiryBuf := make([]byte, 4)
	binary.LittleEndian.PutUint32(expiryBuf, uint32(getExpiryDate(60)))
	hashedPasswordWithTimeStamp := append(expiryBuf, []byte(hashedPassword)...)

	db.Put([]byte(account.Email), hashedPasswordWithTimeStamp, nil)
	GetDB().Create(account)

	// Create a new jwt token as well
	tk := &utils.Token{Email: account.Email, Exp: getExpiryDate(60)}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))
	return true, tokenString
}

// Login using the account credentials
func (account *Account) Login() (bool, string) {
	// Check fi account exists in acocuntCacheDB
	tempAccount := &Account{}
	db := GetAccountDB()
	defer db.Close()

	hashedPasswordWithTimeStamp, err := db.Get([]byte(account.Email), nil)
	if err != nil {
		// Get from postgres
		err = GetDB().Table("accounts").Where("email = ?", account.Email).First(tempAccount).Error
		if err != nil || err == gorm.ErrRecordNotFound {
			if err == gorm.ErrRecordNotFound {
				return false, "No such account"
			}
			return false, "Failed to connect to DB"
			// Doesnt exist in the DB as well
		}
		// exists in the main DB. Create a time stamp and append
		hashedPassword := []byte(tempAccount.Password)
		expiryBuf := make([]byte, 4)
		binary.LittleEndian.PutUint32(expiryBuf, uint32(getExpiryDate(60)))
		hashedPasswordWithTimeStamp = append(expiryBuf, []byte(hashedPassword)...)
		// Write into the cache
		db.Put([]byte(account.Email), hashedPasswordWithTimeStamp, nil)
	}

	// Remove the timestamp
	hashedPassword := hashedPasswordWithTimeStamp[4:]
	err = bcrypt.CompareHashAndPassword(hashedPassword, []byte(account.Password))
	if err != nil {
		return false, err.Error()
	}
	// Generate a jwt token for this new login
	tk := &utils.Token{Email: account.Email, Exp: getExpiryDate(60)}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))
	return true, tokenString
}

func getExpiryDate(multiplier int) int64 {
	start := time.Now()
	end := start.Add(time.Second * time.Duration(multiplier))
	return end.Unix()
}
