package relation

import "fmt"

/**
学生可以不用电脑，但学生不能没有老师
 */

type Teacher struct {
}

type StudentA struct {
	teacher Teacher
}
func (p *StudentA) Study() {
	fmt.Println("Study")
}