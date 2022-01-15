package dp

import (
	"fmt"
	"github.com/hammer/algorithmDesign/list"
	"sync"
)

//编辑举例算法

type DisActionEnum int

const(
	DIS_ACTION_ADD DisActionEnum = iota
	DIS_ACTION_DEL
	DIS_ACTION_REP
)


var DIS_ACTION_NAME =  map[DisActionEnum]string{
	DIS_ACTION_ADD:"add",
	DIS_ACTION_DEL:"del",
	DIS_ACTION_REP:"replace",
}

type DisTrail struct {
	action DisActionEnum
	row int
	column int
}

func minDistance(a string,b string) (int,map[int]*DisTrail) {
	//行数，代表列长度
	len1 := len(a)
	//列数，打榜行长度
	len2 := len(b)
	//新增空字符串的位置
	dp := make([][]int,len1+1)
	for a := range dp{
		dp[a] = make([]int,len2+1)
	}
	//trailMap := make(map[string]bool)

	//初始化第一列
	for row := 0; row < len1+1; row++ {
		dp[row][0] = row
	}
	//初始化第一行
	for column := 0; column < len2+1; column++ {
		dp[0][column] = column
	}

	var disTrailMap = make(map[int]*DisTrail)

	//循环两次,逐行扫描
	for row := 1; row < len1+1; row++ {
		for column := 1; column < len2+1; column++ {
			var newDisTrail DisTrail
			//转移方程
			//当左右相等，直接取左上角值
			if a[row - 1] == b[column - 1] {
				//当值相同时，复用上一步的操作，意味着{m+1,n+1}的DisTrail等同于{m,n}的DisTrail
				dp[row][column] = dp[row - 1][column - 1]
			}else{
				//取左，上，左上最小值 + 1
				var temp int
				//左小于上
				if dp[row][column - 1] < dp[row-1][column]{
					temp = dp[row][column - 1]
					newDisTrail.action = DIS_ACTION_ADD
					newDisTrail.row = row
					newDisTrail.column = column -1
				}else{
					temp = dp[row-1][column]
					newDisTrail.action = DIS_ACTION_DEL
					newDisTrail.row = row-1
					newDisTrail.column = column
				}
				//左上小于temp
				if dp[row-1][column - 1] < temp{
					temp = dp[row-1][column-1]
					newDisTrail.action = DIS_ACTION_REP
					newDisTrail.row = row-1
					newDisTrail.column = column-1
				}
				dp[row][column] = temp + 1
				disTrailMap[temp+1] = &newDisTrail
				fmt.Println("step",temp+1,newDisTrail)
			}
		}
	}

	//返回表格最右下角，就是最优解
	return dp[len1][len2],disTrailMap
}

func trailsToStep(a,b string,ans int,disTrailMap map[int]*DisTrail)  {
	//打印每一步
	var newA  = a
	for i := 1; i <= ans ; i++ {
		step := disTrailMap[i]
		fmt.Println(DIS_ACTION_NAME[step.action],step)
		//增加
		if step.action == DIS_ACTION_ADD{
			newA = newA[:step.row] + string(b[step.column]) + newA[step.row:]
			fmt.Println("a 执行add操作后",newA)
		}else if step.action == DIS_ACTION_REP{
			//更换
			newA = newA[:step.row] + string(b[step.column]) + newA[step.row+1:]
			fmt.Println("a 执行replace操作后",newA)
		} else if step.action == DIS_ACTION_DEL{
			//删除
			end := step.row+1
			if end>len(newA){
				newA = newA[:step.row]
			}else{
				newA = newA[:step.row] + newA[end:]
			}

			fmt.Println("a 执行del操作后",newA)
		}
	}
	//return stepActions
}

//新增数据结构StringLinkList
//支持
//2 更新指定index的数据
//3 删除指定index的数据

type StringLinkList struct {
	mutex *sync.RWMutex
	Head  *list.SingleNode
	Tail  *list.SingleNode
	Size  uint
}

func NewStringLinListr(from string) *StringLinkList  {
	ins := StringLinkList{
		mutex: new(sync.RWMutex),
	}
	//初始化
	if len(from) < 1{
		panic("from len < 1")
	}
	head := &list.SingleNode{
		Data: string(from[0]),
		Next: nil,
	}
	tail :=head
	for i := 0; i < len(from); i++ {
		ins.Size++
		if i < len(from)-1{
			tail.Next = &list.SingleNode{
				Data: string(from[i+1]),
				Next: nil,
			}
			tail = tail.Next
		}
	}
	ins.Head = head
	ins.Tail = tail

	return &ins
}

//1 指定index插入数据，如果当前index存在，自动往后寻找直到空白位置
func (p *StringLinkList) Insert(index int,val string) bool  {
	if val == "" {
		return false
	}
	//
	p.mutex.Lock()
	defer p.mutex.Unlock()
	return false
}
