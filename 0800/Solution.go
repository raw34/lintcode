package _800

import (
    "fmt"
    "math"
)

/*
n = 10
prices = [4,4,5]
probability = [0.1,0.2,0.3]
dp = [
    [0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0],
    [0, 0, 0, 0, 0.1, 0.1, 0.1, 0.1, 0.1, 0.1, 0.1],
    [0, 0, 0, 0, 0.2, 0.2, 0.2, 0.2, 0.28, 0.28, 0.28],
    [0, 0, 0, 0, 0.2, 0.3, 0.3, 0.3, 0.3, 0.44, 0.44]
]
*/
func backpackIX(n int, prices []int, probability []float64) float64 {
    m := len(prices)
    dp := make([][]float64, m+1)
    for i := 0; i < m+1; i++ {
        dp[i] = make([]float64, n+1)
    }
    for i := 1; i < m+1; i++ {
        ps, py := prices[i-1], probability[i-1]
        for j := 0; j < n+1; j++ {
            dp[i][j] = dp[i-1][j]
            if j >= ps {
                dp[i][j] = math.Max(dp[i-1][j], dp[i-1][j-ps]+py-dp[i-1][j-ps]*py)
            }
        }
    }
    fmt.Println(dp)

    return dp[m][n]
}

/*
n = 10
prices = [4,4,5]
probability = [0.1,0.2,0.3]
dp =[0, 0, 0, 0, 0.2, 0.3, 0.3, 0.3, 0.3, 0.44, 0.44]
*/
func backpackIX2(n int, prices []int, probability []float64) float64 {
    m := len(prices)
    dp := make([]float64, n+1)
    for i := 1; i < m+1; i++ {
        ps, py := prices[i-1], probability[i-1]
        for j := n; j >= ps; j-- {
            dp[j] = math.Max(dp[j], dp[j-ps]+py-dp[j-ps]*py)
        }
    }

    return dp[n]
}
