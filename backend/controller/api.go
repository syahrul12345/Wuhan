package controller

import (
	"backend/models"
	"backend/utils"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"math/big"
	"net/http"
	"os"
	"path/filepath"

	"github.com/syndtr/goleveldb/leveldb"
)

type Request struct {
	Message string `json:"message"`
}
type Payload struct {
	Country string `json:"Country"`
	Deaths  uint32 `json:"Deaths"`
}

// Serve will serve the frontend
var Serve = func(w http.ResponseWriter, r *http.Request) {
	// Deal with the authentication first
	authenticated := utils.Auth(w, r)
	if !authenticated {
		return
	}
	const staticPath = "./build"
	const indexPath = "index.html"
	fileServer := http.FileServer(http.Dir(staticPath))
	path, err := filepath.Abs(r.URL.Path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	path = filepath.Join(staticPath, path)
	_, err = os.Stat(path)
	if os.IsNotExist(err) {
		// file does not exist, serve index.html
		http.ServeFile(w, r, filepath.Join(staticPath, indexPath))
		return
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fileServer.ServeHTTP(w, r)
}
var Save = func(w http.ResponseWriter, r *http.Request) {
	resp := make(map[string]interface{})
	request := &Request{}
	json.NewDecoder(r.Body).Decode(request)
	db, _ := leveldb.OpenFile("./db", nil)
	defer db.Close()

	data, err := db.Get([]byte("counts"), nil)
	if err != nil {
		// Does not exist..
		fmt.Println("First time run... no clicks")
		db.Put([]byte("counts"), []byte{0}, nil)
	}
	// Use big, in case it becoems big
	z := big.NewInt(0).Add(big.NewInt(0).SetBytes(data), big.NewInt(1))
	db.Put([]byte("counts"), z.Bytes(), nil)
	resp["count"] = z.Text(10)
	w.WriteHeader(http.StatusOK)
	utils.Respond(w, resp)
}
var Get = func(w http.ResponseWriter, r *http.Request) {
	resp := make(map[string]interface{})
	db, _ := leveldb.OpenFile("./db", nil)
	defer db.Close()
	// Get prayer count
	data, err := db.Get([]byte("counts"), nil)
	if err != nil {
		// Does not exist..
		fmt.Println("No Clicks stored")
	}
	z := big.NewInt(0).SetBytes(data)
	resp["count"] = z.Text(10)
	// get death count
	data, err = db.Get([]byte("deaths"), nil)
	if err != nil {
		// Does not exist...
		fmt.Println("deaths count first initiated")
		db.Put([]byte("deaths"), []byte{0}, nil)
	}
	z = big.NewInt(0).SetBytes(data)
	resp["death"] = z.Text(10)
	w.WriteHeader(http.StatusOK)
	utils.Respond(w, resp)
}

// UpdateDeath updates the death count
var UpdateDeath = func(w http.ResponseWriter, r *http.Request) {
	resp := make(map[string]interface{})

	tempPayload := &Payload{}
	json.NewDecoder(r.Body).Decode(tempPayload)
	// Get the level db
	db, _ := leveldb.OpenFile("./db", nil)
	defer db.Close()
	if tempPayload.Country == "global" {
		fmt.Println("uploading global...")
		buf := make([]byte, 4)
		binary.BigEndian.PutUint32(buf, tempPayload.Deaths)
		db.Put([]byte("deaths"), buf, nil)
	} else {
		buf := make([]byte, 4)
		fmt.Println("updating others.....")
		binary.BigEndian.PutUint32(buf, tempPayload.Deaths)
		db.Put([]byte(tempPayload.Country), buf, nil)
	}
	resp["message"] = "sucess in updating death count"
	w.WriteHeader(http.StatusOK)
	utils.Respond(w, resp)
}

// GetTotalDeath is a comment
var GetTotalDeath = func(w http.ResponseWriter, r *http.Request) {
	resp := make(map[string]interface{})
	db, _ := leveldb.OpenFile("./db", nil)
	data, err := db.Get([]byte("deaths"), nil)
	defer db.Close()
	if err != nil {
		fmt.Println("Cant get total death count... intiializing it to 0")
		db.Put([]byte("deaths"), []byte{0}, nil)
	}
	z := big.NewInt(0).SetBytes(data)
	resp["death"] = z.Text(10)
	w.WriteHeader(http.StatusOK)
	utils.Respond(w, resp)
}

// CreateAccount will create a new account
var CreateAccount = func(w http.ResponseWriter, r *http.Request) {
	newAcc := &models.NewAccount{}
	json.NewDecoder(r.Body).Decode(newAcc)
	if newAcc.Password != newAcc.ConfirmPassword {
		w.WriteHeader(400)
		utils.Respond(w, utils.Message(false, "Password not the same!"))
		return
	}
	acc := &models.Account{
		Email:    newAcc.Email,
		Password: newAcc.Password,
	}
	ok, message := acc.Validate()
	// Fail validation
	if !ok {
		w.WriteHeader(400)
		utils.Respond(w, utils.Message(false, message))
		return
	}
	// Create account
	ok, message = acc.Create()
	if !ok {
		w.WriteHeader(400)
		utils.Respond(w, utils.Message(false, message))
		return
	}
	// return the token as an authorization header as well as a string
	w.Header().Add("Authorization", message)
	utils.Respond(w, map[string]interface{}{
		"status": true,
		"token":  message,
	})
}

// Login
var Login = func(w http.ResponseWriter, r *http.Request) {
	acc := &models.Account{}
	json.NewDecoder(r.Body).Decode(acc)
	// Compare the password
	ok, message := acc.Login()
	if !ok {
		utils.Respond(w, utils.Message(false, message))
		return
	}
	// return the token as an authorization header as well as a string
	w.Header().Add("Authorization", message)
	utils.Respond(w, map[string]interface{}{
		"status": true,
		"token":  message,
	})
}
