/**
 * Description:单向链表
 * version 1.0.0
 * Created by GoLand.
 * Company sdbean
 * Author: hammercui
 * Date: 2020/12/26
 * Time: 14:41
 * Mail: hammercui@163.com
 *
 */
package list

import (
	"fmt"
	"sync"
)

//单链表节点
type SingleNode struct {
	Data interface{}
	Next *SingleNode
}

type SingleList struct {
	mutex *sync.RWMutex
	Head  *SingleNode
	Tail  *SingleNode
	Size  uint
}

//获得单链表实例
func NewSingleList() *SingleList {
	return &SingleList{
		mutex: new(sync.RWMutex),
	}
}

//添加节点到链表尾部
func (p *SingleList) Append(date interface{}) bool {
	if date == nil {
		return false
	}
	node := &SingleNode{Data: date}
	p.mutex.Lock()
	defer p.mutex.Unlock()
	if p.Size == 0 {
		p.Size = 1
		p.Head = node
		p.Tail = node
		return true
	}
	//插入尾部
	oldTail := p.Tail
	oldTail.Next = node
	p.Tail = node
	p.Size++
	return true
}

//插入节点到指定位置
func (p *SingleList) Insert(index uint, date interface{}) bool {
	if date == nil {
		return false
	}
	if index > p.Size {
		return false
	}
	node := &SingleNode{Data: date}

	p.mutex.Lock()
	defer p.mutex.Unlock()
	//当前链表为空时
	if p.Size == 0 {
		p.Size = 1
		p.Head = node
		p.Tail = node
		return true
	}
	//遍历查找指定位置
	var i uint
	ptr := p.Head
	for i = uint(1); i < index; i++ {
		ptr = ptr.Next
	}
	oldSelNode := ptr.Next
	ptr.Next = node
	node.Next = oldSelNode
	p.Size++
	return true
}

//删除指定位置的节点
func (p *SingleList) Delete(index uint) bool {
	if p == nil || p.Size == 0 || index > p.Size {
		return false
	}

	p.mutex.Lock()
	defer p.mutex.Unlock()

	//删除头
	if index == 0 {
		newHead := p.Head.Next
		p.Head = newHead
		if p.Size == 1 {
			p.Tail = nil
		}
		p.Size--
		return true
	}

	//删除其他位置
	ptr := p.Head
	var i uint
	for i = 1; i < index; i++ {
		ptr = ptr.Next
	}
	next := ptr.Next
	ptr.Next = next
	if index == p.Size-1 {
		p.Tail = ptr
	}
	p.Size--
	return true
}

//查找指定位置的节点
func (p *SingleList) Get(index uint) *SingleNode {
	if p == nil || p.Size == 0 || index > p.Size-1 {
		return nil
	}

	p.mutex.RLock()
	defer p.mutex.RUnlock()
	ptr := p.Head
	for i := uint(0); i < index; i++ {
		ptr = ptr.Next
	}
	return ptr
}

// 输出链表
func (p *SingleList) Display() {
	if p == nil || p.Size == 0 {
		fmt.Println("this single list is nil")
		return
	}
	p.mutex.RLock()
	defer p.mutex.RUnlock()
	fmt.Printf("this single list size is %d \n", p.Size)
	ptr := p.Head
	var i uint
	for i = 0; i < p.Size; i++ {
		fmt.Printf("No%3d data is %v\n", i+1, ptr.Data)
		ptr = ptr.Next
	}
}
