package main

import (
	"cloudhired.com/api/handler"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"strings"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	//fmt.Println(strings.Split(r.Header.Get("Authorization"), " ")[1])
	fmt.Fprint(w, strings.Split(r.Header.Get("Authorization"), " ")[1])
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
	// cloud endpoint has cors setting so do not need to configure here. set to allow all.
	router.GlobalOPTIONS = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := w.Header()
		header.Set("Access-Control-Allow-Methods", "GET, POST, PUT, OPTIONS")
		header.Set("Access-Control-Allow-Origin", "*")
		header.Set("Access-Control-Allow-Headers", "*")
		w.WriteHeader(http.StatusNoContent)
	})

	router.GET("/testing", Index)
	router.GET("/hello/:name", Hello)
	router.GET("/api/uidtousername", handler.HandleGetUsername)
	router.GET("/api/users", handler.HandleGetAllUsers)
	router.GET("/api/username/:username", handler.HandleGetUser)
	router.GET("/api/allcerts", handler.HandleGetAllCerts)
	router.POST("/api/username/:username", handler.HandlePostOneUser)

	log.Fatal(http.ListenAndServe(":8080", router))
}
