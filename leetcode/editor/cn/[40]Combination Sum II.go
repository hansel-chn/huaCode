/**
Given a collection of candidate numbers (candidates) and a target number (
target), find all unique combinations in candidates where the candidate numbers sum
to target.

 Each number in candidates may only be used once in the combination.

 Note: The solution set must not contain duplicate combinations.


 Example 1:


Input: candidates = [10,1,2,7,6,1,5], target = 8
Output:
[
[1,1,6],
[1,2,5],
[1,7],
[2,6]
]


 Example 2:


Input: candidates = [2,5,2,1,2], target = 5
Output:
[
[1,2,2],
[5]
]



 Constraints:


 1 <= candidates.length <= 100
 1 <= candidates[i] <= 50
 1 <= target <= 30


 Related Topics 数组 回溯 👍 1532 👎 0

*/

// leetcode submit region begin(Prohibit modification and deletion)
package main

import "sort"

func combinationSum2(candidates []int, target int) [][]int {
	rlt := make([][]int, 0)
	current := make([]int, 0)
	currentValue := 0
	sort.Ints(candidates)

	var backTrack func(start int)
	backTrack = func(start int) {
		if currentValue == target {
			temp := make([]int, len(current))
			copy(temp, current)
			rlt = append(rlt, temp)
		}

		for i := start; i < len(candidates); i++ {
			if i != start && candidates[i] == candidates[i-1] {
				continue
			}

			if currentValue+candidates[i] > target {
				return
			}

			current = append(current, candidates[i])
			currentValue += candidates[i]
			backTrack(i + 1)
			current = current[:len(current)-1]
			currentValue -= candidates[i]
		}
	}
	backTrack(0)
	return rlt
}

//leetcode submit region end(Prohibit modification and deletion)
