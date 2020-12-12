package handler

import (
	"encoding/json"
	"net/http"
	"cloudhired.com/api/dao"
	"cloudhired.com/api/models"

	"github.com/julienschmidt/httprouter"
	)

var users []models.User

func HandleGetAllUsers(w http.ResponseWriter, r *http.Request, _ httprouter.Params)  {
	payload := dao.AllUsers()
	json.NewEncoder(w).Encode(payload)
}
