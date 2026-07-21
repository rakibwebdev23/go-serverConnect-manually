package products

import (
	"ecommerce/database"
	"ecommerce/utils"
	"net/http"
	"strconv"
)

func (h *Handler) GetProductByID(w http.ResponseWriter, r *http.Request) {
	productID := r.PathValue("id")

	id, err := strconv.Atoi(productID)
	if err != nil {
		utils.HandleError(w, http.StatusBadRequest, "Invalid product ID")
		return
	}

	product := database.SingleGet(id)
	if product == nil {
		utils.HandleError(w, http.StatusNotFound, "Product not found")
		return
	}
	utils.GetSingleProductResponse(w, http.StatusOK, "Success", product)
}
