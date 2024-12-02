package main

import (
	"fmt"
	"strings"
)

func reverse(str string) (result string) {
	for _, v := range str {
		result = string(v) + result
	}
	return
}
func main() {
	var a string
	fmt.Scan(&a)
	var array = strings.Split(a, "")
	var maxim = 00
	var intValue = 0
	for i := 0; i < len(a); i++ {
		fmt.Sscan(string(array[i]), &intValue)
		if intValue > maxim {
			maxim = intValue
		}
	}
	fmt.Println(maxim)
}
