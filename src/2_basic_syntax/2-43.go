package main

import "fmt"
import "time"
import "math/rand"

func main(){
  t := time.Now().UnixNano()
  rand.Seed(t)
  s := rand.Intn(6)
  fmt.Println(s)
  
  if s == 0 {
    fmt.Printf("大吉!\n")
  } else if s % 5 == 0 || s % 4 == 0 {
    fmt.Printf("中吉!\n")
  } else if s % 3 == 0 || s % 2 == 0 {
    fmt.Printf("吉!\n")
  } else if s == 1 {
    fmt.Printf("凶!\n")
  }
}
