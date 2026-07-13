package utils

import (
	"encoding/json"
	"net/http"
)

func HandleSend(w http.ResponseWriter, data interface{}) {
	encoder := json.NewEncoder(w)
	encoder.Encode(data)
}
