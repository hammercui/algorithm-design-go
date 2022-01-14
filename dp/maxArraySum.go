package dp

func MaxArraySum(input []int)(int,[]int)  {
	n := len(input)
	f := make([]int,n)
	var output []int
	//状态初始化
	f[0] = input[0]
	ans := Max(f[0],0)
	for i := 1; i < n; i++ {
		//分离指标函数，列表状态转移方程
		if f[i-1] > 0 {
			f[i] = f[i-1] + input[i]
		}else{
			output = output[:0]
			f[i] = input[i]
		}

		output = append(output,input[i])

		//最优解是最优子结构当中的最大值
		ans = Max(f[i],ans)
	}
	return ans,output
}