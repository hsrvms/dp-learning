package main

/*
	Problem:
		Climbing Stairs (k steps)

		You are climbing a stair case. It takes n steps to reach the top.
		Each time you can either climb 1..k steps.
		In how many distinct ways you can climb to the top.

	Framework for Solving DP Problems:
		1. Define the objective function
			f(i) is the number of distinct ways to reach the i-th stair
		2. Identify base cases
			f(0) = 1
			f(1) = 1
		3. Write down a recurrence relation for the optimized objective function
			f(n) = f(n-1) + f(n-2) + ... + f(n-k)
		4. What's the order of execution
			bottom-up
		5. Where to look for the answer
			f(n)
*/

// Time complexity: O(nk)
// Space complexity: O(n)

func climbStairs(n, k int) int {
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1
	for i := 2; i <= n; i++ {
		for j := 1; j <= k; j++ {
			if i-j < 0 {
				break
			}
			dp[i] += dp[i-j]
		}
	}
	return dp[n]
}
