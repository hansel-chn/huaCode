/**
Given an array nums of n integers, return an array of all the unique
quadruplets [nums[a], nums[b], nums[c], nums[d]] such that:


 0 <= a, b, c, d < n
 a, b, c, and d are distinct.
 nums[a] + nums[b] + nums[c] + nums[d] == target


 You may return the answer in any order.


 Example 1:


Input: nums = [1,0,-1,0,-2,2], target = 0
Output: [[-2,-1,1,2],[-2,0,0,2],[-1,0,0,1]]


 Example 2:


Input: nums = [2,2,2,2,2], target = 8
Output: [[2,2,2,2]]



 Constraints:


 1 <= nums.length <= 200
 -10⁹ <= nums[i] <= 10⁹
 -10⁹ <= target <= 10⁹


 Related Topics 数组 双指针 排序 👍 1916 👎 0

*/

// leetcode submit region begin(Prohibit modification and deletion)
package main

import (
	"sort"
)

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	n := 4
	if len(nums) < n {
		return [][]int{}
	}

	var nSum func(num []int, n int, target int) [][]int
	nSum = func(num []int, n int, target int) [][]int {
		if n == 2 {
			var doubletRlt [][]int
			for i, j := 0, len(num)-1; i < j; {
				if i > 0 && num[i] == num[i-1] {
					i++
					continue
				}

				if j < len(num)-1 && num[j] == num[j+1] {
					j--
					continue
				}

				if num[i]+num[j] < target {
					i++
				} else if num[i]+num[j] > target {
					j--
				} else {
					doubletRlt = append(doubletRlt, []int{num[i], num[j]})
					i++
					j--
				}
			}
			return doubletRlt
		} else {
			var rlt [][]int
			for i := 0; i < len(num); i++ {
				if i > 0 && num[i] == num[i-1] {
					continue
				}
				iterRlt := nSum(num[i+1:], n-1, target-num[i])
				for _, v := range iterRlt {
					rlt = append(rlt, append(v, num[i]))
				}
			}
			return rlt
		}
	}
	return nSum(nums, 4, target)
}

//leetcode submit region end(Prohibit modification and deletion)
