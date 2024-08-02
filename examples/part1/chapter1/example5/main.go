package main

import (
	"fmt"
)

func main() {
	var x int = 3
	fmt.Printf("Type of x: %T Value of x: %v\n", x, x)
	fmt.Println(" 3 in binary is 11")
	x = x << 1
	fmt.Printf("After 1 left shift - Type of x: %T Value of x: %v\n", x, x)
	fmt.Println(" 3 in binary is 11, in binary 110 is [2^2 + 2^1 + 0 = 5] in decimal")
	x = 3
	x = x << 2
	fmt.Printf("After 2 left shift - Type of x: %T Value of x: %v\n", x, x)
	x = 3
	x = x << 3
	fmt.Printf("After 3 left shift - Type of x: %T Value of x: %v\n", x, x)
	x = 3
	x = x << 4
	fmt.Printf("After 4 left shift - Type of x: %T Value of x: %v\n", x, x)
	x = 3
	x = x << 5
	fmt.Printf("After 5 left shift - Type of x: %T Value of x: %v\n", x, x)
	x = 3
	x = x << 6
	fmt.Printf("After 6 left shift - Type of x: %T Value of x: %v\n", x, x)
	x = 3
	x = x << 7
	fmt.Printf("After 7 left shift - Type of x: %T Value of x: %v\n", x, x)
	x = 3
	x = x << 8
	fmt.Printf("After 8 left shift - Type of x: %T Value of x: %v\n", x, x)
	x = 3
	x = x << 9
	fmt.Printf("After 9 left shift - Type of x: %T Value of x: %v\n", x, x)
	x = 3
	x = x << 10
	fmt.Printf("After 10 left shift - Type of x: %T Value of x: %v\n", x, x)

	/// ---
	fmt.Println("/////////////////////////////////////////")
	var y int = 3000
	fmt.Printf("Type of y: %T Value of y: %v\n", y, y)
	y = y >> 1
	fmt.Printf("After 1 left right - Type of y: %T Value of y: %v\n", y, y)
	y = 3000
	y = y >> 2
	fmt.Printf("After 2 left right - Type of y: %T Value of y: %v\n", y, y)
	y = 3000
	y = y >> 3
	fmt.Printf("After 3 left right - Type of y: %T Value of y: %v\n", y, y)
	y = 3000
	y = y >> 4
	fmt.Printf("After 4 left right - Type of y: %T Value of y: %v\n", y, y)
	y = 3000
	y = y >> 5
	fmt.Printf("After 5 left right - Type of y: %T Value of y: %v\n", y, y)
	y = 3000
	y = y >> 6
	fmt.Printf("After 6 left right - Type of y: %T Value of y: %v\n", y, y)
	y = 3000
	y = y >> 7
	fmt.Printf("After 7 left right - Type of y: %T Value of y: %v\n", y, y)


	/// ---
	fmt.Println("/////////////////////////////////////////")
	var z int = 1
	fmt.Printf("Type of z: %T Value of z: %v\n", z, z)
	z = z << 1
	fmt.Printf("After 1 left shift - Type of z: %T Value of z: %v\n", z, z)
	z = 1
	z = z << 2
	fmt.Printf("After 2 left shift - Type of z: %T Value of z: %v\n", z, z)
	z = 1
	z = z << 3
	fmt.Printf("After 3 left shift - Type of z: %T Value of z: %v\n", z, z)
	z = 1
	z = z << 4
	fmt.Printf("After 4 left shift - Type of z: %T Value of z: %v\n", z, z)
	z = 1
	z = z << 5
	fmt.Printf("After 5 left shift - Type of z: %T Value of z: %v\n", z, z)
	z = 1
	z = z << 6
	fmt.Printf("After 6 left shift - Type of z: %T Value of z: %v\n", z, z)
	z = 1
	z = z << 7
	fmt.Printf("After 7 left shift - Type of z: %T Value of z: %v\n", z, z)
	fmt.Println("After 7 left shift - [10000000] in binary is 2^7 = 2 * 8 * 8 = 128")
	z = 1
	z = z << 10
	fmt.Printf("After 10 left shift - Type of z: %T Value of z: %v\n", z, z)
	fmt.Println("After 10 left shift - [10000000] in binary is 2^10 = 2^7 * 2^3 = 128 * 8 = 256 * 4 = 1024 ")
	z = 1
	z = z << 100
	fmt.Printf("After 100 left shift - Type of z: %T Value of z: %v\n", z, z)

}
