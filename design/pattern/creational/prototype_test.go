package creational

import (
	"fmt"
	"testing"
)

func TestPrototype(t *testing.T) {
	//总监
	directorNode := NewElement("Director of Engineering")
	//研发经理
	engManagerNode := NewElement("Engineering Manager")
	engManagerNode.SetParent(directorNode)
	engManagerNode.AddChild(NewElement("Lead Software Engineer"))
	//办公室经理
	officeManagerNode := NewElement("Office Manager")
	officeManagerNode.SetParent(directorNode)

	//经理是总监的下级
	directorNode.AddChild(engManagerNode)
	directorNode.AddChild(officeManagerNode)

	//
	fmt.Println("# Company Hierarchy")
	fmt.Print(directorNode)
	fmt.Println(" ")
	//
	fmt.Println("# Team Hierarchy")
	fmt.Print(engManagerNode)
	fmt.Println(" ")
	// 从研发经理节点克隆出一颗新的树
	fmt.Println("# New Team Hierarchy")
	fmt.Print(engManagerNode.Clone())

}
