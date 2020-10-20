package main

import "fmt"

type person struct { //ninja 6.5
	first string
	last  string
	age   int
}

func (p person) speak() {
	fmt.Println("I am", p.first, p.last, "I am", p.age, "years old.")
}

func main() {
	p1 := person{
		first: "James",
		last:  "Bond",
		age:   30,
	}

	p1.speak()
}

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
