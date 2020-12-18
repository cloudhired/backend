package main

import (
	"cloudhired.com/api/handler"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
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
	router.GlobalOPTIONS = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := w.Header()
		header.Set("Access-Control-Allow-Methods", "GET, POST, PUT, OPTIONS")
		header.Set("Access-Control-Allow-Origin", "*")
		header.Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		w.WriteHeader(http.StatusNoContent)
	})

	router.GET("/", Index)
	router.GET("/hello/:name", Hello)
	router.GET("/api/users", handler.HandleGetAllUsers)
	router.GET("/api/username/:username", handler.HandleGetUser)
	router.POST("/api/username/:username", handler.HandlePostOneUser)

	log.Fatal(http.ListenAndServe(":8081", router))
}
