package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"ecommerce/product"
	"ecommerce/utils"
)

// POST create api
func CreateProduct(w http.ResponseWriter, r *http.Request) {
	// json struct a convert 
	var newProduct product.Product
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newProduct)
	if err != nil {
		fmt.Println("Error decoding request body:", err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	newProduct.ID = len(product.ProductList) + 1
	product.ProductList = append(product.ProductList, newProduct)

	response := product.APIResponseCreate{
		StatusCode: http.StatusCreated,
		Success:    true,
		Message:    "Product created successfully",
		Data:       newProduct,
	}
	
	utils.HandleSend(w, response);

}
