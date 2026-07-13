package main

import "net/http"

func HandleGlobarlRouter(mux *http.ServeMux) http.Handler {
	handleAllRequests := func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "OPTIONS" {
			w.Header().Set("Content-Type", "application/json")
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Rakib")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
			w.WriteHeader(http.StatusOK)
			return
		}

		mux.ServeHTTP(w, r)
	}

	return http.HandlerFunc(handleAllRequests)
}
