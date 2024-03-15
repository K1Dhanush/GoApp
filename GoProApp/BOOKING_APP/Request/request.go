package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// struct
type course struct {
	Name     string //`json:"coursename"`
	Price    int
	Platform string //`json:"platform"`
}

// gobal Variable map
var myCourse = make(map[string]course)

// using http package
func main() {

	// Add values to the map
	myCourse["math"] = course{Name: "Mathematics", Price: 100, Platform: "Online"}
	myCourse["physics"] = course{Name: "Physics", Price: 150, Platform: "Offline"}
	myCourse["history"] = course{Name: "History", Price: 75, Platform: "Online"}

	r := mux.NewRouter()
	r.HandleFunc("/home", addCourse).Methods("POST")
	r.HandleFunc("/home", getAll).Methods("GET")
	r.HandleFunc("/home/{Name}", get).Methods("GET")
	r.HandleFunc("/home/{Name}/update", updateCourse).Methods("PUT")
	r.HandleFunc("/home/{Name}/delete", deleteCourse).Methods("DELETE")

	//to start server
	http.ListenAndServe(":8081", r)
}
func addCourse(w http.ResponseWriter, r *http.Request) {
	var book course //struct variable
	w.Header().Add("Content-Type", "application/json")

	decode := json.NewDecoder(r.Body)
	err := decode.Decode(&book)

	//add into the map
	myCourse[book.Name] = book

	if err != nil {
		fmt.Println("Unknown Error")
		return
	}
	// json.NewEncoder(w).Encode(myCourse)  //to get all the Courses
	json.NewEncoder(w).Encode(book)
}

func getAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json") //to get in JSON format
	encoder := json.NewEncoder(w)
	_ = encoder.Encode(myCourse)
}

func get(w http.ResponseWriter, r *http.Request) {

	//to get the values from the user
	vars := mux.Vars(r)
	coursename := vars["Name"]
	get, ok := myCourse[coursename]
	if !ok {
		fmt.Fprintf(w, "course is not present")
	} else {
		json.NewEncoder(w).Encode(get)
	}
}

func updateCourse(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	coursename := vars["Name"]

	_, ok := myCourse[coursename]
	if !ok {
		fmt.Fprintf(w, "User is not present")
	} else {
		//creating a -- struct variable
		var update course
		_ = json.NewDecoder(r.Body).Decode(&update)
		myCourse[update.Name] = update

		//for printing the
		json.NewEncoder(w).Encode(update)
	}
}

func deleteCourse(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	coursename := vars["Name"] //getting the variable
	_, ok := myCourse[coursename]
	if !ok {
		fmt.Fprintf(w, "User is not present")
	} else {
		delete(myCourse, coursename)
		fmt.Fprintf(w, "Deleted Successfully")
	}
}
