package handler

import (
	"cloudhired.com/api/dao"
	"cloudhired.com/api/models"
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

var users []models.User

func HandleGetAllUsers(w http.ResponseWriter, r *http.Request, _ httprouter.Params)  {
	payload := dao.AllUsers()
	m := make(map[string][]models.User)
	m["results"] = payload
	json.NewEncoder(w).Encode(m)
}
