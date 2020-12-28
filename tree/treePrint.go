/**
 * Description:打印数
 * version 1.0.0
 * Created by GoLand.
 * Company sdbean
 * Author: hammercui
 * Date: 2020/12/26
 * Time: 10:59
 * Mail: hammercui@163.com
 *
 */
package tree

import (
	"fmt"
	"log"
)

//深度优先遍历打印
/*
深度优先遍历，从左子树触发，一直找到左子节点
*/
func PrintDFS(tree BinaryTree) {
}

/*
广度优先遍历，先遍历本层节点，遍历本层的基础上再纵向遍历,
借助于queue结构进行输出
*/

func PrintBFS(tree BinaryTree) {
	//根节点
	currNode := tree.GetRoot()
	if currNode == nil || currNode.Data == nil {
		return
	}

	var arr []*TreeNode
	arr = append(arr, currNode)
	for len(arr) > 0 {
		//
		root := arr[0]
		log.Println("根节点", root.Data.Key)
		printChild := ""
		//打印当前点
		//左，入队
		if root.Left != nil {
			arr = append(arr, root.Left)
			printChild = fmt.Sprintf("[%d", root.Left.Data.Key)
		}

		//右，入队
		if root.Right != nil {
			arr = append(arr, root.Right)
			printChild = fmt.Sprintf("%s,%d]", printChild, root.Right.Data.Key)
		}
		log.Println("子节点", printChild)
		//队首出队
		arr = arr[1:]
	}
}
