package main

import (
	"fmt"
	"log"
	"net/http"

	configs "github.com/newlinedeveloper/go-api/Configs"

	routes "github.com/newlinedeveloper/go-api/Routes"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	// MongoDB Connection
	configs.ConnectDB()

	routes.MemberRoutes(router)

	fmt.Print("Server is running on port 8000 !!!!")
	log.Fatal(http.ListenAndServe(":8000", router))
}
