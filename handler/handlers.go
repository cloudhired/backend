package handler

import (
	"cloudhired.com/api/dao"
	"cloudhired.com/api/models"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

var users []models.User

func HandleGetAllUsers(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	payload := dao.AllUsers()
	m := make(map[string][]models.User)
	m["results"] = payload
	json.NewEncoder(w).Encode(m)
}

func HandleGetUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	m := make(map[string]models.User)
	m["data"] = dao.FindOneUser(ps.ByName("username"))
	json.NewEncoder(w).Encode(m)
}

func HandlePostOneUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	username := ps.ByName("username")
	var user interface{}
	var res = make(map[string]interface{})
	json.NewDecoder(r.Body).Decode(&user)
	fmt.Println(user)
	var isGood = dao.UpdateOneUser(username, user)
	if isGood {
		res["error"] = nil
	} else {
		res["error"] = true
	}
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	json.NewEncoder(w).Encode(res)
}
