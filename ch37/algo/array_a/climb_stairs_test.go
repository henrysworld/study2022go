package array_a

import (
	"fmt"
	"testing"
)

func TestClimbStairs(t *testing.T) {
	count := climbStairs(4)
	fmt.Println(count)
}

func climbStairs(n int) int {
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1

	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}
