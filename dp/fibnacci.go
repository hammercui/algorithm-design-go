package dp

//获得斐波那契数,两种方案，一种使用数据，空间复杂度o（n）,另一种o（3）

//v1 空间复杂度o（n）
func getFibnacciV1(n int) int  {
	var dp = make([]int,n)
	dp[0] =1
	dp[1] = 1
	for i := 2; i < n; i++ {
		//状态转移方程
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n-1]
}

//o（3）
func getFibnacciV2(n int) int  {
	//解题思路是每次只需要n-1,n-2的数局，n-3及以前并不需要保存
}