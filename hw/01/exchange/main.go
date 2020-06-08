package main

// Написать программу для конвертации рублей в доллары. Программа запрашивает сумму в
//  рублях и выводит сумму в долларах. Курс доллара задайте константой.

import (
	"GO_1_GB/hw/01/exchange/conversion"
	"fmt"
)

func main() {

	fmt.Println(conversion.RubToUSD(64000, 0.0164))

}
