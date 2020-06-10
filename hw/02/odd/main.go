// Написать функцию, которая определяет, четное ли число.

package main

import (
	"errors"
	"fmt"
)

func isOdd(a int) (bool, error) {
	if a%2 == 0 {
		return true, nil
	} else if a%2 != 0 {
		return false, nil
	}
	defer fmt.Println("Паника обработана")
	return false, errors.New("Wrong")
}

func main() {
	num := 4
	result, err := isOdd(num)
	fmt.Printf("is %v odd? - %v| err - %v\n", num, result, err)
}
