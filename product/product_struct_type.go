package product

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