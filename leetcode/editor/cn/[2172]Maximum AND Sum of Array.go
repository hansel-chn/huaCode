//You are given an integer array nums of length n and an integer numSlots such
//that 2 * numSlots >= n. There are numSlots slots numbered from 1 to numSlots.
//
// You have to place all n integers into the slots such that each slot contains
//at most two numbers. The AND sum of a given placement is the sum of the bitwise
//AND of every number with its respective slot number.
//
//
// For example, the AND sum of placing the numbers [1, 3] into slot 1 and [4, 6]
// into slot 2 is equal to (1 AND 1) + (3 AND 1) + (4 AND 2) + (6 AND 2) = 1 + 1 +
// 0 + 2 = 4.
//
//
// Return the maximum possible AND sum of nums given numSlots slots.
//
//
// Example 1:
//
//
//Input: nums = [1,2,3,4,5,6], numSlots = 3
//Output: 9
//Explanation: One possible placement is [1, 4] into slot 1, [2, 6] into slot 2,
// and [3, 5] into slot 3.
//This gives the maximum AND sum of (1 AND 1) + (4 AND 1) + (2 AND 2) + (6 AND 2
//) + (3 AND 3) + (5 AND 3) = 1 + 0 + 2 + 2 + 3 + 1 = 9.
//
//
// Example 2:
//
//
//Input: nums = [1,3,10,4,7,1], numSlots = 9
//Output: 24
//Explanation: One possible placement is [1, 1] into slot 1, [3] into slot 3, [4
//] into slot 4, [7] into slot 7, and [10] into slot 9.
//This gives the maximum AND sum of (1 AND 1) + (1 AND 1) + (3 AND 3) + (4 AND 4
//) + (7 AND 7) + (10 AND 9) = 1 + 1 + 3 + 4 + 7 + 8 = 24.
//Note that slots 2, 5, 6, and 8 are empty which is permitted.
//
//
//
// Constraints:
//
//
// n == nums.length
// 1 <= numSlots <= 9
// 1 <= n <= 2 * numSlots
// 1 <= nums[i] <= 15
//
//
// Related Topics 位运算 数组 动态规划 状态压缩 👍 67 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
package main

import (
	"math/bits"
)

func maximumANDSum(nums []int, numSlots int) int {
	states := 1<<(numSlots*2) - 1

	dp := make([]int, states+1)

	for state := range states + 1 {
		//for state := 0; state <= states; state++ {
		count := bits.OnesCount64(uint64(state))
		if count > len(nums) {
			continue
		}

		for j := 0; j < numSlots*2; j++ {
			if ((state >> j) & 1) == 1 {
				dp[state] = max(dp[state], dp[state^(1<<j)]+((j>>1)+1)&nums[count-1])
			}
		}
	}

	rlt := 0
	for _, v := range dp {
		rlt = max(rlt, v)
	}

	return rlt
}

//leetcode submit region end(Prohibit modification and deletion)
