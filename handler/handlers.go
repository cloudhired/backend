package handler

import (
	"cloudhired.com/api/dao"
	"cloudhired.com/api/models"
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

var users []models.User

func HandleGetUsername(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if tokenString := r.Header.Get("x-auth-token"); len(tokenString) > 0 {
		fmt.Println(tokenString)
		if token, _ := jwt.Parse(tokenString, nil); token != nil {
			claims, _ := token.Claims.(jwt.MapClaims)
			uid, _ := claims["user_id"].(string)
			name, _ := claims["name"].(string)
			email, _ := claims["email"].(string)
			emailVerified, _ := claims["email_verified"].(bool)
			if len(uid) > 0 {
				m := make(map[string]string)
				m["username"] = dao.GetUsernameByUid(uid, name, email, emailVerified)
				json.NewEncoder(w).Encode(m)
			} else {
				fmt.Println("UID is empty, will not proceed")
			}
		} else {
			fmt.Println("token is empty")
		}
	} else {
		fmt.Println("auth header is empty, not authorized.")
	}
}

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
