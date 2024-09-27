//You are given an integer array rewardValues of length n, representing the
//values of rewards.
//
// Initially, your total reward x is 0, and all indices are unmarked. You are
//allowed to perform the following operation any number of times:
//
//
// Choose an unmarked index i from the range [0, n - 1].
// If rewardValues[i] is greater than your current total reward x, then add
//rewardValues[i] to x (i.e., x = x + rewardValues[i]), and mark the index i.
//
//
// Return an integer denoting the maximum total reward you can collect by
//performing the operations optimally.
//
//
// Example 1:
//
//
// Input: rewardValues = [1,1,3,3]
//
//
// Output: 4
//
// Explanation:
//
// During the operations, we can choose to mark the indices 0 and 2 in order,
//and the total reward will be 4, which is the maximum.
//
// Example 2:
//
//
// Input: rewardValues = [1,6,4,3,2]
//
//
// Output: 11
//
// Explanation:
//
// Mark the indices 0, 2, and 1 in order. The total reward will then be 11,
//which is the maximum.
//
//
// Constraints:
//
//
// 1 <= rewardValues.length <= 2000
// 1 <= rewardValues[i] <= 2000
//
//
// Related Topics 数组 动态规划 👍 13 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
package main

import (
	"math/big"
	"sort"
)

func maxTotalReward(rewardValues []int) int {
	sort.Ints(rewardValues)
	rewardValues = deduplicate(rewardValues)
	dp := big.NewInt(1)
	one := big.NewInt(1)

	for _, v := range rewardValues {
		mask := base().Lsh(
			base().And(
				base().Sub(
					base().Lsh(one, uint(v)),
					one),
				dp),
			uint(v))
		dp = dp.Or(dp, mask)
	}
	return dp.BitLen() - 1
}

func base() *big.Int {
	return new(big.Int)
}

func deduplicate(nums []int) []int {
	unique := []int{nums[0]}
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			continue
		}
		unique = append(unique, nums[i])
	}
	return unique
}

//leetcode submit region end(Prohibit modification and deletion)
