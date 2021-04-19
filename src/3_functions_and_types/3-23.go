package main

import "fmt"

func main(){
  var ns1 [5]int
  var ns2 = [5]int{10,20,30,40,50}
  var ns3 = [...]int{10,20,30,40,50,60}
  ns4 := [...]int{5:50, 10:100}
  
  var ns5 []int
  ns5 = make([]int, 3, 10)

  fmt.Printf("ns1=%d\n", ns1)
  fmt.Printf("ns2=%d\n", ns2)
  fmt.Printf("ns3=%d\n", ns3)
  fmt.Printf("ns4=%d\n", ns4)
  fmt.Printf("ns5=%d\n", ns5)
}