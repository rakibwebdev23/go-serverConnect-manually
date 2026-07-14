package handlers

import ( 
	"net/http"
	"ecommerce/database"
	"ecommerce/utils"
)

// GET products api
func GetProducts(w http.ResponseWriter, r *http.Request) {

	response := database.APIResponse{
		StatusCode: http.StatusOK,
		Success:    true,
		Message:    "Success",
		Data: database.ProductData{
			Meta: database.Meta{
				Total:      len(database.ProductList),
				Page:       1,
				Limit:      10,
				TotalPages: 1,
			},
			Result: database.ProductList,
		},
	}

	utils.HandleSend(w, http.StatusOK, response)
}