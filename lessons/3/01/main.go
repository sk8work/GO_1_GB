// arrs

package main

import "fmt"

func main() {
	var arr [1]int
	fmt.Println(arr, len(arr), cap(arr))
	fmt.Printf("%T\n", arr)

	arr1 := [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr1, len(arr1), cap(arr1))
	fmt.Printf("%T\n", arr1)

	emptyArr := [5]int{0}
	fmt.Println(emptyArr, len(emptyArr), cap(emptyArr))
	fmt.Printf("%T\n", emptyArr)

	sl1 := []int{1, 2, 3, 4, 5}
	sl1 = append(sl1, 6)
	fmt.Println(sl1, len(sl1), cap(sl1))
	fmt.Printf("%T\n", sl1)

	sl2 := make([]int, 5, 5)
	fmt.Println(sl2, len(sl2), cap(sl2))
	fmt.Printf("%T\n", sl2)
}
