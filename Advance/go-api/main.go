package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func getMessage(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"data": "Golang project setup test"})

}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", getMessage).Methods("GET")
	fmt.Print("Server is running on port 8000 !!!")
	log.Fatal(http.ListenAndServe(":8000", router))

}
