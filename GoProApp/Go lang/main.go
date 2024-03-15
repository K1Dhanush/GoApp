/* TASK -- 1
package main //pakage name should always be "main"

import "fmt"   //fmt -- format

func main() { //this is the main function to the run the program.
	fmt.Println("HelloWorld")
}
*/

/*
package main

import (

	"fmt"

)

	func main() {
		var num uint = 4 //variable is a temporary storage -- I am assigning integer so,that it is "int" datatype
		name := "Dhanesh"
		boolean := true
		var num1 float32
		fmt.Println("Enter the value :")
		fmt.Scan(&num1)
		fmt.Println(name)
		fmt.Printf("%.2f\n", num1)
		// fmt.Println("datatype of num is:", reflect.TypeOf(num))
		// fmt.Println("datatype of num1 is: ", reflect.TypeOf(num1))
		// fmt.Println("datatype of boolean is: ", reflect.TypeOf(boolean))
		fmt.Printf("Type : %T , Value = %v\n", num, num)
		fmt.Printf("Type : %T , Value = %v\n", name, name)
		fmt.Printf("Type : %T , Value = %v\n", num1, num1)
		fmt.Printf("Type : %T , Value = %v\n", boolean, boolean)
	}
*/

// Array
package main

import (
	pratice "Project/Slice"
	"fmt"
) //importing packages
var a int = 10

func main() {
	//var arr = [4]int{1, 2, 3}
	//arr := [4]int{4, 7, 8, 9}
	var b = 10

	// arr := [4]int{}
	// arr[1] = 7

	//arr := [4]int{1: 10, 2: 78}

	arr := [...]int{1, 2}
	fmt.Println(len(arr))
	//fmt.Print("Slice :\n")
	//main2()
	//condition()
	a = 1
	help(b)
	// var x, y int
	// fmt.Scan(&x, &y)
	// fmt.Println(Calculator(x, y))

	//Maps
	// cars := make(map[int]string)
	// cars[1] = "BMW"
	// cars[2] = "Frod"

	// vehicles := cars
	// fmt.Println(cars)
	// fmt.Println(vehicles)
	// vehicles[2] = "Maruthi"

	//fmt.Println(cars, vehicles)
	// delete(cars, 2)
	// fmt.Println(cars)
	// for k, v := range vehicles {
	// 	fmt.Printf("%v\t%v\n", k, v)
	// }

	//main2()
	//mySlice()

	//for reading the File
	// a, e := os.ReadFile("C:\\Users\\OGS\\Documents\\toBeDone.txt")
	// if e != nil {
	// 	e = errors.New("error in file opening")
	// }
	// fmt.Println(string(a), e)
	fmt.Println(b)
	pratice.MySlice()
}
func help(b int) {
	b = 1

}

// func Calculator(x int, y int) (add int, sub int, mul int) {
// 	add = x + y
// 	sub = x - y
// 	mul = x * y
// 	return
// }
