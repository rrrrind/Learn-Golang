package main

import "fmt"

func main(){
  type User struct{
    name string
    score [3]int
    sum int
  }

  var user1 User

  user1.name = "A"
  user1.score[0] = 10
  user1.score[1] = 40
  user1.score[2] = 60
  user1.sum = (user1.score[0]+user1.score[1]+user1.score[2]) / 3
  
  fmt.Printf("name = %s\n", user1.name)
  fmt.Printf("score = %d\n", user1.score)
  fmt.Printf("sum = %d\n", user1.sum)
}