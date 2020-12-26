/**
 * Description:数据结构，队列
 * version 1.0.0
 * Created by GoLand.
 * Company sdbean
 * Author: hammercui
 * Date: 2020/12/26
 * Time: 17:16
 * Mail: hammercui@163.com
 *
 */
package queue

import "github.com/hammer/algorithmDesign/list"

// Queue 队列信息
type Queue struct {
	list *list.SingleList
}

// 新建queue实例
func NewQueue() *Queue {
	return &Queue{
		list: list.NewSingleList(),
	}
}

// Size 获取队列长度
func (q *Queue) Size() uint {
	return q.list.Size
}

// Enqueue 进入队列
func (q *Queue) Enqueue(data interface{}) bool {
	return q.list.Append(data)
}

// Dequeue 出列
func (q *Queue) Dequeue() interface{} {
	node := q.list.Get(0)
	if node == nil {
		return nil
	}
	q.list.Delete(0)
	return node.Data
}

// Peek 查看队头信息
func (q *Queue) Peek() interface{} {
	node := q.list.Get(0)
	if node == nil {
		return nil
	}
	return node.Data
}
