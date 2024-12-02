package main

import (
	"fmt"
	"strconv"
)

func reverse(str string) (result string) {
	for _, v := range str {
		result = string(v) + result
	}
	return
}
func main() {
	var a int
	fmt.Scan(&a)
	t := strconv.Itoa(a)
	var r string
	for i := 0; i < len(t); i++ {
		var g = a % 10
		a = a / 10
		r = r + reverse(strconv.Itoa(g*g))
	}
	r2 := reverse(r)
	fmt.Println(r2)
}



