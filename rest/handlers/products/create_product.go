package products

import (
	"ecommerce/repo"
	"ecommerce/utils"
	"encoding/json"
	"fmt"
	"net/http"
)

type ReqCreateProduct struct {
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImgUrl      string  `json:"img_url"`
}


// POST create api
func (h *Handler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	// json struct a convert
	var createProduct ReqCreateProduct
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&createProduct)
	if err != nil {
		fmt.Println("Error decoding request body:", err)
		utils.HandleError(w, http.StatusBadRequest, "Invalid request body")
		return
	}
	createdProduct, err := h.productRepo.Create(repo.Product{
		Title: createProduct.Title,
		Description: createProduct.Description,
		Price: createProduct.Price,
		ImgUrl: createProduct.ImgUrl,
	})
	if err != nil {
		utils.HandleError(w, http.StatusInternalServerError, "Failed to create product")
		return
	}
	utils.CreateResponse(w, http.StatusCreated, "Product created successfully", createdProduct)

}
