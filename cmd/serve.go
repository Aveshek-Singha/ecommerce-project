package cmd

import (
	handler "ecommerice-project/handlers"
	"ecommerice-project/util"
	"fmt"
	"net/http"
)

func Serve() {
	mux := http.NewServeMux() //router

	mux.Handle("GET /products", http.HandlerFunc(handler.GetProducts))

	mux.Handle("POST /products", http.HandlerFunc(handler.CreateProduct))

	mux.Handle("GET /products/{productId}", http.HandlerFunc(handler.GetProductById))

	fmt.Println("Server running on :3000")

	globalRouter := util.GlobalRouter(mux)

	err := http.ListenAndServe(":3000", globalRouter) // "Failed to start the server"
	if err != nil {
		fmt.Println("Error starting the server", err)
	}

}
