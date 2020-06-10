// Написать функцию, которая определяет, делится ли число без остатка на

package main

import (
	"errors"
	"fmt"
)

func isDivByThree(a int) (bool, error) {
	if a%3 == 0 {
		return true, nil
	} else if a%3 != 0 {
		return false, nil
	}
	defer fmt.Println("Паника обработана")
	return false, errors.New("Wrong")
}

func main() {
	num := 8
	result, err := isDivByThree(num)
	fmt.Printf("is %v odd? - %v| err - %v\n", num, result, err)
}
