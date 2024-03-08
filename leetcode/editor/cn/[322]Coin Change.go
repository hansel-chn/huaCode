/**
You are given an integer array coins representing coins of different
denominations and an integer amount representing a total amount of money.

 Return the fewest number of coins that you need to make up that amount. If
that amount of money cannot be made up by any combination of the coins, return -1.

 You may assume that you have an infinite number of each kind of coin.


 Example 1:


Input: coins = [1,2,5], amount = 11
Output: 3
Explanation: 11 = 5 + 5 + 1


 Example 2:


Input: coins = [2], amount = 3
Output: -1


 Example 3:


Input: coins = [1], amount = 0
Output: 0



 Constraints:


 1 <= coins.length <= 12
 1 <= coins[i] <= 2³¹ - 1
 0 <= amount <= 10⁴


 Related Topics 广度优先搜索 数组 动态规划 👍 2726 👎 0

*/

// leetcode submit region begin(Prohibit modification and deletion)
package main

//func coinChange(coins []int, amount int) int {
//
//	dp := make([]int, amount+1)
//
//	for i := 1; i < amount+1; i++ {
//		dp[i] = math.MaxInt32
//	}
//
//	for i := 1; i <= amount; i++ {
//		for _, coin := range coins {
//			if i-coin >= 0 {
//				dp[i] = min(dp[i-coin]+1, dp[i])
//			}
//		}
//	}
//
//	if dp[amount]==math.MaxInt32{
//		return -1
//	}else {
//		return dp[amount]
//	}
//}

func coinChange(coins []int, amount int) int {

	dp := make([]int, amount+1)

	for i := 1; i < amount+1; i++ {
		// 其实取amount+1就行，硬币最小为1
		dp[i] = amount + 1
	}

	for i := 1; i <= amount; i++ {
		for _, coin := range coins {
			if i-coin >= 0 {
				dp[i] = min(dp[i-coin]+1, dp[i])
			}
		}
	}

	if dp[amount] == amount+1 {
		return -1
	} else {
		return dp[amount]
	}
}

//leetcode submit region end(Prohibit modification and deletion)
