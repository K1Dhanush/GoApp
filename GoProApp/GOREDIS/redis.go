package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/gorilla/mux"
)

// Struct
type User struct {
	Id        int    `json:"id"`
	FirstName string `json:"firstname"`
	Age       int    `json:"age"`
}

// for variable
type JSONResponse struct {
	Users []User `json:"users"`
}

var ctx = context.Background()

func main() {
	//Creating the client to interact with the redis server.
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", //"redis:6379",
		Password: "",
		DB:       0, //default database
	})

	_, err := client.Ping(ctx).Result()
	if err != nil {
		panic(err)
	}

	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		ExplicitApi(w, r, client)
	})

	http.ListenAndServe(":8085", r)

}
func ExplicitApi(w http.ResponseWriter, r *http.Request, client *redis.Client) {
	responseApi, errr := http.Get("https://dummyjson.com/users")
	if errr != nil {
		fmt.Fprint(w, errr)
		return
	}

	var response JSONResponse
	//json to struct
	err := json.NewDecoder(responseApi.Body).Decode(&response)
	if err != nil {
		http.Error(w, "Data is not Json- Format", http.StatusInternalServerError)
		return
	}
	w.Header().Add("Content-Type", "application/json")

	//for stroing the data into the redis -- struct to json
	jsonData, err1 := json.Marshal(response)
	if err1 != nil {
		http.Error(w, "Data is not Json- Format", http.StatusInternalServerError)
		return
	}
	key := "DummyJson"
	err2 := client.Set(ctx, key, jsonData, 1*time.Minute)
	if err2 != nil {
		fmt.Fprint(w, err2)
		return
	}
}
