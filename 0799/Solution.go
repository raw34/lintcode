package _0799

func backPackVIII(n int, value []int, amount []int) int {
    // TODO，调试
    m := len(value)
    dp := make([][]int, m+1)
    for i := 0; i < m+1; i++ {
        dp[i] = make([]int, n+1)
    }
    dp[0][0] = 1
    for i := 1; i < m+1; i++ {
        v, a := value[i-1], amount[i-1]
        for j := 0; j < n+1; j++ {
            for k := 1; k <= a; k++ {
                if j >= v*k {
                    dp[i][j] += dp[i-1][j-v*k] + 1
                }
            }
        }
    }

    return dp[m][n]
}
