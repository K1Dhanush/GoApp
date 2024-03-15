package main

import (
	"fmt"
	"sync"
)

var mut sync.Mutex //for locking
var count int

func main() {
	var wg sync.WaitGroup

	wg.Add(2)
	go func() {
		defer wg.Done()
		mut.Lock()
		for i := 0; i < 10; i++ {
			count++
			fmt.Println(count)
		}
		mut.Unlock()
		fmt.Println(count)
	}()

	go func() {
		defer wg.Done()
		//mut.Lock()
		//array = append(array, 2)
		for i := 0; i < 5; i++ {
			count--
			fmt.Println(count)
		}
		// mut.Unlock()
		fmt.Println(count)
	}()

	// fmt.Println(count)
	//go hello(&wg, count)
	wg.Wait()
	fmt.Println(count)
}

// func hello(wg *sync.WaitGroup, count int) {
// 	defer wg.Done()
// 	mut.Lock()
// 	count = count + 2
// 	// array = append(array, 3)
// 	mut.Unlock()
// 	//fmt.Println(count)
// }
