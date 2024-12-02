package main

import (
	"fmt"
	"strings"
)

func main() {
	var a, c string
	fmt.Scan(&a)
	var b = strings.Split(a, "")
	for i := 0; i < len(a)-1; i = i + 2 {
		c = c + string(b[i]) + "*" + string(b[i+1]) + "*"
	}
    c = c + string(b[len(b)-1])
	fmt.Println(c)
}






