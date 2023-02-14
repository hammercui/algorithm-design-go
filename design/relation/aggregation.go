package relation

import "fmt"

/**
是association的更进一步耦合，表示个体与集体，比如多个学生和老师组成了教室
有多维关联的意味，
 */

type Classes struct {
	student *StudentA
}
//student 构造时传入
//可以与classes拥有不同的生命周期
func test()  {
	student1 := StudentA{}
	classes1 := Classes{
		student: &student1,
	}
	fmt.Println(classes1)
}