package main

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Auth is a cool type Yo!
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
		res.Header().Set("token", sEnc)
		handler(res, req, ps)
		return
	}
}

func unauthorized(res http.ResponseWriter, err error) {
	log.Println("Error: ", err)
	http.Error(res, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
	return
}
