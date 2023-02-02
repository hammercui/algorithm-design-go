package main

import (
	"fmt"
)

/**
结构型设计模式
轻量模式：目的就是节省内存，前提是要大量创造对象，并且对象拥有部分相同的属性可拆分出来。
以下以游戏渲染树木举例
 */

//TreeFlyweight
type TreeFlyweight struct {
	name string
	color string
	texture string
	treeType int
}
func (p *TreeFlyweight) draw(x,y int)  {
	fmt.Printf("tree type: %d,color: %s,texture: %s,draw position: %d,%d \n", p.treeType,p.color,p.texture,x,y)
}

type TreeFactory struct {
	pools map[int]*TreeFlyweight
}
func (f *TreeFactory) get(treeType int,name string,color string,texture string) *TreeFlyweight {
	if v, ok := f.pools[treeType]; ok {
		return v
	}else{
		v := &TreeFlyweight{name: name,color: color,texture: texture,treeType: treeType}
		f.pools[treeType] = v
		return v
	}
}

//单例模式工厂
var TreeFactoryIns =  &TreeFactory{pools: make(map[int]*TreeFlyweight)}

type TreeContext struct {
	flyweight *TreeFlyweight
	x int
	y int
}
func (p *TreeContext) draw()  {
	p.flyweight.draw(p.x,p.y)
}
func NewTreeContext(treeType int,name string,color string,texture string,x,y int) *TreeContext {
	fly := TreeFactoryIns.get(treeType,name,color,texture)
	context := &TreeContext{
		flyweight: fly,
		x:         x,
		y:         y,
	}
	return context
}
