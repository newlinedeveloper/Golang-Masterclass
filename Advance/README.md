# Golang-Masterclass
Golang Masterclass - Advanced Concepts

### prerequisites


```
Please check this link https://go.dev/doc/install

go version


# MongoDB installation

Please check this link https://www.mongodb.com/docs/manual/installation/


```

##### Create go-api project 

```
mkdir go-api

cd go-api

go mod init github.com/newlinedeveloper/go-api


```

###### Install Required Packages

```
go get -u github.com/gorilla/mux go.mongodb.org/mongo-driver/mongo github.com/joho/godotenv github.com/go-playground/validator/v10

```

```
github.com/gorilla/mux

go.mongodb.org/mongo-driver/mongo

github.com/joho/godotenv

github.com/go-playground/validator/v10

```


### Go simple Web server



```

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

```

To run the application

```
go run main.go

Server is running on port 8000 !!!

```

#### Project structure



```
.
└── go-api/
    ├── Routes/
    │   └── member_routes.go
    ├── Controllers/
    │   └── member_controllers.go
    ├── Models/
    │   └── member_models.go
    ├── Configs/
    │   ├── env.go
    │   └── connection.go
    ├── Responses/
    │   └── response.go
    ├── main.go
    ├── go.mod
    ├── go.sum
    ├── .env
    └── .env.example


```


#### MongoDB Connection setup

Create .env file and Add Mongo DB connection uri

```
MONGOURI=mongodb://localhost:27017
```

Create `env.go` file in `Configs` folder


```
package configs

import (
    "log"
    "os"
    "github.com/joho/godotenv"
)

func EnvMongoURI() string {
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    return os.Getenv("MONGOURI")
}

```


Create `connection.go` file in `Configs` folder


```
package configs

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoDB connection function
func ConnectDB() *mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI(EnvMongoURI()))
	if err != nil {
		log.Fatal(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	//ping the database
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("MongoDB Connected successfully !!!! ")
	return client
}

// // MongoDB Client instance
// var DB *mongo.Client = ConnectDB()

//Getting database collection
func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	collection := client.Database("golang-masterclass").Collection(collectionName)
	return collection
}


```

in `main.go` file


```

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



```

#### Create Member Model

Create `member_models.go` file in `Models` folder



```

package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Member struct {
	Id    primitive.ObjectID `json:"id,omitempty"`
	Name  string             `json:"name,omitempty" validate:"required"`
	Email string             `json:"email,omitempty" validate:"required"`
	City  string             `json:"city,omitempty" validate:"required"`
}


```


#### Create Member Response struct

Create `member_responses.go` file in `Responses` folder



```

package responses

type MemberResponse struct {
	Status  int                    `json:"status"`
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data"`
}


```


#### Create Members Api Routes

Create `member_routes.go` file in `Routes` folder


```

package routes

import "github.com/gorilla/mux"

func MemberRoutes(router *mux.Router) {

}



```


import routes to `main.go` file



```
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

    // Imported Members routes
	routes.MemberRoutes(router)

	fmt.Print("Server is running on port 8000 !!!!")
	log.Fatal(http.ListenAndServe(":8000", router))
}


```



#### Create Members Controller functions

Create `member_controllers.go` file in `Controllers` folder



Create `CreateMember` function


```

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



```

update routes file


```
package routes

import (
	"github.com/gorilla/mux"
	controllers "github.com/newlinedeveloper/go-api/Controllers"
)

func MemberRoutes(router *mux.Router) {

	router.HandleFunc("/member", controllers.CreateMember()).Methods("POST")

}

```









