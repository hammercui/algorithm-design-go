package creational

import (
	"bytes"
	"fmt"
)

/**
原型模式
1 定义接口 包含Clone方法
2 所有的实例都有实现这个接口
 */
type Node interface {
	String() string
	//
	Parent() Node
	//
	SetParent(node Node)
	//
	Children() []Node
	//
	AddChild(node Node)
	//
	Clone() Node
}

type Element struct {
	text string
	parent Node
	children []Node
}

func NewElement(text string)*Element  {
	return &Element{
		text:    text,
		parent:   nil,
		children: make([]Node, 0),
	}
}
func (p *Element) String() string {
	buffer := bytes.NewBufferString(p.text)
	for _,c := range p.children {
		text := c.String()
		fmt.Fprintf(buffer,"\n %s",text)
	}
	return buffer.String()
}

func (p *Element) Parent() Node {
	return p.parent
}

func (p *Element) SetParent(node Node) {
	p.parent = node
}

func (p *Element) Children() []Node {
	return p.children
}

func (p *Element) AddChild(node Node) {
	_copy := node.Clone()
	_copy.SetParent(p)
	p.children = append(p.children,_copy)
}

func (p *Element) Clone() Node {
	_copy := &Element{
		text:     p.text,
		parent:   nil,
		children: make([]Node,0),
	}
	for _,child := range p.children{
		_copy.AddChild(child)
	}
	return _copy
}



