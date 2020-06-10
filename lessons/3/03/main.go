// struct

package main

import (
	"GO_1_GB/lessons/3/03/contact"
	"fmt"
)

func main() {
	person1 := contact.Contact{}
	fmt.Println(person1)
	person1.Name = "Eugene"
	person1.Mobile = 85133214654
	person1.SetsuccessfulPerson(true)
	fmt.Printf("%+v", person1)
}
