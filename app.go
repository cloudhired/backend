package main

import (
	"fmt"
    "net/http"
	"log"
	"encoding/json"

	"github.com/julienschmidt/httprouter"
	"./dao"
	"./models"
)

var dao = UsersDao{}

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    fmt.Fprint(w, "Welcome!\n")
}

func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
}

func AllUsers(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	users, err := dao.FindAllUsers()
	if err != nil {
		return
	}
	respondWithJson(w, http.StatusOK, users)

}

func main() {
    router := httprouter.New()
    router.GET("/", Index)
    router.GET("/hello/:name", Hello)
	router.GET("/users", AllUsers)

    log.Fatal(http.ListenAndServe(":8081", router))
}