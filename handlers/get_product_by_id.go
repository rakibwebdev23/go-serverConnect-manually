package handlers

import (
	"ecommerce/database"
	"ecommerce/utils"
	"net/http"
	"strconv"
)

func GetProductByID(w http.ResponseWriter, r *http.Request) {
	productID := r.PathValue("productId")

	id, err := strconv.Atoi(productID)
	if err != nil {
		utils.HandleError(w, http.StatusBadRequest, "Invalid product ID")
		return
	}

	for _, product := range database.ProductList {
		if product.ID == id {
			utils.HandleSend(w, http.StatusOK, product)
			return
		}
	}
	utils.HandleError(w, http.StatusNotFound, "Product not found")
}
