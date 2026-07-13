package main

import (
	"fmt"
	"net/http"
	"ecommerce/global_router"
	"ecommerce/handlers"
	"ecommerce/product"
)

func main() {
	mux := http.NewServeMux() //router create

	mux.Handle("GET /products", http.HandlerFunc(handlers.GetProducts))
	mux.Handle("POST /products/create", http.HandlerFunc(handlers.CreateProduct))

	fmt.Println("Server port running:8080")

	err := http.ListenAndServe(":8080", global_router.HandleGlobarlRouter(mux))
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}

func init() {
	product1 := product.Product{
		ID:          1,
		Title:       "Banana",
		Description: "A delicious yellow fruit",
		Price:       1.99,
		ImgUrl:      "https://example.com/banana.jpg",
	}
	product2 := product.Product{
		ID:          2,
		Title:       "Apple",
		Description: "A sweet red fruit",
		Price:       0.99,
		ImgUrl:      "https://example.com/apple.jpg",
	}
	product3 := product.Product{
		ID:          3,
		Title:       "Orange",
		Description: "A juicy citrus fruit",
		Price:       1.49,
		ImgUrl:      "https://example.com/orange.jpg",
	}
	product4 := product.Product{
		ID:          4,
		Title:       "Grapes",
		Description: "A bunch of small round fruits",
		Price:       2.99,
		ImgUrl:      "https://example.com/grapes.jpg",
	}
	product5 := product.Product{
		ID:          5,
		Title:       "Strawberry",
		Description: "A red heart-shaped fruit",
		Price:       3.49,
		ImgUrl:      "https://example.com/strawberry.jpg",
	}
	product6 := product.Product{
		ID:          6,
		Title:       "Pineapple",
		Description: "A tropical fruit with a spiky exterior",
		Price:       4.99,
		ImgUrl:      "https://example.com/pineapple.jpg",
	}
	product7 := product.Product{
		ID:          7,
		Title:       "Watermelon",
		Description: "A large green fruit with red flesh",
		Price:       5.99,
		ImgUrl:      "https://example.com/watermelon.jpg",
	}
	product8 := product.Product{
		ID:          8,
		Title:       "Mango",
		Description: "A sweet tropical fruit with orange flesh",
		Price:       2.49,
		ImgUrl:      "https://example.com/mango.jpg",
	}
	product9 := product.Product{
		ID:          9,
		Title:       "Blueberry",
		Description: "A small round blue fruit",
		Price:       3.99,
		ImgUrl:      "https://example.com/blueberry.jpg",
	}
	product10 := product.Product{
		ID:          10,
		Title:       "Kiwi",
		Description: "A small brown fruit with green flesh",
		Price:       1.49,
		ImgUrl:      "https://example.com/kiwi.jpg",
	}

	product.ProductList = append(product.ProductList, product1, product2, product3, product4, product5, product6, product7, product8, product9, product10)
}
