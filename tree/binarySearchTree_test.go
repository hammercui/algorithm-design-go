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
)

func TestBSTree(t *testing.T) {
	tree := NewBSTree()
	tree.Add(&NodeValue{
		Key:   1,
		Value: 1,
	})
	for i := 0; i < 10; i++ {
		randomId := int(rand.Int31n(int32(100)))
		t.Log("randomId:",randomId)
		tree.Add(&NodeValue{
			Key:   randomId,
			Value: randomId,
		})
	}
	t.Logf("tree:%s", toJsonStr(tree.Root))
	values := PreOrder(tree)
	t.Log("values", values)
}

func toJsonStr(v interface{}) string {
	bytes, _ := json.Marshal(v)
	return string(bytes)
}
