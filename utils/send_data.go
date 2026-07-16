package utils

import (
	"encoding/json"
	"net/http"
)

type APIErrorResponse struct {
	StatusCode int    `json:"statusCode"`
	Success    bool   `json:"success"`
	Message    string `json:"message"`
}

type Meta struct {
	Total      int `json:"total"`
	Page       int `json:"page"`
	Limit      int `json:"limit"`
	TotalPages int `json:"totalPages"`
}

type ListData struct {
	Meta   Meta        `json:"meta"`
	Result interface{} `json:"result"`
}

type ListResponse struct {
	StatusCode int      `json:"statusCode"`
	Success    bool     `json:"success"`
	Message    string   `json:"message"`
	Data       ListData `json:"data"`
}

type SingleResponse struct {
	StatusCode int         `json:"statusCode"`
	Success    bool        `json:"success"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
}

type EmptyResponse struct {
	StatusCode int    `json:"statusCode"`
	Success    bool   `json:"success"`
	Message    string `json:"message"`
}

func HandleSend(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	encoder := json.NewEncoder(w)
	encoder.Encode(data)
}

func HandleError(w http.ResponseWriter, statusCode int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	encoder := json.NewEncoder(w)
	encoder.Encode(APIErrorResponse{
		StatusCode: statusCode,
		Success:    false,
		Message:    message,
	})
}

// GetProductsResponse sends a response for products list with metadata
func GetProductsResponse(w http.ResponseWriter, statusCode int, message string, total int, page int, limit int, result interface{}) {
	totalPages := 1
	if limit > 0 {
		totalPages = (total + limit - 1) / limit
	}
	response := ListResponse{
		StatusCode: statusCode,
		Success:    true,
		Message:    message,
		Data: ListData{
			Meta: Meta{
				Total:      total,
				Page:       page,
				Limit:      limit,
				TotalPages: totalPages,
			},
			Result: result,
		},
	}
	HandleSend(w, statusCode, response)
}

// GetResponse sends a success response for getting a single item
func GetSingleProductResponse(w http.ResponseWriter, statusCode int, message string, data interface{}) {
	response := SingleResponse{
		StatusCode: statusCode,
		Success:    true,
		Message:    message,
		Data:       data,
	}
	HandleSend(w, statusCode, response)
}

// CreateResponse sends a success response for creating an item
func CreateResponse(w http.ResponseWriter, statusCode int, message string, data interface{}) {
	response := SingleResponse{
		StatusCode: statusCode,
		Success:    true,
		Message:    message,
		Data:       data,
	}
	HandleSend(w, statusCode, response)
}

// UpdateResponse sends a success response for updating an item
func UpdateResponse(w http.ResponseWriter, statusCode int, message string, data interface{}) {
	response := SingleResponse{
		StatusCode: statusCode,
		Success:    true,
		Message:    message,
		Data:       data,
	}
	HandleSend(w, statusCode, response)
}

// DeleteResponse sends a success response for deleting an item
func DeleteResponse(w http.ResponseWriter, statusCode int, message string) {
	response := EmptyResponse{
		StatusCode: statusCode,
		Success:    true,
		Message:    message,
	}
	HandleSend(w, statusCode, response)
}
