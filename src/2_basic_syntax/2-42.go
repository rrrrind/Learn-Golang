package main

import "fmt"

func main(){
  for i:=0; i<100; i++{
    if i % 2 == 0 {
      fmt.Printf("%0.3dは偶数だね!\n",i)
    }else if i % 2 == 1 {
      fmt.Printf("%0.3dは奇数だね!\n",i)
    }
  }
}
