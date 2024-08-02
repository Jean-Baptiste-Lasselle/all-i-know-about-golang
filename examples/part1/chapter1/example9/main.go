package main

import (
	"fmt"
	"io"
	"os"
)

/**
 * The below function always returns 2.
 * This function increments the
 * return value "i" AFTER the surrounding function returns.
 * Thus, this function returns 2.
 * This is convenient for modifying the error return value of a function;
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
