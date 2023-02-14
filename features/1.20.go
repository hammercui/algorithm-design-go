package features

import "fmt"

/**
切片转换为数组，数组长度必须小于切片，而且转换后数组为切片的副本，修改互不影响
 */
func sliceToArray()  {
	 s1 := []int{1,2,3,4,5,6,7}
	 arr := [7]int(s1)
	 fmt.Println(s1)
	 fmt.Println(arr)

	 s1[0] = 100
	 arr[1] = 200

	fmt.Println(s1)
	fmt.Println(arr)
}

//自定义error类
type MyError struct {
	s string
}
func (p *MyError) Error() string {
	return  p.s
}



