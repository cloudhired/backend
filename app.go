package main

import (
	"fmt"
    "net/http"
	"log"
	"encoding/json"

	"github.com/julienschmidt/httprouter"
	"cloudhired.com/api/dao"
	// "cloudhired.com/api/models"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    fmt.Fprint(w, "Welcome!\n")
}

func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
}

// func AllUsers(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
// 	users, err := dao.FindAllUsers()
// 	if err != nil {
// 		return
// 	}
// 	respondWithJson(w, http.StatusOK, users)

// }

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func init() {

	// dao.Server = "localhost"
	// dao.Database = "users"
	// dao.Connect()
}

func main() {
    router := httprouter.New()
    router.GET("/", Index)
    router.GET("/hello/:name", Hello)
	// router.GET("/users", AllUsers)

	dao.TestDb()

    log.Fatal(http.ListenAndServe(":8081", router))
}