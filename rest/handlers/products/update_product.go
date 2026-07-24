package products

import (
	"ecommerce/repo"
	"ecommerce/utils"
	"encoding/json"
	"net/http"
	"strconv"
)

type ReqUpdateProduct struct {
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImgUrl      string  `json:"img_url"`
}

func (h *Handler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	// Get id from request url params
	productID := r.PathValue("id")

	id, err := strconv.Atoi(productID)
	if err != nil {
		utils.HandleError(w, http.StatusBadRequest, "Invalid product ID")
		return
	}

	var reqUpdate ReqUpdateProduct
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&reqUpdate)
	if err != nil {
		utils.HandleError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	updatedProduct, err := h.productRepo.Update(repo.Product{
		ID:          id,
		Title:       reqUpdate.Title,
		Description: reqUpdate.Description,
		Price:       reqUpdate.Price,
		ImgUrl:      reqUpdate.ImgUrl,
	})
	if err != nil {
		utils.HandleError(w, http.StatusInternalServerError, "Failed to update product")
		return
	}

	utils.UpdateResponse(w, http.StatusOK, "Product updated successfully", updatedProduct)
}

