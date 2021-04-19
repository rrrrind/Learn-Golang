package main

import "fmt"

func main(){
  odd_even(10)
}

func odd_even(x int) {
  for i := 1; i<=x; i++ {
    if i % 2 == 0 {
      fmt.Printf("%d is even\n", i)
    } else {
      fmt.Printf("%d is odd\n", i)
    }
  } 
}