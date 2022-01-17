package dp

//从start点到end点，每次只能向下，向右，一共有多少路径

func uniquePaths(m,n int)  int{
	if m< 2 || n < 2{
		panic("m,n must >=2")
	}

	//1 构造dp二维数组
	var dp = make([][]int,m)
	for a, _ := range dp {
		dp[a] = make([]int,n)
	}

	//2 (i,j)的位置从左和上而来，则f(i,j) = f(i-1,j)+f(i,j-1)

	//3 最左一列都是1
	for i := 0; i < m; i++ {
		dp[i][0] = 1
	}
	//3.2最上一列全是1
	for j := 0; j < n; j++ {
		dp[0][j] = 1
	}
	//4 推导
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[m-1][n-1]
}