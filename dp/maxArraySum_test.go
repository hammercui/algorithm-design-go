package dp

import "testing"

func Test_maxArraySum(t *testing.T)  {
	input :=[]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	ans,out := MaxArraySum(input)
	t.Log("ans",ans)
	t.Log("out",out)
}