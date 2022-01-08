package _0798

/*
n = 8
prices = [3,2]
weights = [300,160]
amounts = [1,6]
dp = [
    [0, 0, 0, 0, 0, 0, 0, 0, 0]
    [0, 0, 0, 300, 300, 300, 300, 300, 300]
    [0, 0, 160, 300, 320, 460, 480, 620, 640]
]
*/
func backPackVII(n int, prices []int, weight []int, amounts []int) int {
    max := func(a, b int) int {
        if a < b {
            return b
        }
        return a
    }
    m := len(prices)
    dp := make([][]int, m+1)
    for i := 0; i < m+1; i++ {
        dp[i] = make([]int, n+1)
    }
    for i := 1; i < m+1; i++ {
        p, w, a := prices[i-1], weight[i-1], amounts[i-1]
        for j := 0; j < n+1; j++ {
            dp[i][j] = dp[i-1][j]
            for k := 0; k <= a; k++ {
                if j >= p*k {
                    dp[i][j] = max(dp[i][j], dp[i-1][j-p*k]+w*k)
                }
            }
        }
    }

    return dp[m][n]
}

/*
n = 8
prices = [3,2]
weights = [300,160]
amounts = [1,6]
dp = [0, 0, 160, 300, 320, 460, 480, 620, 640]
*/
func backPackVII2(n int, prices []int, weight []int, amounts []int) int {
    max := func(a, b int) int {
        if a < b {
            return b
        }
        return a
    }
    m := len(prices)
    dp := make([]int, n+1)
    for i := 1; i < m+1; i++ {
        p, w, a := prices[i-1], weight[i-1], amounts[i-1]
        for j := n; j >= p; j-- {
            for k := 0; k <= a; k++ {
                if j >= p*k {
                    dp[j] = max(dp[j], dp[j-p*k]+w*k)
                }
            }
        }
    }

    return dp[n]
}
