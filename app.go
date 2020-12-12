package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"cloudhired.com/api/handler"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    fmt.Fprint(w, "Welcome!\n")
}

func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
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
	router.GET("/users", handler.HandleGetAllUsers)

    log.Fatal(http.ListenAndServe(":8081", router))
}