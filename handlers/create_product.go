package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// POST create api
func createProduct(w http.ResponseWriter, r *http.Request) {
	// json struct a convert 
	var newProduct Product
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newProduct)
	if err != nil {
		fmt.Println("Error decoding request body:", err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	newProduct.ID = len(productList) + 1
	productList = append(productList, newProduct)

	response := APIResponseCreate{
		StatusCode: http.StatusCreated,
		Success:    true,
		Message:    "Product created successfully",
		Data:       newProduct,
	}
	
	handleSend(w, response)

}
