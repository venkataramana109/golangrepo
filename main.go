package main

import (
	"fmt"
)
func main()  {
	i := 10
	j := 11
	z := i + j
	fmt.Println("the value of z is :=",z)
}
// Adding two numbers
func AddTwoNumbers(n1 ,n2 int){
	var n3 int
	n3 = n1 +n2
	fmt.Println("value after adding two numbers :", n3)
}