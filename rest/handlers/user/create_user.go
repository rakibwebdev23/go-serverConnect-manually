package user

import (
	"ecommerce/database"
	"ecommerce/utils"
	"encoding/json"
	"fmt"
	"net/http"
)

// POST create api
func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	// json struct a convert
	var newUser database.User
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newUser)
	if err != nil {
		fmt.Println("Error decoding request body:", err)
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	createdUser := database.UserStore(newUser)

	utils.CreateResponse(w, http.StatusCreated, "User Registered successfully", createdUser)
}
