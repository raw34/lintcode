package _0092

/**
 * @param m: An integer m denotes the size of a backpack
 * @param A: Given n items with size A[i]
 * @return: The maximum size
 */
func backPack(m int, A []int) int {
    max := func(a, b int) int {
        if a < b {
            return b
        }
        return a
    }
    n := len(A)
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

func backPack2(m int, A []int) int {
    max := func(a, b int) int {
        if a < b {
            return b
        }
        return a
    }
    n := len(A)
    dp := make([]int, m+1)
    for i := 1; i < n+1; i++ {
        a := A[i-1]
        for j := m; j >= a; j-- {
            dp[j] = max(dp[j], dp[j-a]+a)
        }
    }

    return dp[m]
}
