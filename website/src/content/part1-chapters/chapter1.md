---
title: Chapter 1
# slug: my-custom-slug/supports/slashes
tags:
  - golang
  - basics

summary: First Part, Chapter 1, The basics of Golang
index: 0
---

## The Basics of Golang

### Simplest app, one source file and entry point

To build a first executable app in Golang, with one source file and one entrypoint, this is the code that you need:

```Golang
package main

import "fmt"


func main() {

    fmt.Println("Hello Golang App!")
}
```

Add the above content to a single file with an `*.go` extension, then you can build and run it like this:

```bash
# ---
#  Build the executable, the [-o] allows you to tell the compiler the path to the executable file to generate:
go build -o app.exe *.go
# ---
# and well execute it like this:
./app.exe
```

### Primary types in Golang

Here an example of basic operators on primary Golang types  string, booleans, integers, and decimal numbers:

```Golang
package main

import "fmt"


func main() {

    fmt.Println("go" + "lang")

    fmt.Println("1+1 =", 1+1)
    fmt.Println("7.0/3.0 =", 7.0/3.0)
    fmt.Println("7.12548789654/3.0 =", 7.12548789654/3.0)
    fmt.Println("7.12548789654*3.0 =", 7.12548789654*3.0)

    fmt.Println(true && false)
    fmt.Println(true || false)
    fmt.Println(!true)

}
```

### Simple variables

The `var` keyword is used in GOlang to declare variables.
It can be used inside a function declaration, or as top declaration like imports. Note that if a variable is declared both at top of a golang source file, at same level than the imports, and inside a function, then the declaration done inside the function supersedes.

Note that variables declared, if not initialized, always have a default initialization value, called a "_zero value_":

* `""` The empty string for strings (`string`)
* `0` for integers (`int`)
* `false` for boolean

Here below examples of how to declare variables (and also how to display value, and type, of a variable):

```Golang
package main

import (
 "fmt"
 "math/cmplx"
)

var (
 myNb1, myNb2 float32 = 17.0, 13.0
    myNb5, myNb6 float64 = -27.12548789654, 43.46868468464543541534135416354
    myBool1, myBool2 bool = true, false
 myGreatestUnsigned64Integer uint64     = 1<<64 - 1
    /// + ---
    // if you uncomment [myMoreThanGreatestUnsigned64Integer]
    // the source code does not pass compilation checks.
 // myMoreThanGreatestUnsigned64Integer uint64     = 1<<64
 myComplexNumber      complex128 = cmplx.Sqrt(-5 + 12i)
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
}
```

Output of the above will be :

```bash
$ go build -o app.exe *.go && ./app.exe
golang
myInt1+myInt2 = 2
myNb1/myNb2 = 2.3333333
myNb3/myNb4 = 2.054233389411802
myNb3*myNb4 = 24.71607073735472
myNb5/myNb6 = -0.6240236642384905
myNb5*myNb6 = -1179.1092802918636
myBool1 && myBool2 = false
myBool1 || myBool2 = true
!myBool1 =  false
Type: bool Value: true
Type: uint64 Value: 18446744073709551615
Type: complex128 Value: (2+3i)
0 0 false ""

```

The `int`, `uint`, and `uintptr` types are usually 32 bits wide on 32-bit systems and 64 bits wide on 64-bit systems. When you need an integer value you should use `int` unless you have a specific reason to use a sized or unsigned integer type.

Declaring variables can also be done without the var keyword, like this:

```Golang
myvar8 := 42
myvar9 := "Golang rocks"
myvar10 := "Everybody knows"
```

### Type conversions

The expression `T(v)` converts the value `v` to the type `T`.

A few examples of type conversions can be seen in example4, for primary types. Yet, We will see if the same syntax can be used to convert between complex Types I will define, like Upcasting Downcasting between "siblings" types (eg classes with _"inherits"_ relation) in other languages.

_NEXT TODO_: <https://go.dev/tour/basics/14>
