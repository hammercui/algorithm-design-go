/**
 * Description:二叉树
 * version 1.0.0
 * Created by GoLand.
 * Company sdbean
 * Author: hammercui
 * Date: 2020/12/26
 * Time: 11:11
 * Mail: hammercui@163.com
 *
 */
package tree

type BinaryTree interface {
	GetRoot() *TreeNode
	SetRoot(node *TreeNode)
	Add(data *TreeNodeValue)
	Search(key int) (*TreeNodeValue, error)
}
