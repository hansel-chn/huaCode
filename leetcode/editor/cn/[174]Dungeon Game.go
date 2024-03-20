/**
The demons had captured the princess and imprisoned her in the bottom-right
corner of a dungeon. The dungeon consists of m x n rooms laid out in a 2D grid. Our
valiant knight was initially positioned in the top-left room and must fight his
way through dungeon to rescue the princess.

 The knight has an initial health point represented by a positive integer. If
at any point his health point drops to 0 or below, he dies immediately.

 Some of the rooms are guarded by demons (represented by negative integers), so
the knight loses health upon entering these rooms; other rooms are either empty
(represented as 0) or contain magic orbs that increase the knight's health (
represented by positive integers).

 To reach the princess as quickly as possible, the knight decides to move only
rightward or downward in each step.

 Return the knight's minimum initial health so that he can rescue the princess.


 Note that any room can contain threats or power-ups, even the first room the
knight enters and the bottom-right room where the princess is imprisoned.


 Example 1:


Input: dungeon = [[-2,-3,3],[-5,-10,1],[10,30,-5]]
Output: 7
Explanation: The initial health of the knight must be at least 7 if he follows
the optimal path: RIGHT-> RIGHT -> DOWN -> DOWN.


 Example 2:


Input: dungeon = [[0]]
Output: 1



 Constraints:


 m == dungeon.length
 n == dungeon[i].length
 1 <= m, n <= 200
 -1000 <= dungeon[i][j] <= 1000


 Related Topics 数组 动态规划 矩阵 👍 820 👎 0

*/

// leetcode submit region begin(Prohibit modification and deletion)
package main

func calculateMinimumHP(dungeon [][]int) int {
	dp := make([][]int, len(dungeon))
	for i := range dp {
		dp[i] = make([]int, len(dungeon[0]))
	}

	for i := len(dungeon) - 1; i >= 0; i-- {
		for j := len(dungeon[0]) - 1; j >= 0; j-- {
			if i == len(dungeon)-1 && j == len(dungeon[0])-1 {
				dp[i][j] = max(1-dungeon[i][j], 1)
				continue
			}

			if i == len(dungeon)-1 {
				dp[i][j] = max(dp[i][j+1]-dungeon[i][j], 1)
				continue
			}

			if j == len(dungeon[0])-1 {
				dp[i][j] = max(dp[i+1][j]-dungeon[i][j], 1)
				continue
			}

			dp[i][j] = max(min(dp[i+1][j], dp[i][j+1])-dungeon[i][j], 1)
		}
	}
	return dp[0][0]
}

//func calculateMinimumHP(dungeon [][]int) int {
//	maxInt := 1000*200 + 1
//	dp := make([][]int, len(dungeon)+1)
//	for i := range dp {
//		dp[i] = make([]int, len(dungeon[0])+1)
//	}
//
//	for i := 0; i < len(dungeon)+1; i++ {
//		dp[i][len(dungeon[0])] = maxInt
//	}
//
//	for j := 0; j < len(dungeon[0])+1; j++ {
//		dp[len(dungeon)][j] = maxInt
//	}
//
//	for i := len(dungeon) - 1; i >= 0; i-- {
//		for j := len(dungeon[0]) - 1; j >= 0; j-- {
//			if i == len(dungeon)-1 && len(dungeon[0])-1 == j {
//				dp[i][j] = max(1-dungeon[i][j], 1)
//				continue
//			}
//
//			dp[i][j] = max(min(dp[i+1][j], dp[i][j+1])-dungeon[i][j], 1)
//		}
//	}
//	return dp[0][0]
//}

//leetcode submit region end(Prohibit modification and deletion)
