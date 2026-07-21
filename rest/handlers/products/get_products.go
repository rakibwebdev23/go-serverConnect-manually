package products

import (
	"net/http"
	"ecommerce/database"
	"ecommerce/utils"
)

// GET products api
func (h *Handler) GetProducts(w http.ResponseWriter, r *http.Request) {
	utils.GetProductsResponse(w, http.StatusOK, "Success", len(database.GetList()), 1, 10, database.GetList())
}
