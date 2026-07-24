package products

import (
	"net/http"
	"ecommerce/utils"
)

// GET products api
func (h *Handler) GetProducts(w http.ResponseWriter, r *http.Request) {
	products, err := h.productRepo.List()
	if err != nil {
		utils.HandleError(w, http.StatusInternalServerError, "Failed to fetch products")
		return
	}
	utils.GetProductsResponse(w, http.StatusOK, "Success", len(products), 1, 10, products)
}
