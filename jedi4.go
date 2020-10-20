package main

import "fmt"

func main() { //ninga 4.10
	m := map[string][]string{
		`bond_james`:      []string{`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss`: []string{`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:           []string{`Being evil`, `Ice cream`, `Sunsets`},
	}
	for k, v := range m {
		fmt.Println("The record for", k)
		for i, v2 := range v {
			fmt.Println("\t", i, v2)
		}
	}
	fmt.Printf("\nNow to add to the map.\n\n")

	m[`Observation`] = []string{`Mr. Bond`, `is a`, `womanizer`}
	for k, v := range m {
		fmt.Println("The record for ", k)
		for i, v2 := range v {
			fmt.Println("\t", i, v2)
		}
	}
	fmt.Printf("\nNow delete a record from the map.\n\n")

	delete(m, `no_dr`)
	for k, v := range m {
		fmt.Println("The record for ", k)
		for i, v2 := range v {
			fmt.Println("\t", i, v2)
		}
	}
}

// func main() { //ninga 4.9
// 	m := map[string][]string{
// 		`bond_james`:      []string{`Shaken, not stirred`, `Martinis`, `Women`},
// 		`moneypenny_miss`: []string{`James Bond`, `Literature`, `Computer Science`},
// 		`no_dr`:           []string{`Being evil`, `Ice cream`, `Sunsets`},
// 	}
// 	for k, v := range m {
// 		fmt.Println("The record for", k)
// 		for i, v2 := range v {
// 			fmt.Println("\t", i, v2)
// 		}
// 	}
// 	fmt.Printf("\nNow to add to the map.\n\n")

// 	m[`Observation`] = []string{`Mr. Bond`, `is`, `a womanizer`}
// 	for k, v := range m {
// 		fmt.Println("The record for ", k)
// 		for i, v2 := range v {
// 			fmt.Println("\t", i, v2)
// 		}
// 	}
// }

// 	func main() { //ninga 4.8
// 	m := map[string][]string{
// 		`bond_james`:      []string{`Shaken, not stirred`, `Martinis`, `Women`},
// 		`moneypenny_miss`: []string{`James Bond`, `Literature`, `Computer Science`},
// 		`no_dr`:           []string{`Being evil`, `Ice cream`, `Sunsets`},
// 	}
// 	for k, v := range m {
// 		fmt.Println("The record for ", k)
// 		for i, v2 := range v {
// 			fmt.Println("\t", i, v2)
// 		}
// 	}
// }

// func main() { //ninga 4.7
// 	xs1 := []string{"James", "Bond", "Shaken, not stirred"}
// 	xs2 := []string{"Miss", "Moneypenny", "Helloooooo, James!"}
// 	fmt.Println(xs1)
// 	fmt.Println(xs2)
// 	xxs := [][]string{xs1, xs2}
// 	fmt.Println(xxs)

// 	for i, xs := range xxs {
// 		fmt.Println("Record: ", i)
// 		for j, v := range xs {
// 			fmt.Printf("\tPosition: %v\tvalue: %v,\n", j, v)
// 		}
// 	}
// }

// func main() { //ninja 4.6
// 	x := make([]string, 50, 500)
// 	fmt.Println(len(x))
// 	fmt.Println(cap(x))

// 	for i := 0; i < 50; i++ {
// 		x[i] = fmt.Sprintf("Position %d", i)
// 	}
// 	for i := 0; i < len(x); i++ {
// 		fmt.Println(i, x[i])
// 	}
// 	fmt.Println(x)
// 	fmt.Println(len(x))
// 	fmt.Println(cap(x))

// 	y := []string{` Alabama`, ` Alaska`, ` Arizona`, ` Arkansas`, ` California`, ` Colorado`, ` Connecticut`, ` Delaware`, ` Florida`, ` Georgia`, ` Hawaii`, ` Idaho`, ` Illinois`, ` Indiana`, ` Iowa`, ` Kansas`, ` Kentucky`, ` Louisiana`, ` Maine`, ` Maryland`, ` Massachusetts`, ` Michigan`, ` Minnesota`, ` Mississippi`, ` Missouri`, ` Montana`, ` Nebraska`, ` Nevada`, ` New Hampshire`, ` New Jersey`, ` New Mexico`, ` New York`, ` North Carolina`, ` North Dakota`, ` Ohio`, ` Oklahoma`, ` Oregon`, ` Pennsylvania`, ` Rhode Island`, ` South Carolina`, ` South Dakota`, ` Tennessee`, ` Texas`, ` Utah`, ` Vermont`, ` Virginia`, ` Washington`, ` West Virginia`, ` Wisconsin`, ` Wyoming`}

// 	for i := 0; i < 50; i++ {
// 		x[i] = fmt.Sprintf(y[i])
// 	}

// 	fmt.Println(x)
// 	fmt.Println(len(x))
// 	fmt.Println(cap(x))

// 	for i := 0; i < len(x); i++ {
// 		fmt.Println(i, x[i])
// 	}

// 	x = append(x, y...)

// 	fmt.Println(x)
// 	fmt.Println(len(x))
// 	fmt.Println(cap(x))

// 	for i := 0; i < len(x); i++ {
// 		fmt.Println(i, x[i])
// 	}
// }

// func main() { //ninja 4.5
// 	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
// 	fmt.Println(x)
// 	y = append(x[:3]..., x[6:]...)
// 	fmt.Println(x)
// }

// func main() { //ninja 4.4
// 	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
// 	fmt.Println(x)
// 	x = append(x, 52)
// 	fmt.Println(x)
// 	x = append(x, 53, 54, 55)
// 	fmt.Println(x)
// 	y := []int{56, 57, 58, 59, 60}
// 	x = append(x, y...)
// 	fmt.Println(x)
// }

// func main() { //ninja 4.3
// 	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
// 	fmt.Println(x[:5])
// 	fmt.Println(x[5:])
// 	fmt.Println(x[2:7])
// 	fmt.Println(x[1:6])
// 	fmt.Println(x)
// }

// func main() { //ninja 4.2
// 	x := []int{123, 1234, 1, 12345, 99, 55, 655, 755, 899, 9087}
// 	for i, v := range x {
// 		fmt.Println(i, v)
// 	}
// 	fmt.Printf("Type: %T\n", x)
// }

// func main() { //ninja 4.1
// 	x := [5]int{123, 1234, 1, 12345, 99}
// 	for i, v := range x {
// 		fmt.Println(i, v)
// 	}
// 	fmt.Printf("Type: %T\n", x)
// }
