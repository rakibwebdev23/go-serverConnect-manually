package user

import (
	"ecommerce/database"
	"ecommerce/utils"
	"encoding/json"
	"fmt"
	"net/http"
)

type ReqLogin struct{
	Email string `json:"email"`
	Password string `json:"password"`
}

// POST create api
func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var reqLogin ReqLogin

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&reqLogin)
	if err != nil {
		fmt.Println("Error decoding request body:", err)
		http.Error(w, "Invalid request data", http.StatusBadRequest)
		return
	}

	loginUser := database.FindUser(reqLogin.Email, reqLogin.Password)
	if loginUser == nil {
		http.Error(w, "Invalid credentials", http.StatusBadRequest)
		return
	}
	
	utils.CreateResponse(w, http.StatusCreated, "Login successfull", loginUser)

}
