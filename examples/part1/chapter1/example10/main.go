package main

import "fmt"

func zeroval(ival int) {
	ival = 0
}

func zeroptr(iptr *int) {
	*iptr = 0
}

func main() {
	i := 1
	fmt.Println("initial:", i)

	zeroval(i)
	fmt.Println("zeroval:", i)

	zeroptr(&i)
	fmt.Println("zeroptr:", i)

	fmt.Println("pointer:", &i)

	var myVar int = 42

	var thisIsAFirstGeneratedPointerToMyVar = &myVar
	var thisIsASecondGeneratedPointerToMyVar = &myVar

	fmt.Printf("thisIsAFirstGeneratedPointerToMyVar = %v \n", thisIsAFirstGeneratedPointerToMyVar)
	fmt.Printf("thisIsASecondGeneratedPointerToMyVar = %v \n", thisIsASecondGeneratedPointerToMyVar)

	fmt.Println("read [myVar] through the [thisIsAFirstGeneratedPointerToMyVar] pointer:")
	fmt.Println(*thisIsAFirstGeneratedPointerToMyVar) // read [myVar] through the [thisIsAFirstGeneratedPointerToMyVar] pointer
	*thisIsASecondGeneratedPointerToMyVar = 21        // set myVar through the [thisIsASecondGeneratedPointerToMyVar] pointer
	fmt.Printf("after setting [myVar] through the [thisIsASecondGeneratedPointerToMyVar] pointer, myVar = %v \n", myVar)
}
