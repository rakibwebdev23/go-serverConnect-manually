package user

import (
	"ecommerce/repo"
	"ecommerce/utils"
	"encoding/json"
	"fmt"
	"net/http"
)

type ReqRegister struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	IsShopOwner bool   `json:"is_shop_owner"`
}

// POST create api
func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	// json struct a convert
	var reqRegister ReqRegister
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&reqRegister)
	if err != nil {
		fmt.Println("Error decoding request body:", err)
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	createdUser, err := h.userRepo.Create(repo.User{
		FirstName:   reqRegister.FirstName,
		LastName:    reqRegister.LastName,
		Email:       reqRegister.Email,
		Password:    reqRegister.Password,
		IsShopOwner: reqRegister.IsShopOwner,
	})
	if err != nil {
		utils.HandleError(w, http.StatusInternalServerError, "Failed to register user")
		return
	}

	accessToken, err := utils.CreateJwt(h.cnf.JwtSecretKey, UserPayload{
		User:      createdUser.ID,
		FirstName: createdUser.FirstName,
		LastName:  createdUser.LastName,
		Email:     createdUser.Email,
	})
	if err != nil {
		utils.HandleError(w, http.StatusInternalServerError, "Failed to create token")
		return
	}

	response := map[string]interface{}{
		"accessToken": accessToken,
		"user":        createdUser,
	}

	utils.CreateResponse(w, http.StatusCreated, "User Registered successfully", response)
}

