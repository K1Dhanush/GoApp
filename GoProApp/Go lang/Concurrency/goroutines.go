// package main

// import (
// 	"fmt"
// 	"sync"
// 	"time"
// )

// // Gobal Variable
// var (
// 	counter int
// )

// func increment(id int) {

// 	fmt.Printf("Goroutine %d: Started\n", id)
// 	time.Sleep(3 * time.Second)
// 	counter++
// 	fmt.Printf("Goroutine %d: Incremented the counter to %d\n", id, counter)
// }

// func main() {
// 	var wg sync.WaitGroup

// 	for i := 1; i <= 5; i++ {
// 		wg.Add(1)
// 		go func(id int) {
// 			defer wg.Done()
// 			increment(id)
// 		}(i)
// 	}

// 	wg.Wait()
// 	fmt.Println("Main: All Goroutines Completed")
// 	fmt.Println("Final Counter Value:", counter)
// }

package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup
var mut sync.Mutex

func main() {
	wg.Add(2)
	go paytabs("Hello")

	//to sleep to execute the another GoRoutine
	time.Sleep(3 * time.Second)

	go paytabs("World")

	wg.Wait()
	fmt.Println("Two GoRoutines are executed.")
}

func paytabs(str string) {
	defer wg.Done()
	mut.Lock()
	fmt.Println("Enter your Name: ")
	mut.Unlock()
}
