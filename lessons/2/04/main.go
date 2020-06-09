// functions func name(argument type) (argument type) {}

package main

import (
	"fmt"
	"math"
)

func funcName(b bool) (bool, int, string) {
	return b, 42, "42"
}

func calculateParameters(a, b float64, callBackFunc func()) (float64, float64) {
	c := math.Sqrt(a*a + b*b)
	callBackFunc()
	return c, a + b + c
}

func main() {
	condition, world, str := funcName(true)
	fmt.Println(condition, world, str)

	fmt.Println(calculateParameters(3, 4, func() { fmt.Println("Hello world") }))
}
