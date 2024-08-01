package main

import (
	"fmt"
	"math/cmplx"
)

var (
	myNb1, myNb2                float32 = 17.0, 13.0
	myNb5, myNb6                float64 = -27.12548789654, 43.46868468464543541534135416354
	myBool1, myBool2            bool    = true, false
	myGreatestUnsigned64Integer uint64  = 1<<64 - 1
	/// + ---
	// if you uncomment [myMoreThanGreatestUnsigned64Integer]
	// the source code does not pass compilation checks.
	// myMoreThanGreatestUnsigned64Integer uint64     = 1<<64
	myComplexNumber complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {

	var myWord1, myWord2 string = "go", "lang"
	fmt.Println(myWord1 + myWord2)

	var myInt1, myInt2 int = 1, 1
	fmt.Println("myInt1+myInt2 =", myInt1+myInt2)
	var myNb1, myNb2 float32 = 7.0, 3.0
	fmt.Println("myNb1/myNb2 =", myNb1/myNb2)
	var myNb3, myNb4 float64 = 7.12548789654, 3.46868468464543541534135416354
	fmt.Println("myNb3/myNb4 =", myNb3/myNb4)
	fmt.Println("myNb3*myNb4 =", myNb3*myNb4)

	fmt.Println("myNb5/myNb6 =", myNb5/myNb6)
	fmt.Println("myNb5*myNb6 =", myNb5*myNb6)

	fmt.Println("myBool1 && myBool2 =", myBool1 && myBool2)
	fmt.Println("myBool1 || myBool2 =", myBool1 || myBool2)
	fmt.Println("!myBool1 = ", !myBool1)

	fmt.Printf("Type: %T Value: %v\n", myBool1, myBool1)
	fmt.Printf("Type: %T Value: %v\n", myGreatestUnsigned64Integer, myGreatestUnsigned64Integer)
	fmt.Printf("Type: %T Value: %v\n", myComplexNumber, myComplexNumber)

	///

	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)

	myvar8 := 42
	myvar9 := "Golang rocks"
	myvar10 := "Everybody knows"
	myvar14 := "knows"

	fmt.Printf("%v %v %v %q\n", myvar8, myvar9, myvar10, myvar14)
}
