package main

import (
	"fmt"
	"testing"
)

type T struct {
	A string
	B []string
}

/**
slice是浅拷贝
 */
func TestCopy(t *testing.T)  {
	x := T{"hammer",[]string{"on"}}
	y := x
	y.A = "hansir"
	y.B[0] = "off"
	fmt.Printf("x value %v \n",x)
	fmt.Printf("y value %v \n",y)

	fmt.Printf("x addr %p \n",&x)
	fmt.Printf("y addr %p \n",&y)

	fmt.Printf("x.B[0] addr %p \n",&x.B[0])
	fmt.Printf("y.B[0] addr %p \n",&y.B[0])
}
