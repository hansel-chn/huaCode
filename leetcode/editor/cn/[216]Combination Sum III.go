/**
Find all valid combinations of k numbers that sum up to n such that the
following conditions are true:


 Only numbers 1 through 9 are used.
 Each number is used at most once.


 Return a list of all possible valid combinations. The list must not contain
the same combination twice, and the combinations may be returned in any order.


 Example 1:


Input: k = 3, n = 7
Output: [[1,2,4]]
Explanation:
1 + 2 + 4 = 7
There are no other valid combinations.

 Example 2:


Input: k = 3, n = 9
Output: [[1,2,6],[1,3,5],[2,3,4]]
Explanation:
1 + 2 + 6 = 9
1 + 3 + 5 = 9
2 + 3 + 4 = 9
There are no other valid combinations.


 Example 3:


Input: k = 4, n = 1
Output: []
Explanation: There are no valid combinations.
Using 4 different numbers in the range [1,9], the smallest sum we can get is 1+2
+3+4 = 10 and since 10 > 1, there are no valid combination.



 Constraints:


 2 <= k <= 9
 1 <= n <= 60


 Related Topics 数组 回溯 👍 815 👎 0

*/

// leetcode submit region begin(Prohibit modification and deletion)
package main

func combinationSum3(k int, n int) [][]int {
	rlt := make([][]int, 0)
	current := make([]int, 0)

	var backTrack func(start int)
	backTrack = func(start int) {
		if len(current) == k {
			if SliceSum(current) == n {
				temp := make([]int, len(current))
				copy(temp, current)
				rlt = append(rlt, temp)
			}
			return
		}

		for i := start; i <= 9; i++ {
			current = append(current, i)
			backTrack(i + 1)
			current = current[:len(current)-1]
		}
	}
	backTrack(1)
	return rlt
}

func SliceSum(s []int) int {
	sum := 0
	for _, v := range s {
		sum += v
	}
	return sum
}

//leetcode submit region end(Prohibit modification and deletion)
