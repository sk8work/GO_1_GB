// arrs and slices

package main

import "fmt"

func main() {
	// sl1 := []int{1, 2, 3, 4, 5}
	// fmt.Println(sl1, cap(sl1), len(sl1))
	// for idx, v := range sl1 {
	// 	fmt.Printf("idx: %v| v: %v\n", idx, v)
	// }
	// sl1 = append(sl1, 10)
	// fmt.Println(sl1, cap(sl1), len(sl1))
	// for idx, v := range sl1 {
	// 	fmt.Printf("idx: %v| v: %v\n", idx, v)
	// }
	// return

	sl2 := []int{}
	for i := 0; i <= 10; i++ {
		sl2 = append(sl2, i)
		fmt.Println(sl2, len(sl2), cap(sl2))
	}

}
