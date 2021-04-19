package main

import "fmt"

type MyInt int
func (n *MyInt) Inc() { (*n)++ }

func main(){
  var m MyInt
  fmt.Println(m)
  m.Inc()
  fmt.Println(m)
}

