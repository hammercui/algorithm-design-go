package features

import (
	"fmt"
	"testing"
)

/**
泛型测试
*/

func TestFindIndex(t *testing.T) {

	sel := 1
	arr := []int{1, 2, 3, 4, 5}
	pos := 0
	if pos == findIndex(arr, sel) {
		t.Log("success")
	} else {
		t.Fatal("findIndex error")
	}
}

func TestContacts(t *testing.T) {
	a := []int{1, 2}
	b := []int{3, 4}
	fmt.Println(concatenate(a, b))
}

// func TestContactsLint(t *testing.T) {
// 	c := []string{"a", "b"}
// 	d := []string{"c", "d"}
// 	fmt.Println(concatenate(c, d))
// }
