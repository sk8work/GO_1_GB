// json
package main

import (
	"GO_1_GB/lessons/3/04/contact"
	"encoding/json"
	"fmt"
	"log"
)

func backToRow(j string) (contact.Person, error) {
	var p contact.Person
	err := json.Unmarshal([]byte(j), &p)
	if err != nil {
		return contact.Person{}, err
	}
	return p, nil
}

func main() {
	s := contact.Person{
		Name:     "name",
		Lastname: "Lastname",
	}

	data, err := json.Marshal(s)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(data)
	fmt.Println(string(data))

	fmt.Println(backToRow(`{"Name": "Petya", "Lastname": "Ivanov"}`))
}
