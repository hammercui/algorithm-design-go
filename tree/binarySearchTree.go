/**
 * Description
 * version 1.0.0
 * Created by GoLand.
 * Company sdbean
 * Author: hammercui
 * Date: 2020/12/25
 * Time: 13:51
 * Mail: hammercui@163.com
 *
 */
package tree

import "errors"

/*
二叉搜索树
二叉树：每个节点至多有两个子节点
二叉查找树的特点：对于每一个节点X,左子树中所有项的值小于X中的值，而它的右子树种所有项的值大于X中的值。
*/

type NodeValue struct {
	Key   int
	Value interface{}
}

//树节点
type Node struct {
	Data  *NodeValue //节点的值
	Left  *Node
	Right *Node
}

type BSTree struct {
	Root *Node
}

func NewBSTree() *BSTree {
	ins := &BSTree{
		Root: &Node{
			Data:  nil,
			Left:  nil,
			Right: nil,
		},
	}
	return ins
}

//插入
func (p *BSTree) Add(data *NodeValue) {
	//若root节点是空，则将data所指節点作为根節点插入，否则：
	p.add(p.Root, data)
}

//插入实现
func (p *BSTree) add(root *Node, data *NodeValue) {
	//若root节点是空，则将data所指節点作为根節点插入，否则：
	if root.Data == nil {
		root.Data = data
		return
	}
	//若data等于根節点的数据域之值，则返回，否则：
	if data.Key == root.Data.Key {
		return
	}
	//若data小于根節点的数据域之值，则把data插入到左子树中，否则：
	if data.Key < root.Data.Key {
		if root.Left == nil {
			root.Left = &Node{Data: nil}
		}
		p.add(root.Left, data)
	} else {
		if root.Right == nil {
			root.Right = &Node{Data: nil}
		}
		//把data插入到右子树中。
		p.add(root.Right, data)
	}
}

//查询
func (p *BSTree) Search(key int) (*NodeValue, error) {
	return p.search(p.Root, key)
}

//查询实现
func (p *BSTree) search(root *Node, key int) (*NodeValue, error) {
	if root.Data == nil {
		return nil, errors.New("key not exist!")
	}
	if root.Data != nil && root.Data.Key == key {
		return root.Data, nil
	}
	//key小于节点值,继续左子树查询
	if key < root.Data.Key {
		return p.search(root.Left, key)
	} else {
		return p.search(root.Right, key)
	}
}

func PreOrder(tree *BSTree) []int {
	values := make([]int, 0, 50)
	//进行前序遍历，把结果输出，看看二叉查找树构造是否正确
	current := tree.Root

	if current == nil {
		return values
	}

	leftTree := new(BSTree)
	leftTree.Root = current.Left
	rightTree := new(BSTree)
	rightTree.Root = current.Right
	values = append(values, current.Data.Key)
	left := PreOrder(leftTree)
	right := PreOrder(rightTree)
	values = append(values, left...)
	values = append(values, right...)
	return values
}
