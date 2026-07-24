package products

import (
	"ecommerce/repo"
	"ecommerce/utils"
	"encoding/json"
	"net/http"
	"strconv"
)

func (h *Handler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	// Get id from request url params
	productID := r.PathValue("id")

	id, err := strconv.Atoi(productID)
	if err != nil {
		utils.HandleError(w, http.StatusBadRequest, "Invalid product ID")
		return
	}

	var newProduct repo.Product
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&newProduct)
	if err != nil {
		utils.HandleError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	newProduct.ID = id
	updatedProduct, err := h.productRepo.Update(newProduct)
	if err != nil {
		utils.HandleError(w, http.StatusInternalServerError, "Failed to update product")
		return
	}

	utils.UpdateResponse(w, http.StatusOK, "Product updated successfully", updatedProduct)
}

