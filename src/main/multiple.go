//multiple return value
package main

import (
  "fmt"
)

func vals()(int, int){
	return 3, 7
}

func main() {
	a, b:= vals() 
	fmt.Println(a)
	fmt.Println(b)

	//"_" output yang pertama tdk dimunculkan
	_, c := vals()
	fmt.Println(c)
}
