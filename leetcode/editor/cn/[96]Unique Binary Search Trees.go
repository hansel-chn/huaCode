/**
Given an integer n, return the number of structurally unique BST's (binary
search trees) which has exactly n nodes of unique values from 1 to n.


 Example 1:


Input: n = 3
Output: 5


 Example 2:


Input: n = 1
Output: 1



 Constraints:


 1 <= n <= 19


 Related Topics 树 二叉搜索树 数学 动态规划 二叉树 👍 2521 👎 0

*/

// leetcode submit region begin(Prohibit modification and deletion)
package main

func numTrees(n int) int {
	memo := make(map[int]int)
	var count func(n int) int
	count = func(n int) int {
		if v, ok := memo[n]; ok {
			return v
		}

		if n == 1 || n == 0 {
			return 1
		}

		res := 0
		for i := 1; i <= n; i++ {
			res +=count(i-1) * count(n-i)
		}

		memo[n] = res
		return res
	}
	a:=count(n)
	return a
}

//leetcode submit region end(Prohibit modification and deletion)
