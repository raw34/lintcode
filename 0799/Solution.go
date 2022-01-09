package _0799

/*
n = 5
value = [1,4]
amount = [2,1]
dp =[
    [true, false, false, false, false, false],
    [true, true, true, false, false, false]
    [true, true, true, false, true, true]
]
*/
func backPackVIII(n int, value []int, amount []int) int {
    m := len(value)
    dp := make([][]bool, m+1)
    for i := 0; i < m+1; i++ {
        dp[i] = make([]bool, n+1)
    }
    dp[0][0] = true
    for i := 1; i < m+1; i++ {
        v, a := value[i-1], amount[i-1]
        for j := 0; j < n+1; j++ {
            for k := 0; k <= a; k++ {
                if j >= v*k && !dp[i][j] && dp[i-1][j-v*k] {
                    dp[i][j] = true
                }
            }
        }
    }

    res := 0
    for j := 1; j < n+1; j++ {
        if dp[m][j] {
            res++
        }
    }

    return res
}

/*
n = 5
value = [1,4]
amount = [2,1]
dp = [true, true, true, false, true, true]
*/
func backPackVIII2(n int, value []int, amount []int) int {
    m := len(value)
    dp := make([]bool, n+1)
    dp[0] = true
    for i := 1; i < m+1; i++ {
        v, a := value[i-1], amount[i-1]
        for k := 1; k <= a; k++ {
            for j := n; j >= v; j-- {
                if !dp[j] && dp[j-v] {
                    dp[j] = true
                }
            }
        }
    }

    res := 0
    for j := 1; j < n+1; j++ {
        if dp[j] {
            res++
        }
    }

    return res
}
