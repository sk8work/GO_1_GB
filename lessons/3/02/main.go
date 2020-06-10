// maps

package main

import (
	"fmt"
	"hash/crc32"
)

type hashMap [16]int

func hash(key string) uint32 {
	return crc32.ChecksumIEEE([]byte(key))
}

func findPosition(key string, length int) int {
	h := hash(key)
	return int(h % uint32(length))
}

func main() {
	// sl := make([]int, 4)
	// key := "some_key"
	// position := findPosition(key, len(sl))
	// fmt.Println(position)
	// fmt.Println(sl)

	mp := make(map[int]string)
	fmt.Println(mp)
	mp[1] = "odin"
	mp[2] = "dva"
	fmt.Println(mp)
}
