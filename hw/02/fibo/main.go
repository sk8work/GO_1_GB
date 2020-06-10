// Написать функцию, которая последовательно выводит на экран 100 первых чисел Фибоначчи, начиная от 0.

package main

import (
	"fmt"
)

func getFiboNums(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 || n == 2 {
		return 1
	}
	return getFiboNums(n-1) + getFiboNums(n-2)
}

func main() {
	n := 40
	for i := 0; i <= n; i++ {
		fmt.Println(getFiboNums(i))
	}
}
