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

	tree.SetRoot(&TreeNode{Data: &TreeNodeValue{Key: 1,Value: 1}})
	if tree.GetRoot().Data == nil {
		t.Errorf("empty bsTree set fail ")
	} else {
		t.Errorf("empty bsTree set success")
	}
}



func toJsonStr(v interface{}) string {
	bytes, _ := json.Marshal(v)
	return string(bytes)
}
