package main

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/gorilla/mux"
)

// handling url's
const urlRaw = "http://localhost:8080/hello"

/*using Gorilla/Mux*/
func main() {

	//creating a new router
	r := mux.NewRouter()

	r.HandleFunc("/hello", helloHandler).Methods("GET")
	r.HandleFunc("/hello", addData).Methods("POST")
	r.HandleFunc("/hello", updateData).Methods("PUT")
	r.HandleFunc("/hello", deleteData).Methods("DELETE")
	r.HandleFunc("/hi", hello)

	//starting the server
	http.ListenAndServe(":8080", r)

	//for url handling -- parsing
	result, _ := url.Parse(urlRaw)
	urlHandling(result)

}
func helloHandler(w http.ResponseWriter, r *http.Request) {
	// vars := mux.Vars(r) //to get data from the segments , passing http-request as a parameter
	// title := vars["title"]
	fmt.Fprintf(w, "Welcome to my website!")
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "HomePage")
}

func addData(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Data is added Successfully")
}

func updateData(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Data is updated Successfully")
}

func deleteData(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Data is deleted Successfully")
}

// for Handling the -- url's
func urlHandling(result *url.URL) {
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.RawQuery) //parameter in url
}
