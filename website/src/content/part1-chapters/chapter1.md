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

# ---
# note you can prettify/format your go code by running:
go fmt *.go
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

### Type inference

When declaring a variable without specifying an explicit type (either by using the `:=` syntax or `var =` expression syntax), the variable's type is inferred from the value on the right hand side.

### Constants

Syntax of declaring a constant is:

```Golang
package main

import "fmt"

const Pi = 3.14

func main() {
 const World = "世界"
 fmt.Println("Hello", World)
 fmt.Println("Happy", Pi, "Day")

 const Truth = true
 fmt.Println("Go rules?", Truth)
}
```

* Constants can be character, string, boolean, or numeric values.
* Constants cannot be declared using the `:=` syntax.

### For loops

Here is a classicla for loop syntax in golang:

```Golang
package main

import "fmt"

func main() {
 // first loop example
 sum := 0
 for i := 0; i < 10; i++ {
  sum += i
 }
 fmt.Println(sum)

 // init statement is optional
 sum2 := 1
 for ; sum2 < 1000; {
  sum2 += sum2
 }
 fmt.Println(sum2)

 // same than above, simplified with less semi-colon
 sum3 := 1
 for sum3 < 1000 {
  sum3 += sum3
 }
 fmt.Println(sum3)
 
 // below this is an infinite loop (don't do that!)
 for {
 }

}
```

### IF statement

* Syntax of IF statements:

```Golang
package main

import (
 "fmt"
 "math"
)

func sqrt(x float64) string {
 if x < 0 {
  return sqrt(-x) + "i"
 }
 return fmt.Sprint(math.Sqrt(x))
}

func pow(x, n, lim float64) float64 {
 /**
  * Like for, the if statement can start with a
  * short statement to execute before the condition.
  */
 if alpha := math.Pow(x, n); alpha < lim {
  return alpha
 }
 // fmt.Printf("// variable 'alpha' does not exist here: %v", alpha)
 return lim
}

func pow2(x, n, lim float64) float64 {
 if v := math.Pow(x, n); v < lim {
  return v
 } else {
  fmt.Printf("%g >= %g\n", v, lim)
 }
 // can't use v here, though
 return lim
}

func main() {
 fmt.Println(sqrt(2), sqrt(-4))
 fmt.Println(
  pow(3, 2, 10),
  pow(3, 3, 20),
  pow2(3, 3, 20),
 )
}
```

### Switch/case statement

* The switch statement syntax in Golang does not need the `break` keyword, like in many other languages:

```Golang
package main

import (
 "fmt"
 "runtime"
    "time"
)

func main() {
 fmt.Print("Go runs on ")
 switch os := runtime.GOOS; os {
 case "darwin":
  fmt.Println("OS X.")
 case "linux":
  fmt.Println("Linux.")
 case "windows":
  fmt.Println("OS X.")
 default:
  // freebsd, openbsd,
  // plan9, windows...
  fmt.Printf("%s.\n", os)
 }
    /// another switch syntax:

 fmt.Println("When's Saturday?")
 today := time.Now().Weekday()
 switch time.Saturday {
 case today + 0:
  fmt.Println("Today.")
 case today + 1:
  fmt.Println("Tomorrow.")
 case today + 2:
  fmt.Println("In two days.")
 default:
  fmt.Println("Too far away.")
 }

    /////
    // a third syntax of switch statement, neat
    // way to write an if with many else:
 t := time.Now()
 switch {
 case t.Hour() < 12:
  fmt.Println("Good morning!")
 case t.Hour() < 17:
  fmt.Println("Good afternoon.")
 default:
  fmt.Println("Good evening.")
 }
}
```

### The `defer` keyword (and Panic and Recover)

A `defer` statement defers the execution of a function until the surrounding function returns.

The deferred call's arguments are evaluated immediately, but the function call is not executed until the surrounding function returns.

Deferred function calls are pushed onto a stack. When a function returns, its deferred calls are executed in last-in-first-out order.

For example:

```Golang
package main

import "fmt"

func main() {
    // will print "world", only after the 
    // [main] funtion returns
 defer fmt.Println("after world 2")
 defer fmt.Println("after world 1")
 defer fmt.Println("world")

 fmt.Println("hello")
}
```

Defer is commonly used to simplify functions that perform various clean-up actions. typical example:

```Golang
package main

import (
 "fmt"
 "os"
 "io"
)

/**
 * The below function always returns 2.
 * This function increments the 
 * return value "i" AFTER the surrounding function returns.
 * Thus, this function returns 2.
 * This is convenient for modifying the error return value 
 * of a function;
 */
func c() (i int) {
    defer func() { i++ }()
    return 1
}
func f() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recovered in f", r)
        }
    }()
    fmt.Println("Calling g.")
    g(0)
    fmt.Println("Returned normally from g.")
}

func g(i int) {
    if i > 3 {
        fmt.Println("Panicking!")
        panic(fmt.Sprintf("%v", i))
    }
    defer fmt.Println("Defer in g", i)
    fmt.Println("Printing in g", i)
    g(i + 1)
}

func CopyFile(dstName, srcName string) (written int64, err error) {
    src, err := os.Open(srcName)
    if err != nil {
        return
    }
    //////
    // below deferred call, is useful: if the os.Create 
    // function call fails, then the file will
    // still be closed, after copyFile function returns.
    defer src.Close()

    dst, err := os.Create(dstName)
    if err != nil {
        return
    }
    defer dst.Close()

    return io.Copy(dst, src)
}

func main() {
    // will print "world", only after the 
    // [main] funtion returns
 defer fmt.Println("after world 2")
 defer fmt.Println("after world 1")
 defer fmt.Println("world")

 fmt.Println("hello")

    CopyFile("./README.md", "README.copied.md")
    f()
    fmt.Println("Returned normally from f.")
 fmt.Printf("The c() function returns: %v \n", c())
}

```

**`Panic`** is a built-in function that stops the ordinary flow of control and begins panicking. When the function `F` calls panic, execution of `F` stops, any deferred functions in `F` are executed normally, and then `F` returns to its caller. To the caller, `F` then behaves like a call to panic. The process continues up the stack until all functions in the current goroutine have returned, at which point the program crashes. Panics can be initiated by invoking panic directly. They can also be caused by runtime errors, such as out-of-bounds array accesses.

**`Recover`** is a built-in function that regains control of a panicking goroutine. Recover is only useful inside deferred functions. During normal execution, a call to recover will return nil and have no other effect. If the current goroutine is panicking, a call to recover will capture the value given to panic and resume normal execution.

In Golang, funtion paramters are passed by values, unless its type is prefixed by the `*` parameter, which then means the function expects to be passed a pointer, not a bare value. If a variable `var myVar int = 42` is declared, you can generate a pointer to that variable's value using the ampersand operator, like this: `&myVar`.

The `*` operator denotes the pointer's underlying value. Using the `*` is known as "dereferencing" or "indirecting".

```Golang
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

```

The pointer of a variable is the address, in memory, where the value of that variable is stored.

Unlike C, Go has no pointer arithmetic.

_NEXT TODO_: <https://go.dev/tour/moretypes/2>
