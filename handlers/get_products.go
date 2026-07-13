package main

import "net/http"

// GET products api
func getProducts(w http.ResponseWriter, r *http.Request) {

	response := APIResponse{
		StatusCode: http.StatusOK,
		Success:    true,
		Message:    "Success",
		Data: ProductData{
			Meta: Meta{
				Total:      len(productList),
				Page:       1,
				Limit:      10,
				TotalPages: 1,
			},
			Result: productList,
		},
	}

	handleSend(w, response)
}