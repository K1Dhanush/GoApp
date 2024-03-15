package pratice

import (
	"fmt"
)

type Person struct {
	name   string
	age    int
	height float32
}

func main2() {
	// myslice := []string{}
	// fmt.Printf("length of myslice is :%v\n", len(myslice))
	// fmt.Printf("length of myslice is :%v\n", cap(myslice))
	// fmt.Printf("length of myslice is :%v\n", myslice)

	// my_slice := make([]int, 5, 9)

	// arr := [5]int{1, 5, 7, 8, 8}
	// my_slice := arr[2:4]
	// fmt.Printf("length of myslice is :%v\n", len(my_slice))
	// fmt.Printf("capacity of myslice is :%v\n", cap(my_slice))
	// fmt.Printf("myslice is :%v\n", my_slice)
	// fmt.Println("Grow Slice is ", my_slice[:cap(my_slice)]) //to get the elements after the length index

	// my_slice = append(my_slice, 4, 7, 8)
	// fmt.Println("Grow Slice is ", my_slice[:cap(my_slice)])
	// fmt.Printf("length of myslice is :%v\n", len(my_slice))
	// fmt.Printf("capacity of myslice is :%v\n", cap(my_slice))
	// return "Slice_file"

	var pers1 Person
	//var pers2 Person
	myFunction(pers1)
}

func myFunction(pers Person) {
	pers.age = 21
	pers.height = 5.11
	pers.name = "Dhanesh"
	fmt.Printf("%v\t%v\t%v", pers.name, pers.age, pers.height)
}

func MySlice() {
	my_slice := []int{1, 2, 3, 4, 5}
	my_slice = append(my_slice[:2], my_slice[3:]...)
	fmt.Println(my_slice)

}
