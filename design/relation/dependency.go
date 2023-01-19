package relation

import "fmt"

/**
*依赖关系，最弱的关系
* Student play game 需要computer
 */

type Computer struct {
}
func (p *Computer) Start()  {
	fmt.Println("Starting")
}

type Student struct {}
func (p *Student) PlayGame() {
	computer := Computer{}
	computer.Start()
	fmt.Println("Play Game")
}