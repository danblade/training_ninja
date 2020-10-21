package main

import "fmt"

func main() { //ninja 6.9.1
	f := foo()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
}

func foo() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}

// func main() { //ninja 6.9
// 	g := func(xi []int) int {
// 		if len(xi) == 0 {
// 			return 0
// 		}
// 		if len(xi) == 1 {
// 			return xi[0]
// 		}
// 		return xi[0] + xi[len(xi)-1]
// 	}
// 	x := foo(g, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
// 	fmt.Println(x)
// }

// func foo(f func(xi []int) int, ii []int) int {
// 	n := f(ii)
// 	n++
// 	return n
// }

// func foo() func() int { //ninja 6.8
// 	return func() int {
// 		return 7
// 	}
// }

// func main() { //ninja 6.8
// 	f := foo()
// 	fmt.Println(f())

// }

// func main() { //ninja 6.7
// 	f := func() {
// 		for i := 0; i < 10; i++ {
// 			fmt.Println(i)
// 		}
// 	}
// 	f()
// }

// func main() { //ninja 6.6
// 	func() {
// 		for i := 0; i < 100; i++ {
// 			fmt.Println(i)
// 		}
// 	}()

// }

// import ( //ninja 6.5
// 	"fmt"
// 	"math"
// )

// type circle struct {
// 	radius float64
// }

// type square struct {
// 	length float64
// }

// func (s square) area() float64 {
// 	return s.length * s.length
// }

// func (c circle) area() float64 {
// 	return math.Pi * (c.radius * c.radius)
// }

// type shape interface {
// 	area() float64
// }

// func info(t string, s shape) {
// 	fmt.Println(t, s.area())
// }

// func main() {
// 	circ := circle{radius: 12.345}
// 	squ := square{length: 15}

// 	info("Circle area:", circ)
// 	info("Square area:", squ)
// }

// type person struct { //ninja 6.4
// 	first string
// 	last  string
// 	age   int
// }

// func (p person) speak() {
// 	fmt.Println("I am", p.first, p.last, "I am", p.age, "years old.")
// }

// func main() {
// 	p1 := person{
// 		first: "James",
// 		last:  "Bond",
// 		age:   30,
// 	}

// 	p1.speak()
// }

// func foo(n ...int) int { //ninja 6.3
// 	total := 0
// 	for _, v := range n {
// 		total += v
// 	}
// 	return total
// }

// func bar(a []int) int {
// 	total := 0
// 	for _, v := range a {
// 		total += v
// 	}
// 	return total
// }

// func main() {

// 	defer fmt.Println(
// 		"Calling and running foo() and defering",
// 		foo([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}...),
// 		"\n",
// 	)

// 	fmt.Println(
// 		"Calling and running bar()",
// 		bar([]int{1, 6, 7, 8, 9, 44, 12}),
// 	)
// }

// func foo(n ...int) int { //ninja 6.2
// 	total := 0
// 	for _, v := range n {
// 		total += v
// 	}
// 	return total
// }

// func bar(a []int) int {
// 	total := 0
// 	for _, v := range a {
// 		total += v
// 	}
// 	return total
// }

// func main() {

// 	fmt.Println(foo([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}...),
// 		bar([]int{1, 6, 7, 8, 9, 44, 12}))

// }

// func foo() int { //ninja 6.1
// 	return 7
// }

// func bar() (int, string) {
// 	return 11, "How bout you?"
// }

// func main() {

// 	x, y := bar()

// 	fmt.Println(foo(), x, y)

// }
