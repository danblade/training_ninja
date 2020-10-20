package main

import "fmt"

//ninja lvl 2.4
const (
	num  = 2020 + iota
	numA = 2020 + iota
	numB = 2020 + iota
	numC = 2020 + iota
)

func main() {

	fmt.Println(num)
	fmt.Println(numA)
	fmt.Println(numB)
	fmt.Println(numC)
}

// //ninja lvl 2.4
// func main() {
// 	num := `The number I like is
// 	733 (decimal's),
// 	"1335" octal,
// 	'2dd' hex, and
// 	"1011011101" binary.`
// 	fmt.Println(num)
// }

// //ninja lvl 2.3
// func main() {
// 	num := 733
// 	fmt.Printf("%5d (decimal), %[1]o octal, %[1]x hex, and %[1]b binary.\n", num)
// 	num = num << 1
// 	fmt.Printf("%d (decimal), %[1]o octal, %[1]x hex, and %[1]b binary.\n", num)
// }

//output: The number I like is 733 (decimal), 1335 octal, 2dd hex, and 1011011101 binary.

// //ninja lvl 2.2
// func main() {
// 	g := (733 == 733)
// 	h := (733 <= 733)
// 	i := (733 >= 733)
// 	j := (733 != 733)
// 	k := (733 < 777)
// 	l := (733 > 700)
// 	fmt.Printf("%v,\n%v,\n%v,\n%v,\n%v,\n%v\n", g, h, i, j, k, l)
// }

// //ninja lvl 2.1
// func main() {
// 	num := 733
// 	fmt.Printf("The number I like is %d (decimal), %[1]o octal, %[1]x hex, and %[1]b binary.\n", num)
// }

//output: The number I like is 733 (decimal), 1335 octal, 2dd hex, and 1011011101 binary.

// //Find architecture
// func main() {
// 	fmt.Println(runtime.GOOS)   //dawin
// 	fmt.Println(runtime.GOARCH) //amd64
// }
