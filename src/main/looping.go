package main

import (
  "fmt"
)

func main() {
  i := 1

  //do while
  for i <= 3{
  	fmt.Println(i)
  	i = i + 1
  }

  //for
  for j := 7; j<=9; j++{
  	fmt.Println(j)
  }

  //infinite
  for{
  	fmt.Println("loop")
  	break
  }
}
