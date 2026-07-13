package main

import (
	"encoding/json"
	"net/http"
)

func handleSend(w http.ResponseWriter, data interface{}) {
	encoder := json.NewEncoder(w)
	encoder.Encode(data)
}
