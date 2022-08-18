package main

import (
	"fmt"
	"testing"
)

/**
泛型测试
 */

func TestAny(t *testing.T)  {
	a := func(a any) {
		switch a.(type) {
		case int:
			fmt.Println("int ")
		case float64:
			fmt.Println("float64")
		case string:
			fmt.Println("string")
		}
	}
	a(1)
	a(1.111)
	a("aaa")
}