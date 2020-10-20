package main

import "fmt"

func main() { //ninja 5.4
	s := struct {
		first     string
		friends   map[string]int
		favDrinks []string
	}{
		first: "James",
		friends: map[string]int{
			"Moneypenny": 333,
			"Q":          1313,
			"M":          888,
		},
		favDrinks: []string{
			"Mt. Dew",
			"Creme Soda",
			"Cherry Coke",
			"Gatorade",
		},
	}
	fmt.Println(s.first)
	fmt.Println(s.friends)
	fmt.Println(s.favDrinks, "\n")
	fmt.Println(s.first)
	for k, v := range s.friends {
		fmt.Println(k, v)
	}
	for _, val := range s.favDrinks {
		fmt.Println(val)
	}
}

// type vehicle struct {
// 	doors int
// 	color string
// }
// type truck struct {
// 	vehicle
// 	fourWheel bool
// }
// type sedan struct {
// 	vehicle
// 	luxury bool
// }

// func main() { //ninja 5.3

// 	chevy := truck{
// 		vehicle: vehicle{
// 			doors: 4,
// 			color: "black"},
// 		fourWheel: true,
// 	}
// 	bmw := sedan{
// 		vehicle: vehicle{
// 			doors: 4,
// 			color: "red"},
// 		luxury: true,
// 	}
// 	fmt.Println(bmw)
// 	fmt.Println(chevy)
// 	fmt.Println(chevy.doors, "\t", chevy.color, "\tFour Wheel Drive:", chevy.fourWheel)

// }

// func main() { //ninja 5.2

// 	type person struct {
// 		first      string
// 		last       string
// 		ffIceCream []string
// 	}

// 	p1 := person{
// 		first:      "James",
// 		last:       "Bond",
// 		ffIceCream: []string{"champagne", "olive", "pizza"},
// 	}
// 	p2 := person{
// 		first:      "Money",
// 		last:       "Penny",
// 		ffIceCream: []string{"strawberry", "vanilla", "rabbit tracks"},
// 	}
// 	m := map[string]person{
// 		p1.last: p1,
// 		p2.last: p2,
// 	}

// 	fmt.Println(m)
// 	for _, v := range m {
// 		fmt.Println(v.first)
// 		fmt.Println(v.last)
// 		for _, val := range v.ffIceCream {
// 			fmt.Println("\t", val)
// 		}
// 	}
// }

// func main() { //ninja 5.1
// 	type person struct {
// 		first      string
// 		last       string
// 		ffIceCream []string
// 	}
// 	p1 := person{
// 		first:      "James",
// 		last:       "Bond",
// 		ffIceCream: []string{"champagne", "olive", "pizza"},
// 	}
// 	p2 := person{
// 		first:      "Money",
// 		last:       "Penny",
// 		ffIceCream: []string{"strawberry", "vanilla", "rabbit tracks"},
// 	}

// 	fmt.Println(p1.first, p1.last)
// 	for _, v := range p1.ffIceCream {
// 		fmt.Println("\t", v)
// 	}
// 	fmt.Println(p2.first, p2.last)
// 	for _, v := range p2.ffIceCream {
// 		fmt.Println("\t", v)
// 	}
// }
