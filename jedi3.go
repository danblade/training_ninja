package main

import (
	"fmt"
	"os"
)

//ninja 3.9
func main() {
	favSport := "Snowboarding"

	switch favSport {
	case "skiing":
		os.Exit(0)
	case "Snowboarding":
		os.Exit(0)

	default:
		fmt.Println("Go to the pool")
	}

}

// //ninja 3.8
// func main() {
// 	bd := 1980
// 	for {
// 		fmt.Println(bd)

// 		switch bd {
// 		case 2020:
// 			os.Exit(0)

// 		default:
// 			bd++
// 		}
// 	}
// }

// //ninja 3.7
// func main() {
// 	bd := 1980
// 	for {
// 		if bd <= 2020 {
// 			fmt.Println(bd)
// 			bd++
// 		} else {
// 			break
// 		}
// 	}
// }

// //ninja 3.6
// func main() {
// 	bd := 1980
// 	for {
// 		if bd <= 2020 {
// 			fmt.Println(bd)
// 			bd++
// 		} else {
// 			break
// 		}
// 	}
// }

// //ninja 3.5
// func main() {
// 	for i := 10; i <= 100; i++ {
// 		fmt.Println(i, "remainder:", i%4)

// 	}
// }

// //ninja 3.4
// func main() {
// 	bd := 1980
// 	for {
//	fmt.Println(bd)
// 			bd++
// 		if bd <= 2020 {
// 			break
// 		}
// 	}
// }

// //ninja 3.3
// func main() {
// 	bd := 1980
// 	for bd <= 2020 {
// 		fmt.Println(bd)
// 		bd++
// 	}
// }

// //ninja 3.2
// func main() {
// 	for i := 0; i <= 3; i++ {
// 		for j := 65; j <= 90; j++ {
// 			fmt.Printf("\t%#U\n", j)
// 		}
// 		fmt.Println("End of iteration: ", i)
// 	}
// }

// //ninja 3.1
// func main() {
// 	for i := 0; i < 10001; i++ {
// 		fmt.Printf("%6d\n", i)
// 	}
// }

// //exercise for video 60
// func main() {
// 	for i := 0; i < 201; i++ {
// 		fmt.Printf("Decimal: %d, Unicode: %#[1]U.\n", i)
// 	}
//}
