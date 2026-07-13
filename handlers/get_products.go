package handlers

import ( 
	"net/http"
	"ecommerce/product"
	"ecommerce/utils"
)

// GET products api
func GetProducts(w http.ResponseWriter, r *http.Request) {

	response := product.APIResponse{
		StatusCode: http.StatusOK,
		Success:    true,
		Message:    "Success",
		Data: product.ProductData{
			Meta: product.Meta{
				Total:      len(product.ProductList),
				Page:       1,
				Limit:      10,
				TotalPages: 1,
			},
			Result: product.ProductList,
		},
	}

	utils.HandleSend(w, response)
}