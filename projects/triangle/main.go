package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c float64
	fmt.Scan(&a, &b)
	c = a*a + b*b
	var y = math.Sqrt(c)
	fmt.Println(y)
}
