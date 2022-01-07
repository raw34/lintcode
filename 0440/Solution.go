package _0440

/*
A = [2, 3, 5, 7]
V = [1, 5, 2, 4]
m = 10
dp = [
    [0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0],
    [0, 0, 1, 1, 2, 2, 3, 3, 4, 4, 5],
    [0, 0, 1, 5, 5, 6, 10, 10, 11, 15, 15],
    [0, 0, 1, 5, 5, 6, 10, 10, 11, 15, 15],
    [0, 0, 1, 5, 5, 6, 10, 10, 11, 15, 15]
]
*/
func backPackIII(A []int, V []int, m int) int {
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
        a, v := A[i-1], V[i-1]
        for j := a; j < m+1; j++ {
            for k := 0; a*k <= j; k++ {
                dp[i][j] = max(dp[i][j], dp[i-1][j-a*k]+v*k)
            }
        }
    }

    return dp[n][m]
}

/*
A = [2, 3, 5, 7]
V = [1, 5, 2, 4]
m = 10
dp = [0, 0, 1, 5, 5, 6, 10, 10, 11, 15, 15]
*/
func backPackIII2(A []int, V []int, m int) int {
    max := func(a, b int) int {
        if a < b {
            return b
        }
        return a
    }
    n := len(A)
    dp := make([]int, m+1)
    for i := 1; i < n+1; i++ {
        a, v := A[i-1], V[i-1]
        for j := a; j < m+1; j++ {
            dp[j] = max(dp[j], dp[j-a]+v)
        }
    }

    return dp[m]
}
