/**
Given an array of distinct integers candidates and a target integer target,
return a list of all unique combinations of candidates where the chosen numbers sum
to target. You may return the combinations in any order.

 The same number may be chosen from candidates an unlimited number of times.
Two combinations are unique if the frequency of at least one of the chosen numbers
is different.

 The test cases are generated such that the number of unique combinations that
sum up to target is less than 150 combinations for the given input.


 Example 1:


Input: candidates = [2,3,6,7], target = 7
Output: [[2,2,3],[7]]
Explanation:
2 and 3 are candidates, and 2 + 2 + 3 = 7. Note that 2 can be used multiple
times.
7 is a candidate, and 7 = 7.
These are the only two combinations.


 Example 2:


Input: candidates = [2,3,5], target = 8
Output: [[2,2,2,2],[2,3,3],[3,5]]


 Example 3:


Input: candidates = [2], target = 1
Output: []



 Constraints:


 1 <= candidates.length <= 30
 2 <= candidates[i] <= 40
 All elements of candidates are distinct.
 1 <= target <= 40


 Related Topics 数组 回溯 👍 2751 👎 0

*/

// leetcode submit region begin(Prohibit modification and deletion)
package main

import "sort"

func combinationSum(candidates []int, target int) [][]int {
	rlt := make([][]int, 0)
	current := make([]int, 0)
	currentSum :=0

	sort.Ints(candidates)

	var backTrack func(start int)
	backTrack = func(start int) {
		if currentSum == target {
			temp := make([]int, len(current))
			copy(temp, current)
			rlt = append(rlt, temp)
		}

		for i := start; i < len(candidates); i++ {
			if currentSum+candidates[i]>target {
				return
			}

			current = append(current,candidates[i])
			currentSum+=candidates[i]
			backTrack(i)
			currentSum-=candidates[i]
			current =current[:len(current)-1]
		}
	}

	backTrack(0)
	return rlt
}

func SumSlice(s []int) int {
	sum := 0
	for _, v := range s {
		sum += v
	}
	return sum
}

//leetcode submit region end(Prohibit modification and deletion)
