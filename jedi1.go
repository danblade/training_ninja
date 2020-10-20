package main

//ninga level 1
import (
	"fmt"
)

//ninga level 1

//lvl 1.5
type theNumber int

var x theNumber
var y = "James Bond"
var z = true
var xy int

func main() {
	x = 42
	fmt.Printf("The secret code was %v (which was of type, %[1]T), %v finally figured out it was, %v. \n", x, y, z)
	xy = int(x)
	fmt.Println(xy)
	fmt.Println(x)
}

// //lvl 1.4
// type theNumber int

// var x theNumber
// var y = "James Bond"
// var z = true

// func main() {
// 	x = 42
// 	fmt.Printf("The secret code was %v (which was of type, %[1]T), %v finally figured out it was, %v. \n", x, y, z)
// }

// //lvl 1.3
// var x = 42
// var y = "James Bond"
// var z = true

// func main() {
// 	fmt.Printf("The secret code was %v, %v finally figured out it was, %v. \n", x, y, z)
// }

// //lvl 1.2
// var x int
// var y string
// var z bool

// func main() {
// 	fmt.Printf("The secret code was %v, %v finally figured out it was, %v. \n", x, y, z)
// }

// //ninja lvl 1.1
// func main() {
// 	x := 42
// 	y := "James Bond"
// 	z := true
// 	fmt.Printf("The secret code was %v, %v finally figured out it was, %v. \n", x, y, z)
// }
