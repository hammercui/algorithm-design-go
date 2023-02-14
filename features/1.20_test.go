package features

import (
	"errors"
	"fmt"
	"testing"
)

func TestSliceToArray(t *testing.T) {
	sliceToArray()
}

func TestWrapError(t *testing.T) {

	e1 := errors.New("error1")
	e2 := errors.New("error2")
	e3 := errors.New("error3")
	e4 := &MyError{
		s:"error4",
	}
	//多个error包装成一个error
	e := fmt.Errorf("%w,%w,%w,%w,",e1,e2,e3,e4)
	fmt.Printf("e = %s \n",e.Error())

	//判断错误是否能相互转化
	fmt.Printf("e == e1 ? %t \n",errors.Is(e,e1))

	//错误转换
	var ner *MyError
	fmt.Println(errors.As(e,&ner))
	fmt.Println(ner==e4)

	fmt.Printf("e = %s \n",e.Error())
}