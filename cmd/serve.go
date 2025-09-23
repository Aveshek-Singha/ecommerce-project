package cmd

import (
	"ecommerice-project/middleware"
	"ecommerice-project/util"
	"fmt"
	"net/http"
)

func Serve() {
	manager := middleware.NewManager()

	manager.Use(middleware.Logger, middleware.Hudai)

	mux := http.NewServeMux()

	initRoutes(mux, manager)

	globalRouter := util.GlobalRouter(mux)

	fmt.Println("Server running on :3000")

	err := http.ListenAndServe(":3000", globalRouter) // "Failed to start the server"
	if err != nil {
		fmt.Println("Error starting the server", err)
	}

}
