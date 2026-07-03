package main

import (
	"fmt"
	"net/http"
)

func handleHello (w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w, "Hello world");
}

func handleGoodbye (w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w, "Goodbye world");
}

func main() {
	mux := http.NewServeMux(); //router create 

	mux.HandleFunc("/hello", handleHello); // route declaration
	mux.HandleFunc("/good-bye", handleGoodbye);

	fmt.Println("Server port running:3004");

	err := http.ListenAndServe(":3004", mux);

	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}