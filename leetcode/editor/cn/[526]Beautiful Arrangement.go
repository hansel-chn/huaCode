//Suppose you have n integers labeled 1 through n. A permutation of those n
//integers perm (1-indexed) is considered a beautiful arrangement if for every i (1 <=
// i <= n), either of the following is true:
//
//
// perm[i] is divisible by i.
// i is divisible by perm[i].
//
//
// Given an integer n, return the number of the beautiful arrangements that you
//can construct.
//
//
// Example 1:
//
//
//Input: n = 2
//Output: 2
//Explanation:
//The first beautiful arrangement is [1,2]:
//    - perm[1] = 1 is divisible by i = 1
//    - perm[2] = 2 is divisible by i = 2
//The second beautiful arrangement is [2,1]:
//    - perm[1] = 2 is divisible by i = 1
//    - i = 2 is divisible by perm[2] = 1
//
//
// Example 2:
//
//
//Input: n = 1
//Output: 1
//
//
//
// Constraints:
//
//
// 1 <= n <= 15
//
//
// Related Topics 位运算 数组 动态规划 回溯 状态压缩 👍 412 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
package main

func countArrangement(n int) int {
	var temp = 0

	var bits int
	var count int
	var traverse func(idx int)
	traverse = func(idx int) {
		if idx == n+1 {
			count++
			return
		}

		for i := 1; i <= n; i++ {
			if bits >> (i - 1) & 1==1 {
				continue
			}

			if i%idx == 0 || idx%i == 0 {
				bits = bits ^ 1<<(i-1)
				temp++
				traverse(idx + 1)
				temp--
				bits = bits ^ 1<<(i-1)
			}
		}
	}

	traverse(1)
	return count
}

//leetcode submit region end(Prohibit modification and deletion)
