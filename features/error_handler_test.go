package main

import (
	"errors"
	"fmt"
	"os"
	"testing"
)

type CustomeError struct {
	msg string
	err error
}

func (c CustomeError) Wrap(err error)  {
	c.err = err
}

func (c CustomeError) Error() string {
	return c.msg
}
func (c CustomeError) Unwrap() error{
	return c.err
}



var (
	//ERROR_NON_EXISTING = errors.New("non-existing")
	ERROR_NON_EXISTING = &CustomeError{msg: "non-existing"}
)
/**
golang的错误处理
 */
func TestErrHandler(t *testing.T) {
	//1 test as
	if _, err := os.Open("non-existing"); err != nil {
		var pathError *os.PathError
		if errors.As(err, &pathError) {
			fmt.Println("Failed at path:", pathError.Path)
		} else {
			fmt.Println(err)
		}
	}

	e := errors.New("first error")
	//嵌套error
	w := fmt.Errorf("Wrap new err %w",e)
	fmt.Println(w)
	//解嵌套
	fmt.Println(errors.Unwrap(w))
	//is判断是这个错误，或者是否嵌套错误
	if errors.Is(errors.Unwrap(w),e) {
		fmt.Println("Unwrap success!")
	}
	//as判断为是否可转换为这个错误，是包含赋值的 因此必须是指针
	second := fmt.Errorf("Wrap second err %w", ERROR_NON_EXISTING)
	if errors.Is(errors.Unwrap(second), ERROR_NON_EXISTING) {
		fmt.Println("Unwrap custom error success!")
	}else{
		fmt.Println("Unwrap custom error fail!")
	}
	//wrap 多个error等golang1.20

}