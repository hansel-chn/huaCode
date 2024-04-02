/**
Given an array nums of distinct integers, return all the possible permutations.
You can return the answer in any order.


 Example 1:
 Input: nums = [1,2,3]
Output: [[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]

 Example 2:
 Input: nums = [0,1]
Output: [[0,1],[1,0]]

 Example 3:
 Input: nums = [1]
Output: [[1]]


 Constraints:


 1 <= nums.length <= 6
 -10 <= nums[i] <= 10
 All the integers of nums are unique.


 Related Topics 数组 回溯 👍 2858 👎 0

*/

// leetcode submit region begin(Prohibit modification and deletion)
package main

func permute(nums []int) [][]int {
	rlt := make([][]int, 0)
	used := make([]bool, len(nums))

	var backTrack func(current []int)
	backTrack = func(current []int) {
		if len(current) == len(nums) {
			temp := make([]int, len(current))
			copy(temp, current)
			rlt = append(rlt, temp)
		}

		for i := 0; i < len(nums); i++ {
			if used[i] == true {
				continue
			}
			used[i] = true
			backTrack(append(current, nums[i]))
			used[i] = false
		}
	}
	backTrack([]int{})
	return rlt
}

//leetcode submit region end(Prohibit modification and deletion)
