package dp

import (
	"testing"
)

func Test_lis(t *testing.T)  {
	//ax>ap，肯定能构造一个以ax结尾的子序列，长度为f(p)+1
	var a = []int{1,5,3,4,6,9,7,8}

	t.Log("ans",Lis(a))
}

