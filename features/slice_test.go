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

/**
slice新切片 长度=high-low 容量=max-low
 */
func TestSplit(t *testing.T)  {
	slice1 := make([]int, 0)
	slice2 := make([]int, 1, 3)
	slice3 := []int{}
	slice4 := []int{1: 2, 3}
	arr := []int{1, 2, 3}
	slice5 := arr[1:2]
	slice6 := arr[1:2:2]
	slice7 := arr[1:]
	slice8 := arr[:1]
	slice9 := arr[3:]

	fmt.Printf("%s = %v,\t len = %d, cap = %d\n", "slice1", slice1, len(slice1), cap(slice1))
	fmt.Printf("%s = %v,\t len = %d, cap = %d\n", "slice2", slice2, len(slice2), cap(slice2))
	fmt.Printf("%s = %v,\t len = %d, cap = %d\n", "slice3", slice3, len(slice3), cap(slice3))
	fmt.Printf("%s = %v,\t len = %d, cap = %d\n", "slice4", slice4, len(slice4), cap(slice4))
	fmt.Printf("%s = %v,\t len = %d, cap = %d\n", "slice5", slice5, len(slice5), cap(slice5))
	fmt.Printf("%s = %v,\t len = %d, cap = %d\n", "slice6", slice6, len(slice6), cap(slice6))
	fmt.Printf("%s = %v,\t len = %d, cap = %d\n", "slice7", slice7, len(slice7), cap(slice7))
	fmt.Printf("%s = %v,\t len = %d, cap = %d\n", "slice8", slice8, len(slice8), cap(slice8))
	fmt.Printf("%s = %v,\t len = %d, cap = %d\n", "slice9", slice9, len(slice9), cap(slice9))
}