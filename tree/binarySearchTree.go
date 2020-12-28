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

import (
	"errors"
)

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
func (p *BSTree) Add(data *TreeNodeValue) bool {
	if data == nil {
		return false
	}
	//若root节点是空，则将data所指節点作为根節点插入，否则：
	return p.add(p.root, data)
}

//插入实现
func (p *BSTree) add(root *TreeNode, data *TreeNodeValue) bool {
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

//删除
func (p *BSTree) Del(key int) bool {
	if p == nil || p.root.Data == nil {
		return false
	}
	return p.delBsf(nil, p.root, key)
}

//深度优先遍历，找到待删除
func (p *BSTree) delBsf(father *TreeNode, node *TreeNode, key int) bool {
	if node == nil || node.Data == nil {
		return false
	}
	//根节点
	if father == nil {
		if node == nil || node.Data == nil {
			return false
		} else if node.Data.Key == key {
			p.root = nil
			return true
		}
	}

	if key < node.Data.Key {
		return p.delBsf(node, node.Left, key)
	} else if key > node.Data.Key {
		return p.delBsf(node, node.Right, key)
	} else {
		return p.delImp(father, node)
	}
}

func (p *BSTree) delImp(father *TreeNode, node *TreeNode) bool {
	//1 node是叶子节点，直接删除
	if node.Left == nil && node.Right == nil {
		if father.Left.Data.Key == node.Data.Key {
			father.Left = nil
			return true
		}
		if father.Right.Data.Key == node.Data.Key {
			father.Right = nil
			return true
		}
	} else if node.Right == nil {
		//2 node只有左字树或者右子树，node的l或者r,直接成为双亲节点的l或者r
		//left不空 right空
		father.Left = node.Left
	} else if node.Left == nil {
		//left空，right不空
		father.Right = node.Right
	} else {
		//3 node左右都不为空。
		//3.1 node的左子节点nl有右子节点nlr，
		//		找到右最下的节点nlrmax,该节点为node前驱节点，即仅次于node节点大小的节点。
		//      然后nlrmax赋值给node,然后nlrmax此时已经没有右节点，把它剩余的左分支挂载与父节点的右分支。相当于删除了nlrmax。
		nl := node.Left
		if nl.Right != nil {
			nlrmaxFather := nl
			nlrmax := nl.Right
			for nlrmax.Right != nil {
				nlrmaxFather = nlrmax
				nlrmax = nlrmax.Right
			}
			node.Data = &TreeNodeValue{
				Key:   nlrmax.Data.Key,
				Value: nlrmax.Data.Value,
			}
			nlrmaxFather.Right = nlrmax.Left
		} else {
			//3.2 否
			//     nl提为node的父亲的左子树。node的右节点变为nl的右节点。相当于删除了node
			father.Left = nl
			nl.Right = node.Right
		}
	}

	return true
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
