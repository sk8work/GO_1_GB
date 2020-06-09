// if-else, switch

package main

import "fmt"

func f() bool {
	return true
}

func main() {
	// if f() {
	// 	fmt.Println("Hello")
	// } else {
	// 	fmt.Println("World")
	// }
	// return

	i := 2
	switch i {
	case 1:
		fmt.Println(1)
	case 2:
		fmt.Println(2)
		fallthrough
	case 3:
		fmt.Println(3)
	}
}
