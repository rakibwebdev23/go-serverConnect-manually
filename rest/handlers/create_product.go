package handlers

import (
	"ecommerce/database"
	"ecommerce/utils"
	"encoding/json"
	"fmt"
	"net/http"
)

// POST create api
func CreateProduct(w http.ResponseWriter, r *http.Request) {
	// json struct a convert
	var newProduct database.Product
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newProduct)
	if err != nil {
		fmt.Println("Error decoding request body:", err)
		utils.HandleError(w, http.StatusBadRequest, "Invalid request body")
		return
	}
	createdProduct := database.Store(newProduct)
	utils.CreateResponse(w, http.StatusCreated, "Product created successfully", createdProduct)

}
