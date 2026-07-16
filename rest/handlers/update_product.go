package handlers

import (
	"ecommerce/database"
	"ecommerce/utils"
	"encoding/json"
	"net/http"
	"strconv"
)

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	// Get id from request url params
	productID := r.PathValue("id")

	id, err := strconv.Atoi(productID)
	if err != nil {
		utils.HandleError(w, http.StatusBadRequest, "Invalid product ID")
		return
	}

	var newProduct database.Product
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&newProduct)
	if err != nil {
		utils.HandleError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	newProduct.ID = id
	database.Update(newProduct)

	utils.UpdateResponse(w, http.StatusOK, "Product updated successfully", newProduct)

}
