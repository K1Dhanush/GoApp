package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	A()
}
func A() {
	defer func() {
		if x := recover(); x != nil { //assign and condition in one line
			fmt.Printf("Panic error is : %v", x)
		}
	}()
	defer fmt.Println("Hello")
	defer fmt.Println("Function A")
	B()
}

func B() {
	fmt.Println("Function B")
	_, err := os.Open("C:\\Users\\OGS\\Documents\\toBeDone.txt")
	var pathError *os.PathError
	if errors.As(err, &pathError) { //error -- should be nil by Default.
		fmt.Println("Failed to open the File with given FileName : " + pathError.Path)
	}
	//panic(errors.New("Ho no!"))
}
