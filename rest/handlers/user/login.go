package user

import (
	"ecommerce/utils"
	"encoding/json"
	"fmt"
	"net/http"
)

type ReqLogin struct {
	Email    string `json:"email"`
	// Password string `json:"password"`
	Password    string `json:"_"`
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

	loginUser, err := h.userRepo.Find(reqLogin.Email, reqLogin.Password)
	if err != nil {
		fmt.Println("Database error during login:", err)
		utils.HandleError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}
	if loginUser == nil {
		utils.HandleError(w, http.StatusBadRequest, "Invalid credentials")
		return
	}

	accessToken, err := utils.CreateJwt(h.cnf.JwtSecretKey, UserPayload{
		User:      loginUser.ID,
		FirstName: loginUser.FirstName,
		LastName:  loginUser.LastName,
		Email:     loginUser.Email,
	})
	if err != nil {
		utils.HandleError(w, http.StatusInternalServerError, "Failed to create token")
		return
	}

	response := map[string]interface{}{
		"accessToken": accessToken,
	}

	utils.CreateResponse(w, http.StatusCreated, "Login successfully", response)
}
