// json more
package main

import (
	"encoding/json"
	"fmt"
)

type A struct {
	Str  string `json:"str"`
	Num  int    `json:"-"`               // не показывать совсем
	Str2 string `json:"str_2,omitempty"` // не показывать, если пустое
}

func main() {
	a := A{
		Str: "string",
		Num: 42,
	}
	data, _ := json.Marshal(a)
	fmt.Println(string(data))
}
