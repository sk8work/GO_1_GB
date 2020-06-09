// PANIC RECOVER

package main

import (
	"errors"
	"fmt"
)

func div(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("Zero Divizion")
	}
	return a / b, nil
}

func main() {
	c, err := div(5, 0)
	fmt.Println(c, err)
	c, err = div(5, 2)
	fmt.Println(c, err)
}
