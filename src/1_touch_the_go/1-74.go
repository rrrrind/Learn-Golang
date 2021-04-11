package main

import "fmt"

func main(){
  var price int
  fmt.Printf("値段>")
  fmt.Scan("&price")
  fmt.Printf("%d円\n", price)
}