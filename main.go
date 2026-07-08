package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Product struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImgUrl      string  `json:"img_url"`
}

type Meta struct {
	Total      int `json:"total"`
	Page       int `json:"page"`
	Limit      int `json:"limit"`
	TotalPages int `json:"totalPages"`
}

type ProductData struct {
	Meta   Meta      `json:"meta"`
	Result []Product `json:"result"`
}

type APIResponse struct {
	StatusCode int         `json:"statusCode"`
	Success    bool        `json:"success"`
	Message    string      `json:"message"`
	Data       ProductData `json:"data"`
}

// Response for Create Product
type APIResponseCreate struct {
	StatusCode int     `json:"statusCode"`
	Success    bool    `json:"success"`
	Message    string  `json:"message"`
	Data       Product `json:"data"`
}

var productList []Product

func handleHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello world")
}

func handleGoodbye(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Goodbye world")
}

// GET products api
func getProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	if r.Method != "GET" {
		http.Error(w, "Please send GET request", http.StatusBadRequest)
		return
	}

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

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

// GET api request with options & preflight request for products api
func getProductsWithOptions(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Rakib")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	if r.Method == "OPTIONS" {
		w.WriteHeader(200)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Please send GET request", http.StatusBadRequest)
		return
	}

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

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

// POST create api
func createProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != "POST" {
		http.Error(w, "Please send POST request", http.StatusBadRequest)
		return
	}

	var newProduct Product
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newProduct)
	if err != nil {
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

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)

}

func main() {
	mux := http.NewServeMux() //router create

	mux.HandleFunc("/hello", handleHello) // route declaration
	mux.HandleFunc("/good-bye", handleGoodbye)
	mux.HandleFunc("/products", getProductsWithOptions)
	mux.HandleFunc("/products/create", createProduct)

	fmt.Println("Server port running:8080")

	err := http.ListenAndServe(":8080", mux)

	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}

func init() {
	product1 := Product{
		ID:          1,
		Title:       "Banana",
		Description: "A delicious yellow fruit",
		Price:       1.99,
		ImgUrl:      "https://example.com/banana.jpg",
	}
	product2 := Product{
		ID:          2,
		Title:       "Apple",
		Description: "A sweet red fruit",
		Price:       0.99,
		ImgUrl:      "https://example.com/apple.jpg",
	}
	product3 := Product{
		ID:          3,
		Title:       "Orange",
		Description: "A juicy citrus fruit",
		Price:       1.49,
		ImgUrl:      "https://example.com/orange.jpg",
	}
	product4 := Product{
		ID:          4,
		Title:       "Grapes",
		Description: "A bunch of small round fruits",
		Price:       2.99,
		ImgUrl:      "https://example.com/grapes.jpg",
	}
	product5 := Product{
		ID:          5,
		Title:       "Strawberry",
		Description: "A red heart-shaped fruit",
		Price:       3.49,
		ImgUrl:      "https://example.com/strawberry.jpg",
	}
	product6 := Product{
		ID:          6,
		Title:       "Pineapple",
		Description: "A tropical fruit with a spiky exterior",
		Price:       4.99,
		ImgUrl:      "https://example.com/pineapple.jpg",
	}
	product7 := Product{
		ID:          7,
		Title:       "Watermelon",
		Description: "A large green fruit with red flesh",
		Price:       5.99,
		ImgUrl:      "https://example.com/watermelon.jpg",
	}
	product8 := Product{
		ID:          8,
		Title:       "Mango",
		Description: "A sweet tropical fruit with orange flesh",
		Price:       2.49,
		ImgUrl:      "https://example.com/mango.jpg",
	}
	product9 := Product{
		ID:          9,
		Title:       "Blueberry",
		Description: "A small round blue fruit",
		Price:       3.99,
		ImgUrl:      "https://example.com/blueberry.jpg",
	}
	product10 := Product{
		ID:          10,
		Title:       "Kiwi",
		Description: "A small brown fruit with green flesh",
		Price:       1.49,
		ImgUrl:      "https://example.com/kiwi.jpg",
	}

	productList = append(productList, product1, product2, product3, product4, product5, product6, product7, product8, product9, product10)
}
