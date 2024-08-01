package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, z)
	fmt.Printf("Type of x: %T Valueof x: %v\n", x, x)
	fmt.Printf("Type of y: %T Valueof y: %v\n", y, y)
	fmt.Printf("Type of z: %T Valueof z: %v\n", z, z)
	fmt.Printf("Type of f: %T Valueof f: %v\n", f, f)
}
