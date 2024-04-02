/**
Given an integer array nums that may contain duplicates, return all possible
subsets (the power set).

 The solution set must not contain duplicate subsets. Return the solution in
any order.


 Example 1:
 Input: nums = [1,2,2]
Output: [[],[1],[1,2],[1,2,2],[2],[2,2]]

 Example 2:
 Input: nums = [0]
Output: [[],[0]]


 Constraints:


 1 <= nums.length <= 10
 -10 <= nums[i] <= 10


 Related Topics 位运算 数组 回溯 👍 1202 👎 0

*/

// leetcode submit region begin(Prohibit modification and deletion)
package main

import "sort"

func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)

	rlt:=make([][]int,0)
	current := make([]int, 0)
	var backtrack func(start int)
	backtrack = func(start int) {
		temp:=make([]int, len(current))
		copy(temp,current)
		rlt = append(rlt, temp)

		for i := start; i < len(nums); i++ {
			if i != start && nums[i] == nums[i-1] {
				continue
			}
			current = append(current, nums[i])
			backtrack(i+1)
			current = current[:len(current)-1]
		}
	}

	backtrack(0)
	return rlt
}

//func subsetsWithDup(nums []int) [][]int {
//	sort.Ints(nums)
//	var backtrack func(start int, target *[]int)
//	rlt := make([][]int, 0)
//	backtrack = func(start int, target *[]int) {
//		temp := make([]int, len(*target))
//		copy(temp, *target)
//		rlt = append(rlt, temp)
//
//		for i := start; i < len(nums); i++ {
//			if i+1 < len(nums) && nums[i] == nums[i+1] {
//				var repeat int
//				for i+1 < len(nums) && nums[i] == nums[i+1] {
//					i++
//					repeat++
//				}
//				//fmt.Println(repeat)
//				for j := 0; j < repeat+1; j++ {
//					*target = append(*target, nums[i])
//					backtrack(i+1, target)
//				}
//
//				*target = (*target)[:len(*target)-repeat-1]
//			} else {
//				*target = append(*target, nums[i])
//				backtrack(i+1, target)
//				*target = (*target)[:len(*target)-1]
//			}
//		}
//	}
//	initTarget := make([]int, 0)
//	backtrack(0, &initTarget)
//	return rlt
//}

//leetcode submit region end(Prohibit modification and deletion)
