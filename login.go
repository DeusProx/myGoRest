package main

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Auth contains the data for the authentication method
type Auth struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func basicAuth(handler httprouter.Handle) httprouter.Handle {
	return func(res http.ResponseWriter, req *http.Request, ps httprouter.Params) {
		decoder := json.NewDecoder(req.Body)
		var auth Auth
		if err := decoder.Decode(&auth); err != nil {
			unauthorized(res, err)
			return
		}

		if auth.Password != "secret" {
			unauthorized(res, errors.New("wrong password"))
			return
		}

		sEnc := base64.StdEncoding.EncodeToString([]byte(auth.Username))
		res.Header().Set("Token", sEnc)
		handler(res, req, ps)
		return
	}
}

func tokenHandler(handler httprouter.Handle) httprouter.Handle {
	return func(res http.ResponseWriter, req *http.Request, ps httprouter.Params) {
		var token string
		if token = req.Header.Get("Token"); token == "" {
			unauthorized(res, errors.New("No Token"))
			return
		}
		if token != "testToken" {
			unauthorized(res, errors.New("Wrong Token"))
			return
		}

		handler(res, req, ps)
		return
	}
}

func unauthorized(res http.ResponseWriter, err error) {
	log.Println("Error: ", err)
	http.Error(res, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
	return
}

// curl -v -X POST --data '{"username": "userX", "password": "secret"}' 'http://localhost:8080/login'
func login(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	res.WriteHeader(http.StatusCreated)
	fmt.Fprint(res, "LOGGED IN!\n")
}
