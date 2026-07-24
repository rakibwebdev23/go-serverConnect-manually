package products

import (
	"ecommerce/utils"
	"net/http"
	"strconv"
)

func (h *Handler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	// Get id from request url params
	productID := r.PathValue("id")

	id, err := strconv.Atoi(productID)
	if err != nil {
		utils.HandleError(w, http.StatusBadRequest, "Invalid product ID")
		return
	}

	err = h.productRepo.Delete(id)
	if err != nil {
		utils.HandleError(w, http.StatusInternalServerError, "Failed to delete product")
		return
	}
	utils.DeleteResponse(w, http.StatusOK, "Product deleted successfully")
}
