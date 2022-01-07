package _0563

/*
nums = [1,2,3,3,7]
target = 7
dp = [
    [1, 0, 0, 0, 0, 0, 0, 0],
    [1, 1, 0, 0, 0, 0, 0, 0],
    [1, 1, 1, 1, 0, 0, 0, 0],
    [1, 1, 1, 2, 1, 1, 1, 0],
    [1, 1, 1, 3, 2, 2, 3, 1],
    [1, 1, 1, 3, 2, 2, 3, 2],
]
*/
func backPackV(nums []int, target int) int {
    n := len(nums)
    dp := make([][]int, n+1)
    for i := 0; i < n+1; i++ {
        dp[i] = make([]int, target+1)
    }
    dp[0][0] = 1
    for i := 1; i < n+1; i++ {
        w := nums[i-1]
        for j := 0; j < target+1; j++ {
            if j < w {
                dp[i][j] = dp[i-1][j]
            } else {
                dp[i][j] = dp[i-1][j] + dp[i-1][j-w]
            }
        }
    }

    return dp[n][target]
}

/*
nums = [1,2,3,3,7]
target = 7
dp = [1, 1, 1, 3, 2, 2, 3, 2]
*/
func backPackV2(nums []int, target int) int {
    n := len(nums)
    dp := make([]int, target+1)
    dp[0] = 1
    for i := 1; i < n+1; i++ {
        w := nums[i-1]
        for j := target; j >= w; j-- {
            dp[j] += dp[j-w]
        }
    }

    return dp[target]
}
