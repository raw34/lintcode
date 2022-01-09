package _0801

func backPackX(n int) int {
    max := func(a, b int) int {
        if a < b {
            return b
        }
        return a
    }
    prices := []int{150, 250, 350}
    m := len(prices)
    dp := make([][]int, m+1)
    for i := 0; i < m+1; i++ {
        dp[i] = make([]int, n+1)
    }
    for i := 1; i < m+1; i++ {
        p := prices[i-1]
        for j := 0; j < n+1; j++ {
            dp[i][j] = dp[i-1][j]
            for k := 0; p*k <= j; k++ {
                dp[i][j] = max(dp[i][j], dp[i][j-p*k]+p*k)
            }
        }
    }

    return n - dp[m][n]
}

func backPackX2(n int) int {
    max := func(a, b int) int {
        if a < b {
            return b
        }
        return a
    }
    prices := []int{150, 250, 350}
    m := len(prices)
    dp := make([]int, n+1)
    for i := 1; i < m+1; i++ {
        p := prices[i-1]
        for j := p; j < n+1; j++ {
            dp[j] = max(dp[j], dp[j-p]+p)
        }
    }

    return n - dp[n]
}
