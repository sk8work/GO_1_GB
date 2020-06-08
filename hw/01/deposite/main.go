// * Пользователь вводит сумму вклада в банк и годовой процент. Найти сумму вклада через 5 лет.

package main

import (
	"GO_1_GB/hw/01/deposite/calcdepo"
	"fmt"
)

func main() {
	myMoney := calcdepo.GetDepo(1000, 0.2, 12)
	fmt.Printf("%v\n", myMoney)
}
