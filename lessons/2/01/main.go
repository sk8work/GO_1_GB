// for-loops

package main

import (
	"fmt"
)

func main() {

	// basic loop
	// for i := 0; i <= 10; i++ {
	// 	fmt.Println(i)
	// }
	// return

	// for-loop with entered int number
	// var a int
	// fmt.Println("Enter int number: ")
	// fmt.Scan(&a)
	// for i := 0; i <= a; i++ {
	// 	fmt.Println(i)
	// }
	// return

	s := "Hello мир!"
	for idx, v := range s {
		fmt.Println(idx, v)
	}
}
