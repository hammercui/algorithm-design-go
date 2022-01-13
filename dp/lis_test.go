package dp

import (
	"fmt"
	"testing"
)

func Test_lis(t *testing.T)  {
	//ax>ap，肯定能构造一个以ax结尾的子序列，长度为f(p)+1
	var a = []int{1,5,3,4,6,9,7,8}
	var n = len(a)
	var f = make([]int,n)
	t.Log(a)
	f[0]=1
	for x := 0; x < n ; x++ {
		for p := 0; p < x; p++ {
			if a[p]<a[x]{
				f[x]= Max(f[x],f[p]+1)
			}
		}
		fmt.Printf("f[%d]=%d \n",x,f[x])
	}

	//获得最长值
	var ans = f[0]
	for i := 0; i < n; i++ {
		ans = Max(ans,f[i])
	}
	t.Log("ans",ans)
}

func Max(a,b int) int  {
	if a> b{
		return a
	}else{
		return b
	}
}