package behavioral
/**
迭代器模式，常见于兼容的集合，但不同的遍历方式
包含两个要素
1.Collection：一种Collection返回自己的iterator。包含算法与数据结构
2.Iterator：实现遍历
 */
type Node struct {
	name string
	data interface{}
}

type Iterator interface {
	hasNext() bool
	getNext() *Node
}

type Collection interface {
	createIterator() Iterator
}

type UserCollection struct {
	users []*Node
}

func (p *UserCollection) createIterator() Iterator {
	return &UserIterator{users: p.users}
}

type UserIterator struct {
	index int
	users []*Node
}

func (u *UserIterator) hasNext() bool {
	if u.index < len(u.users) {
		return true
	}
	return false
}

func (u *UserIterator) getNext() *Node {
	if u.hasNext(){
		user := u.users[u.index]
		u.index++
		return user
	}
	return nil
}
