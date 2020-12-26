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

type TreeNodeValue struct {
	Key   int
	Value interface{}
}

//树节点
type TreeNode struct {
	Data  *TreeNodeValue //节点的值
	Left  *TreeNode
	Right *TreeNode
}

type BSTree struct {
	root *TreeNode
}

func NewBSTree() *BSTree {
	ins := &BSTree{
		root: &TreeNode{
			Data:  nil,
			Left:  nil,
			Right: nil,
		},
	}
	return ins
}

func (p *BSTree) GetRoot() *TreeNode {
	return p.root
}
func (p *BSTree) SetRoot(node *TreeNode) {
	p.root = node
}

//插入
func (p *BSTree) Add(data *TreeNodeValue) bool{
	if data == nil || data.Value == nil{
		return false
	}
	//若root节点是空，则将data所指節点作为根節点插入，否则：
	return p.add(p.root, data)
}

//插入实现
func (p *BSTree) add(root *TreeNode, data *TreeNodeValue) bool{
	//若root节点是空，则将data所指節点作为根節点插入，否则：
	if root.Data == nil {
		root.Data = data
		return true
	}
	//若data等于根節点的数据域之值，则返回，否则：
	if data.Key == root.Data.Key {
		return true
	}
	//若data小于根節点的数据域之值，则把data插入到左子树中，否则：
	if data.Key < root.Data.Key {
		if root.Left == nil {
			root.Left = &TreeNode{Data: nil}
		}
		return p.add(root.Left, data)
	} else {
		if root.Right == nil {
			root.Right = &TreeNode{Data: nil}
		}
		//把data插入到右子树中。
		return p.add(root.Right, data)
	}
}

//查询
func (p *BSTree) Search(key int) (*TreeNodeValue, error) {
	return p.search(p.root, key)
}

//查询实现
func (p *BSTree) search(root *TreeNode, key int) (*TreeNodeValue, error) {
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


