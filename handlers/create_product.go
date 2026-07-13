package handlers

import (
	"ecommerce/database"
	"ecommerce/utils"
	"encoding/json"
	"fmt"
	"net/http"
)

// POST create api
func CreateProduct(w http.ResponseWriter, r *http.Request) {
	// json struct a convert 
	var newProduct database.Product
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newProduct)
	if err != nil {
		fmt.Println("Error decoding request body:", err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	newProduct.ID = len(database.ProductList) + 1
	database.ProductList = append(database.ProductList, newProduct)

	response := database.APIResponseCreate{
		StatusCode: http.StatusCreated,
		Success:    true,
		Message:    "Product created successfully",
		Data:       newProduct,
	}
	
	utils.HandleSend(w, response);

}
