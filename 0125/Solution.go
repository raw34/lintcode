package _0125

/*
m = 10
A = [2, 3, 5, 7]
V = [1, 5, 2, 4]
dp = [
    [0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0],
    [0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1],
    [0, 0, 1, 5, 5, 6, 6, 6, 6, 6, 6],
    [0, 0, 1, 5, 5, 6, 6, 6, 7, 7, 8],
    [0, 0, 1, 5, 5, 6, 6, 6, 7, 7, 9]
]
*/
func backPackII(m int, A []int, V []int) int {
    max := func(a, b int) int {
        if a < b {
            return b
        }
        return a
    }
    n := len(A)
    // dp[i][j]表示取前i个元素在背包容量为j的情况下的最大值
    dp := make([][]int, n+1)
    for i := 0; i < n+1; i++ {
        dp[i] = make([]int, m+1)
    }
    for i := 1; i < n+1; i++ {
        a, v := A[i-1], V[i-1]
        for j := 0; j < m+1; j++ {
            if j < a {
                dp[i][j] = dp[i-1][j]
            } else {
                dp[i][j] = max(dp[i-1][j], dp[i-1][j-a]+v)
            }
        }
    }

    return dp[n][m]
}

/*
m = 10
A = [2, 3, 5, 7]
V = [1, 5, 2, 4]
dp = [0, 0, 1, 5, 5, 6, 6, 6, 7, 7, 9]
*/
func backPackII2(m int, A []int, V []int) int {
    max := func(a, b int) int {
        if a < b {
            return b
        }
        return a
    }
    n := len(A)
    // dp[j]表示取前i个元素在背包容量为j的情况下的最大值
    dp := make([]int, m+1)
    for i := 1; i < n+1; i++ {
        a, v := A[i-1], V[i-1]
        for j := m; j >= a; j-- {
            dp[j] = max(dp[j], dp[j-a]+v)
        }
    }

    return dp[m]
}
