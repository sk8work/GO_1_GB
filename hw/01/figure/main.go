// Даны катеты прямоугольного треугольника. Найти его площадь, периметр и гипотенузу.
// Используйте тип данных float64 и функции из пакета math.

package main

import (
	"GO_1_GB/hw/01/figure/triangle"
	"fmt"
)

func main() {
	h := triangle.GetHep(3, 4)
	p := triangle.GetPerimetr(3, 4)
	s := triangle.GetSquare(3, 4)

	fmt.Println(h, p, s)
}
