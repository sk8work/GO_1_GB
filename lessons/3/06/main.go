// unmarshall different types
package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	str := `["Hello", "World"]`
	var a []string
	json.Unmarshal([]byte(str), &a)
	fmt.Println(a[0])

	strMap := `{"sorok_dva": 42, "sorok_tri": 43, "life is good?": true}`
	mp := make(map[string]interface{})
	json.Unmarshal([]byte(strMap), &mp)
	fmt.Println(mp)
}
