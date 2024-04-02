/**
Given an integer array nums of unique elements, return all possible subsets (
the power set).

 The solution set must not contain duplicate subsets. Return the solution in
any order.


 Example 1:


Input: nums = [1,2,3]
Output: [[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]


 Example 2:


Input: nums = [0]
Output: [[],[0]]



 Constraints:


 1 <= nums.length <= 10
 -10 <= nums[i] <= 10
 All the numbers of nums are unique.


 Related Topics 位运算 数组 回溯 👍 2270 👎 0

*/

// leetcode submit region begin(Prohibit modification and deletion)
package main

func subsets(nums []int) [][]int {
	rlt := make([][]int, 0)
	current := make([]int, 0)

	var backTrack func(start int)
	backTrack = func(start int) {
		temp := make([]int, len(current))
		copy(temp, current)
		rlt = append(rlt, temp)

		for i := start; i < len(nums); i++ {
			current = append(current, nums[i])
			backTrack(i + 1)
			current = current[:len(current)-1]
		}
	}
	backTrack(0)
	return rlt
}

//leetcode submit region end(Prohibit modification and deletion)
