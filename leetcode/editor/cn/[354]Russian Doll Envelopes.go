/**
You are given a 2D array of integers envelopes where envelopes[i] = [wi, hi]
represents the width and the height of an envelope.

 One envelope can fit into another if and only if both the width and height of
one envelope are greater than the other envelope's width and height.

 Return the maximum number of envelopes you can Russian doll (i.e., put one
inside the other).

 Note: You cannot rotate an envelope.


 Example 1:


Input: envelopes = [[5,4],[6,4],[6,7],[2,3]]
Output: 3
Explanation: The maximum number of envelopes you can Russian doll is 3 ([2,3] =>
 [5,4] => [6,7]).


 Example 2:


Input: envelopes = [[1,1],[1,1],[1,1]]
Output: 1



 Constraints:


 1 <= envelopes.length <= 10⁵
 envelopes[i].length == 2
 1 <= wi, hi <= 10⁵


 Related Topics 数组 二分查找 动态规划 排序 👍 996 👎 0

*/

// leetcode submit region begin(Prohibit modification and deletion)
package main

import (
	"math"
	"sort"
)

func maxEnvelopes(envelopes [][]int) int {
	if len(envelopes) == 0 {
		return 0
	}

	sort.Slice(envelopes, func(i, j int) bool {
		if envelopes[i][0] < envelopes[j][0] {
			return true
		} else if envelopes[i][0] == envelopes[j][0] && envelopes[i][1] > envelopes[j][1] {
			return true
		} else {
			return false
		}
	})
	dp := make([]int, len(envelopes))
	for i := range dp {
		dp[i] = 1
	}

	assist := make([]int, 1)
	assist[0] = math.MinInt

	for i := 0; i < len(dp); i++ {
		idx := sort.SearchInts(assist, envelopes[i][1])
		if idx == len(assist) {
			assist = append(assist, envelopes[i][1])
		} else {
			assist[idx] = envelopes[i][1]
		}
		dp[i] = idx
	}

	res := 0
	for _, count := range dp {
		res = max(res, count)
	}
	return res
}

//func maxEnvelopes(envelopes [][]int) int {
//	sort.Slice(envelopes, func(i, j int) bool {
//		if envelopes[i][0] < envelopes[j][0] {
//			return true
//		} else if envelopes[i][0] == envelopes[j][0] && envelopes[i][1] > envelopes[j][1] {
//			return true
//		} else {
//			return false
//		}
//	})
//	dp := make([]int, len(envelopes))
//	for i := range dp {
//		dp[i] = 1
//	}
//
//	for i := 1; i < len(dp); i++ {
//		for j := 0; j < i; j++ {
//			if envelopes[j][1] < envelopes[i][1] {
//				dp[i] = max(dp[i], dp[j]+1)
//			}
//		}
//	}
//
//	res := 0
//	for _, count := range dp {
//		res = max(res, count)
//	}
//	return res
//}

//leetcode submit region end(Prohibit modification and deletion)
