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