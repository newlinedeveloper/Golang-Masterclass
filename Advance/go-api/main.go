package main

import (
	"fmt"
	"log"
	"net/http"

	configs "github.com/newlinedeveloper/go-api/Configs"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	// MongoDB Connection
	configs.ConnectDB()

	fmt.Print("Server is running on port 8000 !!!!")
	log.Fatal(http.ListenAndServe(":8000", router))
}
