package controllers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	configs "github.com/newlinedeveloper/go-api/Configs"
	models "github.com/newlinedeveloper/go-api/Models"
	responses "github.com/newlinedeveloper/go-api/Responses"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
)

var userCollection *mongo.Collection = configs.GetCollection(configs.DB, "users")
var validate = validator.New()

func CreateMember() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var member models.Member
		defer cancel()

		//validate the request body
		if err := json.NewDecoder(r.Body).Decode(&member); err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			response := responses.MemberResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}}
			json.NewEncoder(rw).Encode(response)
			return
		}

		//use the validator library to validate required fields
		if validationErr := validate.Struct(&member); validationErr != nil {
			rw.WriteHeader(http.StatusBadRequest)
			response := responses.MemberResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}}
			json.NewEncoder(rw).Encode(response)
			return
		}

		newUser := models.Member{
			Id:    primitive.NewObjectID(),
			Name:  member.Name,
			Email: member.Email,
			City:  member.City,
		}
		result, err := userCollection.InsertOne(ctx, newUser)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			response := responses.MemberResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}}
			json.NewEncoder(rw).Encode(response)
			return
		}

		rw.WriteHeader(http.StatusCreated)
		response := responses.MemberResponse{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": result}}
		json.NewEncoder(rw).Encode(response)
	}
}

func GetMember() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		params := mux.Vars(r)
		userId := params["id"]
		var user models.Member

		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(userId)

		err := userCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&user)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			response := responses.MemberResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}}
			json.NewEncoder(rw).Encode(response)
			return
		}

		rw.WriteHeader(http.StatusOK)
		response := responses.MemberResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": user}}
		json.NewEncoder(rw).Encode(response)

	}
}
