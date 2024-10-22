//An array is squareful if the sum of every pair of adjacent elements is a
//perfect square.
//
// Given an integer array nums, return the number of permutations of nums that
//are squareful.
//
// Two permutations perm1 and perm2 are different if there is some index i such
//that perm1[i] != perm2[i].
//
//
// Example 1:
//
//
//Input: nums = [1,17,8]
//Output: 2
//Explanation: [1,8,17] and [17,8,1] are the valid permutations.
//
//
// Example 2:
//
//
//Input: nums = [2,2,2]
//Output: 1
//
//
//
// Constraints:
//
//
// 1 <= nums.length <= 12
// 0 <= nums[i] <= 10⁹
//
//
// Related Topics 位运算 数组 哈希表 数学 动态规划 回溯 状态压缩 👍 124 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
package main

import (
	"math"
	"sort"
)

func numSquarefulPerms(nums []int) int {
	sort.Ints(nums)

	used := make([]bool, len(nums))
	cur := 0
	cur1 := make([]int, 0)
	count := 0

	var traverse func(pre int)
	traverse = func(pre int) {
		if cur == len(nums) {
			count++
			return
		}

		for i := range nums {
			if used[i] == true {
				continue
			}

			if i > 0 && nums[i] == nums[i-1] && used[i-1] == false {
				continue
			}

			if cur == 0 || isPerfectSquare(pre+nums[i]) {
				cur += 1
				cur1 = append(cur1, nums[i])
				used[i] = true
				traverse(nums[i])
				cur -= 1
				cur1 = cur1[:len(cur1)-1]
				used[i] = false
			}
		}
	}
	traverse(0)
	return count
}

func isPerfectSquare(n int) bool {
	if n < 0 {
		return false
	}
	sqrt := math.Sqrt(float64(n))
	return sqrt == float64(int(sqrt))
}

//leetcode submit region end(Prohibit modification and deletion)
