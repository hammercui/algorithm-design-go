/**
 * Description:队列
 * version 1.0.0
 * Created by GoLand.
 * Company sdbean
 * Author: hammercui
 * Date: 2020/12/26
 * Time: 17:18
 * Mail: hammercui@163.com
 *
 */
package queue

import "testing"

func TestQueue_Enqueue(t *testing.T) {
	q := NewQueue()

	q.Enqueue(1)
	if 1 == q.Size() {
		t.Log("queue size and enqueue success")
	} else {
		t.Error("queue size and enqueue failed")
	}
}

func TestQueue_Dequeue(t *testing.T) {
	q := NewQueue()

	d1 := q.Dequeue()
	if d1 == nil {
		t.Log("empty queue dequeue success")
	} else {
		t.Error("empty queue dequeue failed")
	}

	q.Enqueue(1)
	d := q.Dequeue()
	if 1 == d.(int) && q.Size() == 0 {
		t.Log("queue dequeue success")
	} else {
		t.Error("queue dequeue failed")
	}
}

func TestQueue_Peek(t *testing.T) {
	q := NewQueue()

	d1 := q.Peek()
	if d1 == nil {
		t.Log("empty queue peek success")
	} else {
		t.Error("empty queue peek failed")
	}

	q.Enqueue(1)
	d := q.Peek()
	if 1 == d.(int) && q.Size() == 1 {
		t.Log("queue peek success")
	} else {
		t.Error("queue peek failed")
	}
}
