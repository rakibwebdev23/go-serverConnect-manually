package review

import (
	"ecommerce/database"
	"ecommerce/utils"
	"net/http"
)

// GET products api
func (h *Handler) GetReview(w http.ResponseWriter, r *http.Request) {
	utils.GetProductsResponse(w, http.StatusOK, "Success", len(database.GetList()), 1, 10, database.GetReviewList())
}
