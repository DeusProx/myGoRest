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

	router.NotFound = http.FileServer(http.Dir("./static"))
	router.GET("/root", index)
	router.GET("/user/:username", user)
	router.POST("/login", basicAuth(login))
	router.GET("/private", tokenHandler(private))
	router.GET("/websocket", websocketHandler)

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

// curl -v -X GET -H "Token: testToken" 'http://localhost:8080/private'
func private(res http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	fmt.Fprintf(res, "This resource is private and protected by a token\n")
}
