package _0092

/*
m = 10, A = [3, 4, 8, 5]
dp = [
    [0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0],
    [0, 0, 0, 3, 3, 3, 3, 3, 3, 3, 3],
    [0, 0, 0, 3, 4, 4, 4, 7, 7, 7, 7],
    [0, 0, 0, 3, 4, 4, 4, 7, 8, 8, 8],
    [0, 0, 0, 3, 4, 5, 5, 7, 8, 9, 9]
]
*/
func backPack(m int, A []int) int {
    max := func(a, b int) int {
        if a < b {
            return b
        }
        return a
    }
    n := len(A)
    // dp[i][j]表示取前i个元素且背包容量为j时能达到的最大值
    dp := make([][]int, n+1)
    for i := 0; i < n+1; i++ {
        dp[i] = make([]int, m+1)
    }
    for i := 1; i < n+1; i++ {
        a := A[i-1]
        for j := 0; j < m+1; j++ {
            if j < a {
                dp[i][j] = dp[i-1][j]
            } else {
                dp[i][j] = max(dp[i-1][j], dp[i-1][j-a]+a)
            }
        }
    }

    return dp[n][m]
}

/*
m = 10, A = [3, 4, 8, 5]
dp = [0 0 0 3 4 5 5 7 8 9 9]
*/
func backPack2(m int, A []int) int {
    max := func(a, b int) int {
        if a < b {
            return b
        }
        return a
    }
    n := len(A)
    // dp[j]表示取前i个元素且背包容量为j时能达到的最大值
    dp := make([]int, m+1)
    for i := 1; i < n+1; i++ {
        a := A[i-1]
        for j := m; j >= a; j-- {
            dp[j] = max(dp[j], dp[j-a]+a)
        }
    }

    return dp[m]
}
