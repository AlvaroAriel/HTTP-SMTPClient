package main

import (
	"fmt"
	"net/http"
)

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("POST /send", handleSendEmail())

	fmt.Println("Starting server on port 8080...")
	err := http.ListenAndServe(":8080", mux)

	if err != nil {
		fmt.Println(fmt.Errorf("something went wrong while initializating the serve %w", err))
	}
}

func handleSendEmail() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		//implement
	}

}
