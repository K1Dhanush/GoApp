package main

import "fmt"

type Person struct {
	name string
	age  int
}

func main() {
	//pointers declaration
	// var pers1 Person
	// var p *Person = &pers1 //address of the variable n
	// fmt.Println(&p)
	// changeVal(p)
	// fmt.Println(*p)

	// //Person struct type
	// pers1 := Person{"Dhaneshwar", 78}

	// //Here it is pointing to the struct
	// pts := &pers1

	// fmt.Println(pts.age)
	// fmt.Println(&pers1)
	// fmt.Println(&pers1.age)

	//Double Pointers
	v := 45
	var p *int = &v
	var p1 **int = &p
	fmt.Println(p, p1)
}

// func changeVal(v *Person) {
// 	*v = Person{"Dhaneshwar", 45} //assigning the value
// }
