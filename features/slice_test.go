package features

import (
	"fmt"
	"testing"
)

func TestMakeAddr(t *testing.T)  {
	a := []int{1,2,3,4}
	fmt.Printf("%p \n",&a)
	fmt.Println("update before",a)
	update(a)
	fmt.Println("update after",a)
}

func update(a []int)  {
	fmt.Printf("%p \n",&a)
	a[1]  = 10
	fmt.Println("update",a)
}