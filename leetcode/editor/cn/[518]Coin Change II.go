/**
You are given an integer array coins representing coins of different
denominations and an integer amount representing a total amount of money.

 Return the number of combinations that make up that amount. If that amount of
money cannot be made up by any combination of the coins, return 0.

 You may assume that you have an infinite number of each kind of coin.

 The answer is guaranteed to fit into a signed 32-bit integer.


 Example 1:


Input: amount = 5, coins = [1,2,5]
Output: 4
Explanation: there are four ways to make up the amount:
5=5
5=2+2+1
5=2+1+1+1
5=1+1+1+1+1


 Example 2:


Input: amount = 3, coins = [2]
Output: 0
Explanation: the amount of 3 cannot be made up just with coins of 2.


 Example 3:


Input: amount = 10, coins = [10]
Output: 1



 Constraints:


 1 <= coins.length <= 300
 1 <= coins[i] <= 5000
 All the values of coins are unique.
 0 <= amount <= 5000


 Related Topics 数组 动态规划 👍 1235 👎 0

*/

// leetcode submit region begin(Prohibit modification and deletion)
package main

func change(amount int, coins []int) int {
	//if amount == 0 {
	//	return 0
	//}

	dp := make([][]int, len(coins)+1)
	for i := range dp {
		dp[i] = make([]int, amount+1)
	}

	for i := 0; i < len(coins)+1; i++ {
		dp[i][0] = 1
	}

	for i := 1; i < len(coins)+1; i++ {
		for j := 1; j < amount+1; j++ {
			//for k := i - 1; k > 0; k-- {
			//	dp[i][j] += dp[k][j]
			//}
			dp[i][j] += dp[i-1][j]
			for iter := j - coins[i-1]; iter >= 0; iter = iter - coins[i-1] {
				dp[i][j] += dp[i-1][iter]
			}
		}
	}
	return dp[len(coins)][amount]
}

//leetcode submit region end(Prohibit modification and deletion)
