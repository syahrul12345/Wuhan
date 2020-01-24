package controller

import (
	_ "backend/models"
	"backend/utils"
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

// Serve will serve the frontend
var Serve = func(w http.ResponseWriter, r *http.Request) {
	const staticPath = "../website/build"
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
	data, err := db.Get([]byte("counts"), nil)
	if err != nil {
		// Does not exist..
		fmt.Println("No Clicks stored")
	}
	z := big.NewInt(0).SetBytes(data)
	resp["count"] = z.Text(10)
	w.WriteHeader(http.StatusOK)
	utils.Respond(w, resp)
}