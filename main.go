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
	handleCors(w);
	handlePreflightOptionsReq(w, r);

	if r.Method != "GET"{
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

	handleSend(w, response)
}

// GET api request with options & preflight request for products api
func getProductsWithOptions(w http.ResponseWriter, r *http.Request) {
	handleCors(w)
	handlePreflightOptionsReq(w, r);

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

// POST create api
func createProduct(w http.ResponseWriter, r *http.Request) {
	handleCors(w)
	handlePreflightOptionsReq(w, r)

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
	handleSend(w, response)

}

func handleCors(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Rakib")
	w.Header().Set("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,DELETE,OPTIONS")
}

func handlePreflightOptionsReq(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}
}

func handleSend(w http.ResponseWriter, data interface{}) {
	encoder := json.NewEncoder(w)
	encoder.Encode(data)
}

func main() {
	mux := http.NewServeMux() //router create

	mux.Handle("GET /hello", http.HandlerFunc(handleHello)) // route declaration
	mux.Handle("GET /good-bye", http.HandlerFunc(handleGoodbye))
	mux.Handle("GET /products", http.HandlerFunc(getProducts))
	mux.Handle("OPTIONS /products", http.HandlerFunc(getProductsWithOptions))
	mux.Handle("POST /products/create", http.HandlerFunc(createProduct))

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
