package routes

import (
	"github.com/gorilla/mux"
	controllers "github.com/newlinedeveloper/go-api/Controllers"
)

func MemberRoutes(router *mux.Router) {

	router.HandleFunc("/member", controllers.CreateMember()).Methods("POST")
	router.HandleFunc("/member/{id}", controllers.GetMember()).Methods("GET")
	router.HandleFunc("/members", controllers.GetAllMembers()).Methods("GET")
	router.HandleFunc("/member/{id}", controllers.UpdateMember()).Methods("PUT")
	router.HandleFunc("/member/{id}", controllers.DeleteMember()).Methods("DELETE")

}
