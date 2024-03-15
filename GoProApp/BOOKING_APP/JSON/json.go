package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Book struct {
	Name     string `json:"coursename"`
	Price    int
	Platfrom string
	Password string `json:"-"`
}

func main() {
	fmt.Println("Welcome to golang to -- JSON")
	//encodeJson()
	//decodeJson()
	decodeJsonFile()
}

// func encodeJson() {
// 	book := []Book{{"Ramayanam", 452, "Valli", "abc@123"}, {"Bharatha", 789, "Vasu", "def@456"}}

// 	//parsing
// 	jsonData, _ := json.MarshalIndent(book, "", "\t") //json is in byte from

// 	//storing data int json.file
// 	//create the json file
// 	file, err := os.Create("json.file")
// 	if err != nil {
// 		fmt.Println("File is not Created")
// 	}
// 	_, _ = file.Write(jsonData) //int,err and param is in []byte from

// 	//fmt.Println(string(jsonData))
// }

// func decodeJson() {
// 	jsonData := []byte(`
// 	{
// 		"coursename": "Ramayanam",
// 		"Price": 452,
// 		"Platfrom": "Valli"
//     }`)

// 	checkJson := json.Valid(jsonData)

// 	//Decoding json data in STRUCT
// 	var course Book
// 	if checkJson {
// 		fmt.Println("Valid JSON File")
// 		json.Unmarshal(jsonData, &course) //storing json in this -- variable "struct"
// 		fmt.Printf("%#v\n", course)
// 	} else {
// 		fmt.Println("Invalid JSON")
// 	}

// 	//JSON as key value -- using map datastructure
// 	//map
// 	var myMap map[string]interface{}
// 	json.Unmarshal(jsonData, &myMap)
// 	fmt.Printf("%#v\n", myMap)
// 	// _, ok := myMap["Price"]
// 	// fmt.Println(ok)
// 	for k, v := range myMap {
// 		fmt.Printf("Key is %v and value is %v and type of :%T\n", k, v, v)
// 	}
// }

func decodeJsonFile() {
	file, err := os.Open("file.json")
	if err != nil {
		fmt.Print("File is not Found")
	} else {
		fileRead, _ := io.ReadAll(file)

		var books []Book                                         //creating Slice to read multiple structs of type -- book
		if err := json.Unmarshal(fileRead, &books); err != nil { //reading the variable and condition in oneline
			fmt.Println("Error unmarshalling JSON:", err)
			return
		}

		fmt.Println("Books:")
		for _, book := range books { // Print the data
			fmt.Printf("Name: %s, Price: %d, Platform: %s\n", book.Name, book.Price, book.Platfrom)
		}
	}

}
