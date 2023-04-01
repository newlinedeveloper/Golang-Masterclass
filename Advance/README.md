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
func ConnectDB() *mongo.Client  {
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
    fmt.Println("Connected to MongoDB")
    return client
}

// MongoDB Client instance
var DB *mongo.Client = ConnectDB()

//Getting database collections
func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
    collection := client.Database("golang-masterclass").Collection(collectionName)
    return collection
}

```





