/**
Given an integer array nums, return all the triplets [nums[i], nums[j], nums[k]]
 such that i != j, i != k, and j != k, and nums[i] + nums[j] + nums[k] == 0.

 Notice that the solution set must not contain duplicate triplets.


 Example 1:


Input: nums = [-1,0,1,2,-1,-4]
Output: [[-1,-1,2],[-1,0,1]]
Explanation:
nums[0] + nums[1] + nums[2] = (-1) + 0 + 1 = 0.
nums[1] + nums[2] + nums[4] = 0 + 1 + (-1) = 0.
nums[0] + nums[3] + nums[4] = (-1) + 2 + (-1) = 0.
The distinct triplets are [-1,0,1] and [-1,-1,2].
Notice that the order of the output and the order of the triplets does not
matter.


 Example 2:


Input: nums = [0,1,1]
Output: []
Explanation: The only possible triplet does not sum up to 0.


 Example 3:


Input: nums = [0,0,0]
Output: [[0,0,0]]
Explanation: The only possible triplet sums up to 0.



 Constraints:


 3 <= nums.length <= 3000
 -10⁵ <= nums[i] <= 10⁵


 Related Topics 数组 双指针 排序 👍 6877 👎 0

*/

// leetcode submit region begin(Prohibit modification and deletion)

package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	fmt.Println(nums)
	tripletRlt := make([][]int, 0)

	var twoSum func(target int, curNums []int) [][]int
	twoSum = func(target int, curNums []int) [][]int {
		var doubletRlt [][]int
		for i, j := 0, len(curNums)-1; i < j; {
			if i > 0 && curNums[i] == curNums[i-1] {
				i++
				continue
			}

			if j < len(curNums)-1 && curNums[j] == curNums[j+1] {
				j--
				continue
			}

			if curNums[i]+curNums[j] < target {
				i++
			} else if curNums[i]+curNums[j] > target {
				j--
			} else {
				doubletRlt = append(doubletRlt, []int{curNums[i], curNums[j]})
				i++
				j--
			}
		}
		return doubletRlt
	}

	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		doubletRlt := twoSum(-nums[i], nums[i+1:])
		fmt.Println(doubletRlt)
		for _, v := range doubletRlt {
			tripletRlt = append(tripletRlt, append(v, nums[i]))
		}
	}
	return tripletRlt
}

//leetcode submit region end(Prohibit modification and deletion)
