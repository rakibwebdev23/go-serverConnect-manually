package products

import (
	"ecommerce/repo"
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

	product, err := h.productRepo.Get(id)
	if err != nil {
		if err == repo.ErrNotFound {
			utils.HandleError(w, http.StatusNotFound, "Product not found")
			return
		}
		utils.HandleError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}
	utils.GetSingleProductResponse(w, http.StatusOK, "Success", product)
}
