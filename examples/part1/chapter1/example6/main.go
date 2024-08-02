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

}