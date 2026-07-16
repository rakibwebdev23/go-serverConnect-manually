package handlers

import (
	"net/http"
	"ecommerce/database"
	"ecommerce/utils"
)

// GET products api
func GetProducts(w http.ResponseWriter, r *http.Request) {
	utils.GetProductsResponse(w, http.StatusOK, "Success", len(database.GetList()), 1, 10, database.GetList())
}
