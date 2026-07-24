package user

import (
	"ecommerce/repo"
	"ecommerce/utils"
	"encoding/json"
	"fmt"
	"net/http"
)

// POST create api
func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	// json struct a convert
	var newUser repo.User
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newUser)
	if err != nil {
		fmt.Println("Error decoding request body:", err)
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	createdUser, err := h.userRepo.Create(newUser)
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

