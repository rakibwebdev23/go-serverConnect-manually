package cmd

import (
	"ecommerce/global_router"
	"ecommerce/handlers"
	"fmt"
	"net/http"
)

func Serve() {
	mux := http.NewServeMux() //router create

	mux.Handle("GET /products", http.HandlerFunc(handlers.GetProducts))
	mux.Handle("GET /products/{productId}", http.HandlerFunc(handlers.GetProductByID));
	mux.Handle("POST /products", http.HandlerFunc(handlers.CreateProduct))

	fmt.Println("Server port running:8080")

	err := http.ListenAndServe(":8080", global_router.HandleGlobarlRouter(mux))
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
