package main

import "fmt"

func main(){
  var m1 map[string]int
  var m2 map[string]int
  m1= make(map[string]int)
  m2 = make(map[string]int, 10)
  m3 := map[string]int{"x":10, "y":20}
  m4 := m3
  m4["z"] = 30

  fmt.Printf("m1=%d\n", m1)
  fmt.Printf("m2=%d\n", m2)
  fmt.Printf("m3=%d\n", m3["x"])
  fmt.Printf("m4=%d\n", m4["z"])
}