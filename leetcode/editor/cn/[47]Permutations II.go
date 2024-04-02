/**
Given a collection of numbers, nums, that might contain duplicates, return all
possible unique permutations in any order.


 Example 1:


Input: nums = [1,1,2]
Output:
[[1,1,2],
 [1,2,1],
 [2,1,1]]


 Example 2:


Input: nums = [1,2,3]
Output: [[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]



 Constraints:


 1 <= nums.length <= 8
 -10 <= nums[i] <= 10


 Related Topics 数组 回溯 👍 1558 👎 0

*/

// leetcode submit region begin(Prohibit modification and deletion)
package main

import (
	"math"
	"sort"
)

func permuteUnique(nums []int) [][]int {
	rlt := make([][]int, 0)
	current := make([]int, 0)
	//used := make([]bool, len(nums))
	used := 0
	sort.Ints(nums)

	var backTrack func()
	backTrack = func() {
		if len(current) == len(nums) {
			temp := make([]int, len(current))
			copy(temp, current)
			rlt = append(rlt, temp)
			return
		}

		pre:=math.MinInt
		for i := 0; i < len(nums); i++ {
			if used>>i&1 == 1 {
				continue
			}

			if pre==nums[i] {
				continue
			}

			pre=nums[i]
			used = used | 1<<i
			current = append(current, nums[i])
			backTrack()
			used = used ^ 1<<i
			current = current[:len(current)-1]
		}
	}
	backTrack()
	return rlt
}

//// 主要考虑怎么剪枝
//func permuteUnique(nums []int) [][]int {
//	rlt := make([][]int, 0)
//	current := make([]int, 0)
//	//used := make([]bool, len(nums))
//	used := 0
//	sort.Ints(nums)
//
//	var backTrack func()
//	backTrack = func() {
//		if len(current) == len(nums) {
//			temp := make([]int, len(current))
//			copy(temp, current)
//			rlt = append(rlt, temp)
//		}
//
//		for i := 0; i < len(nums); i++ {
//			//if used[i] == true {
//			//	continue
//			//}
//			if used>>i&1 == 1 {
//				continue
//			}
//
//			if i != 0 && nums[i-1] == nums[i] && !((used >> (i - 1) & 1) == 1) {
//				continue
//			}
//
//			//used[i]=true
//			used = used | 1<<i
//			current = append(current, nums[i])
//			backTrack()
//			used = used ^ 1<<i
//			//used[i]=false
//			current = current[:len(current)-1]
//		}
//	}
//	backTrack()
//	return rlt
//}

//leetcode submit region end(Prohibit modification and deletion)
