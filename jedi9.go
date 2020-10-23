package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() { //ninja 9.3 (start) need to build with race condition.

	var wg sync.WaitGroup
	incrementer := 0
	wg.Add(100)
	for i := 0; i < 100; i++ {

		go func() {

			v := incrementer
			runtime.Gosched()
			v++
			incrementer = v
			fmt.Println(incrementer)
			wg.Done()
		}()

	}
	wg.Wait()
	fmt.Println("End:", incrementer)
} //ninja 9.3 (end)

// type person struct { //ninja 9.2 (start)
// 	firstName string
// 	lastName  string
// 	age       int
// }
// type human interface {
// 	speak()
// }

// func (p *person) speak() {
// 	fmt.Println("Hello! \nI am:", p.firstName, p.lastName)
// }

// func saySomething(h human) {
// 	h.speak()
// }

// func main() {
// 	p1 := person{

// 		firstName: "James",
// 		lastName:  "Bond",
// 		age:       32,
// 	}

// 	saySomething(&p1)

// } //ninja 9.2 (end)

// func threadTwo() { //ninga 9.1 (start)
// 	fmt.Println("This is the second thread.")
// }

// func threadThree() {
// 	fmt.Println("This is the third thread.")
// }

// func threadFour() {
// 	fmt.Println("This is the fourth thread.")
// }

// func main() {

// 	var wg sync.WaitGroup

// 	fmt.Println("CPUs:", runtime.NumCPU())
// 	fmt.Println("Goroutines:", runtime.NumGoroutine())

// 	fmt.Println("This is the main thread.")

// 	wg.Add(1)
// 	go func() {
// 		threadTwo()
// 		fmt.Println("Goroutines:", runtime.NumGoroutine())
// 		wg.Done()
// 	}()

// 	wg.Add(1)
// 	go func() {
// 		threadThree()
// 		fmt.Println("Goroutines:", runtime.NumGoroutine())
// 		wg.Done()
// 	}()

// 	wg.Add(1)
// 	go func() {
// 		threadFour()
// 		fmt.Println("Goroutines:", runtime.NumGoroutine())
// 		wg.Done()
// 	}()

// 	fmt.Println("Waiting for the other threads to finish")
// 	wg.Wait()
// } //ninja 9.1 (end)
