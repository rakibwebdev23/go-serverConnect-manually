package review

import (
	"ecommerce/utils"
	"net/http"
)

// GET reviews api
func (h *Handler) GetReview(w http.ResponseWriter, r *http.Request) {
	reviews, err := h.reviewRepo.List()
	if err != nil {
		utils.HandleError(w, http.StatusInternalServerError, "Failed to fetch reviews")
		return
	}
	utils.GetProductsResponse(w, http.StatusOK, "Success", len(reviews), 1, 10, reviews)
}

