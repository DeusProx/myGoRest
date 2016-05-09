package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	log.Printf("Start Server")

	router := httprouter.New()
	router.GET("/", index)
	router.GET("/user/:username", user)
	router.POST("/login", basicAuth(login))

	log.Fatal(http.ListenAndServe(":8080", router))
}

// curl -v -X GET 'http://localhost:8080/'
func index(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	fmt.Fprint(res, "Welcome!\n")
}

// curl -v -X GET 'http://localhost:8080/user/userX'
func user(res http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	fmt.Fprintf(res, "Welcome %s!\n", ps.ByName("username"))
}

// curl -v -X POST --data '{"username": "userX", "password": "secret"}' 'http://localhost:8080/login'
func login(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	res.WriteHeader(http.StatusCreated)
	fmt.Fprint(res, "LOGGED IN!\n")
}
