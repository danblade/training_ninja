package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type person struct { //ninja 7.2
	first string
	last  string
}

func changeMe(p *person) {

	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Change the first name to: ")
	text, _ := reader.ReadString('\n')
	text = strings.TrimSuffix(text, "\n")
	p.first = text
}

func main() {
	p1 := person{
		first: "James",
		last:  "Bond",
	}

	fmt.Println(p1)
	changeMe(&p1)
	fmt.Println(p1)
}

// func main() { //ninja 7.1
// 	v := 1340
// 	fmt.Println("The variable value is,", v, "the address of the variable is", &v)
// }
