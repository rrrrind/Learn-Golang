package main

import "fmt"

func main(){
  for i:=1; i<=100; i++{
    if i % 15 == 0 {
      fmt.Printf("%0.3d fizz buzz!\n",i)
    }else if i % 3 == 0 {
      fmt.Printf("%0.3d fizz!\n",i)
    } else if i % 5 == 0 {
      fmt.Printf("%0.3d buzz!\n",i)
    } else {
      fmt.Printf("%0.3d\n",i)
    }
   }
}
