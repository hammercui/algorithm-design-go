package dp

import "fmt"

//longest increase sort length
func Lis(a []int) int {
	var n = len(a)
	var f = make([]int,n)
	fmt.Println(a)
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
	return ans
}

func Max(a,b int) int  {
	if a> b{
		return a
	}else{
		return b
	}
}