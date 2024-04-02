/**
Given an integer array nums and an integer k, return true if it is possible to
divide this array into k non-empty subsets whose sums are all equal.


 Example 1:


Input: nums = [4,3,2,3,5,2,1], k = 4
Output: true
Explanation: It is possible to divide it into 4 subsets (5), (1, 4), (2,3), (2,3
) with equal sums.


 Example 2:


Input: nums = [1,2,3,4], k = 3
Output: false



 Constraints:


 1 <= k <= nums.length <= 16
 1 <= nums[i] <= 10⁴
 The frequency of each element is in the range [1, 4].


 Related Topics 位运算 记忆化搜索 数组 动态规划 回溯 状态压缩 👍 1020 👎 0

*/

// leetcode submit region begin(Prohibit modification and deletion)
package main

import "sort"

func canPartitionKSubsets(nums []int, k int) bool {
	sum := 0
	for _, num := range nums {
		sum += num
	}

	if sum%k != 0 {
		return false
	}

	if k > len(nums) {
		return false
	}

	sort.Ints(nums)
	target := sum / k
	used := 0
	memo := make(map[int]struct{})
	//memo := make(map[int]bool)

	var backtrack func(current int, bucket int, start int) bool
	backtrack = func(current int, bucket int, start int) bool {
		if bucket == 0 {
			return true
		}

		if current == target {
			// 逻辑是走过的路一定是失败的
			if _, ok := memo[used]; ok {
				return false
			} else {
				memo[used] = struct{}{}
			}
			return backtrack(0, bucket-1, 0)

			// // 逻辑是记录状态
			//res, ok := memo[used] //缓存结果
			//if ok {
			//	return res
			//}
			//res = backtrack(0, bucket-1, 0)
			//memo[used] = res
			//return res
		}

		for i := start; i < len(nums); i++ {
			if used>>i&1 == 1 {
				continue
			}

			if current+nums[i] > target {
				return false
			}

			used = used | 1<<i
			if backtrack(current+nums[i], bucket, i+1) {
				return true
			}
			used = used ^ 1<<i
		}
		return false
	}
	return backtrack(0, k, 0)
}

//func canPartitionKSubsets(nums []int, k int) bool {
//	sum := 0
//	for _, num := range nums {
//		sum += num
//	}
//
//	if sum%k != 0 {
//		return false
//	}
//
//	if k > len(nums) {
//		return false
//	}
//
//	sort.Ints(nums)
//	target := sum / k
//	used := make([]bool, len(nums))
//
//	var backtrack func(current int, bucket int, start int) bool
//	backtrack = func(current int, bucket int, start int) bool {
//		if bucket == 0 {
//			return true
//		}
//
//		if current == target {
//			return backtrack(0, bucket-1, 0)
//		}
//
//		for i := start; i < len(nums); i++ {
//			if used[i] == true {
//				continue
//			}
//
//			if current+nums[i] > target {
//				return false
//			}
//
//			used[i] = true
//			if backtrack(current+nums[i], bucket, i+1) {
//				return true
//			}
//			used[i] = false
//		}
//		return false
//	}
//	return backtrack(0, k, 0)
//}

//leetcode submit region end(Prohibit modification and deletion)
