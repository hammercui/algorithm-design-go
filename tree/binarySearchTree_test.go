/**
 * Description
 * version 1.0.0
 * Created by GoLand.
 * Company sdbean
 * Author: hammercui
 * Date: 2020/12/25
 * Time: 13:48
 * Mail: hammercui@163.com
 *
 */
package tree

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestBSTree(t *testing.T) {
	tree := NewBSTree()
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		randomId := int(rand.Int31n(int32(100)))
		t.Log("randomId:", randomId)
		tree.Add(&TreeNodeValue{
			Key:   randomId,
			Value: randomId,
		})
	}
	t.Logf("tree:%s", toJsonStr(tree.GetRoot()))
}

func TestBSTree_SetRoot(t *testing.T) {
	tree := NewBSTree()
	if tree.GetRoot().Data == nil {
		t.Log("empty bsTree get success")
	} else {
		t.Errorf("empty bsTree get fail")
	}

	tree.SetRoot(&TreeNode{Data: &TreeNodeValue{Key: 1, Value: 1}})
	if tree.GetRoot().Data == nil {
		t.Errorf("empty bsTree set fail ")
	} else {
		t.Errorf("empty bsTree set success")
	}
}

func TestBSTree_Add(t *testing.T) {
	tree := NewBSTree()
	if tree.Add(&TreeNodeValue{Key: 1}) {
		t.Log("bstree add success")
	} else {
		t.Fatalf("bstree add fail")
	}

	if tree.Add(nil) {
		t.Fatalf("bstree add nil fail")
	} else {
		t.Log("bstree add nil success")
	}

}

func TestBSTree_Del(t *testing.T) {
	tree := NewBSTree()
	tree.Add(&TreeNodeValue{
		Key: 100,
	})
	tree.Add(&TreeNodeValue{
		Key: 50,
	})
	tree.Add(&TreeNodeValue{
		Key: 60,
	})
	tree.Add(&TreeNodeValue{
		Key: 40,
	})
	tree.Add(&TreeNodeValue{
		Key: 35,
	})
	tree.Add(&TreeNodeValue{
		Key: 45,
	})
	tree.Add(&TreeNodeValue{
		Key: 11,
	})
	tree.Add(&TreeNodeValue{
		Key: 21,
	})
	tree.Add(&TreeNodeValue{
		Key: 32,
	})

	tree.Add(&TreeNodeValue{
		Key: 7,
	})
	if tree.Del(50) {
		t.Log("bstree delete exist success")
	} else {
		t.Fatalf("bstree delete exist fail")
	}

	if tree.Del(5) {
		t.Fatalf("bstree delete not exist fail")
	} else {
		t.Log("bstree delete not exist  success")
	}

}

func TestBSTree_bfsPrint(t *testing.T) {
	tree := NewBSTree()
	tree.Add(&TreeNodeValue{
		Key: 100,
	})
	tree.Add(&TreeNodeValue{
		Key: 50,
	})
	tree.Add(&TreeNodeValue{
		Key: 60,
	})
	tree.Add(&TreeNodeValue{
		Key: 40,
	})
	tree.Add(&TreeNodeValue{
		Key: 35,
	})
	tree.Add(&TreeNodeValue{
		Key: 45,
	})
	tree.Add(&TreeNodeValue{
		Key: 11,
	})
	tree.Add(&TreeNodeValue{
		Key: 21,
	})
	tree.Add(&TreeNodeValue{
		Key: 32,
	})

	tree.Add(&TreeNodeValue{
		Key: 7,
	})

	PrintBFS(tree)

	//
	dekKey := 50
	fmt.Println("删除key:", dekKey)
	tree.Del(dekKey)
	PrintBFS(tree)
}

func toJsonStr(v interface{}) string {
	bytes, _ := json.Marshal(v)
	return string(bytes)
}
