/**
You are given k identical eggs and you have access to a building with n floors
labeled from 1 to n.

 You know that there exists a floor f where 0 <= f <= n such that any egg
dropped at a floor higher than f will break, and any egg dropped at or below floor f
will not break.

 Each move, you may take an unbroken egg and drop it from any floor x (where 1 <
= x <= n). If the egg breaks, you can no longer use it. However, if the egg
does not break, you may reuse it in future moves.

 Return the minimum number of moves that you need to determine with certainty
what the value of f is.


 Example 1:


Input: k = 1, n = 2
Output: 2
Explanation:
Drop the egg from floor 1. If it breaks, we know that f = 0.
Otherwise, drop the egg from floor 2. If it breaks, we know that f = 1.
If it does not break, then we know f = 2.
Hence, we need at minimum 2 moves to determine with certainty what the value of
f is.


 Example 2:


Input: k = 2, n = 6
Output: 3


 Example 3:


Input: k = 3, n = 14
Output: 4



 Constraints:


 1 <= k <= 100
 1 <= n <= 10⁴


 Related Topics 数学 二分查找 动态规划 👍 986 👎 0

*/

// leetcode submit region begin(Prohibit modification and deletion)
package main

//func superEggDrop(k int, n int) int {
//	memo := make(map[[2]int]int)
//	var dp func(k int, n int) int
//	dp = func(k int, n int) int {
//		if n == 1 {
//			return 1
//		}
//
//		if k == 1 {
//			return n
//		}
//
//		if v, ok := memo[[2]int{k, n}]; ok {
//			return v
//		}
//
//		left := 1
//		right := n
//		for left+1 < right {
//			mid := (left + right) / 2
//			if dp(k-1, mid-1) < dp(k, n-mid) {
//				left = mid
//			} else if dp(k-1, mid-1) > dp(k, n-mid) {
//				right = mid
//			} else {
//				left, right = mid, mid
//			}
//		}
//		//rlt := 1+min(max(dp(k, n-left),dp(k-1, left -1)), max(dp(k-1, right-1),dp(k, n-right)))
//		rlt := 1 + min(dp(k, n-left), dp(k-1, right-1))
//		memo[[2]int{k, n}] = rlt
//		return rlt
//	}
//	return dp(k, n)
//}

func superEggDrop(k int, n int) int {
	dp := make([][]int, k+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
		dp[i][1] = 1
	}

	for i := 1; i < n+1; i++ {
		dp[1][i] = i
	}

	for i := 2; i < k+1; i++ {
		for j := 2; j < n+1; j++ {
			left := 1
			right := j
			for left+1 < right {
				mid := (left + right) / 2
				if dp[i-1][mid-1] < dp[i][j-mid] {
					left = mid
				} else if dp[i-1][mid-1] > dp[i][j-mid] {
					right = mid
				} else {
					left, right = mid, mid
				}
			}
			dp[i][j] = 1 + min(dp[i][j-left], dp[i-1][right-1])
		}
	}
	return dp[k][n]
}

//leetcode submit region end(Prohibit modification and deletion)
