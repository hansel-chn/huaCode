/**
You are given an integer array prices where prices[i] is the price of a given
stock on the iᵗʰ day, and an integer k.

 Find the maximum profit you can achieve. You may complete at most k
transactions: i.e. you may buy at most k times and sell at most k times.

 Note: You may not engage in multiple transactions simultaneously (i.e., you
must sell the stock before you buy again).


 Example 1:


Input: k = 2, prices = [2,4,1]
Output: 2
Explanation: Buy on day 1 (price = 2) and sell on day 2 (price = 4), profit = 4-
2 = 2.


 Example 2:


Input: k = 2, prices = [3,2,6,5,0,3]
Output: 7
Explanation: Buy on day 2 (price = 2) and sell on day 3 (price = 6), profit = 6-
2 = 4. Then buy on day 5 (price = 0) and sell on day 6 (price = 3), profit = 3-0
 = 3.



 Constraints:


 1 <= k <= 100
 1 <= prices.length <= 1000
 0 <= prices[i] <= 1000


 Related Topics 数组 动态规划 👍 1148 👎 0

*/

// leetcode submit region begin(Prohibit modification and deletion)
package main

//func maxProfit(k int, prices []int) int {
//	if k > len(prices)/2 {
//		k = k/2 + 1
//	}
//
//	dp := make([][][2]int, len(prices))
//
//	for i := range dp {
//		dp[i] = make([][2]int, k+1)
//	}
//
//	for i := 0; i < len(prices); i++ {
//		for j := 1; j < k+1; j++ {
//			if i < 1 {
//				dp[i][j][1] = -prices[i]
//				dp[i][j][0] = 0
//				continue
//			}
//
//			dp[i][j][1] = max(dp[i-1][j][1], dp[i-1][j-1][0]-prices[i])
//			dp[i][j][0] = max(dp[i-1][j][0], dp[i-1][j][1]+prices[i])
//
//		}
//	}
//	return dp[len(prices)-1][k][0]
//}

func maxProfit(k int, prices []int) int {
	return maxProfitAllInOne(k, prices, 0, 0)
}

func maxProfitAllInOne(maxK int, prices []int, coolDown int, fee int) int {
	if len(prices) <= 1 {
		return 0
	}

	if maxK > len(prices)/2 {
		return maxProfitKInf(prices, coolDown, fee)
	}

	dp := make([][][2]int, len(prices))
	for i := range dp {
		dp[i] = make([][2]int, maxK+1)
	}

	for i := 0; i < len(prices); i++ {
		for j := 1; j < maxK+1; j++ {
			if i < 1 {
				dp[i][j][1] = -prices[i] - fee
				dp[i][j][0] = 0
				continue
			}

			if i < coolDown+1 {
				dp[i][j][1] = max(-prices[i]-fee, dp[i-1][j][1])
				dp[i][j][0] = max(dp[i-1][j][1]+prices[i], dp[i-1][j][0])
				continue
			}

			// 初始化 dp[...][j][0]=0
			dp[i][j][1] = max(dp[i-coolDown-1][j-1][0]-prices[i]-fee, dp[i-1][j][1])
			dp[i][j][0] = max(dp[i-1][j][1]+prices[i], dp[i-1][j][0])
		}
	}
	return dp[len(prices)-1][maxK][0]
}

func maxProfitKInf(prices []int, coolDown int, fee int) int {
	dp := make([][2]int, len(prices))
	for i := 0; i < len(prices); i++ {
		if i < 1 {
			dp[i][1] = -prices[i] - fee
			dp[i][0] = 0
			continue
		}

		if i < coolDown+1 {
			dp[i][1] = max(-prices[i]-fee, dp[i-1][1])
			dp[i][0] = max(dp[i-1][1]+prices[i], dp[i-1][0])
			continue
		}

		dp[i][1] = max(dp[i-coolDown-1][0]-prices[i]-fee, dp[i-1][1])
		dp[i][0] = max(dp[i-1][1]+prices[i], dp[i-1][0])
	}
	return dp[len(prices)-1][0]
}

//leetcode submit region end(Prohibit modification and deletion)
