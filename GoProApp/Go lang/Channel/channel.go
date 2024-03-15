package main

import (
	"fmt"
)

// func main() {
// 	var wg sync.WaitGroup
// 	//var myCh chan int

// 	myCh := make(chan int)

// 	wg.Add(2)
// 	//reciever
// 	go func() {
// 		defer wg.Done()
// 		ch := <-myCh
// 		fmt.Println(ch)
// 	}()

// 	//sender
// 	go func() {
// 		defer wg.Done()
// 		myCh <- 4
// 	}()

// 	//go channel()
// 	//fmt.Printf("%v", myCh)
// 	wg.Wait()
// 	fmt.Printf("%T", myCh)
// }

// func channel() {
// 	defer wg.Done()
// 	myCh <- 78
// 	fmt.Printf("%v", myCh)
// }

//select statement

func main() {
	//multiple channels
	ch := make(chan string)
	ch1 := make(chan string)

	//go routines
	go func() {
		//time.Sleep(2 * time.Second)
		ch <- "Hi"
	}()

	go func() {
		ch1 <- "Hello"
	}()

	//handling multiple channels using -- select
	select {
	case msg := <-ch:
		fmt.Print(msg)
	case msg := <-ch1:
		fmt.Print(msg)
	}
}
