//Given two integers n and k, return an array of all the integers of length n
//where the difference between every two consecutive digits is k. You may return
//the answer in any order.
//
// Note that the integers should not have leading zeros. Integers as 02 and 043
//are not allowed.
//
//
// Example 1:
//
//
//Input: n = 3, k = 7
//Output: [181,292,707,818,929]
//Explanation: Note that 070 is not a valid number, because it has leading
//zeroes.
//
//
// Example 2:
//
//
//Input: n = 2, k = 1
//Output: [10,12,21,23,32,34,43,45,54,56,65,67,76,78,87,89,98]
//
//
//
// Constraints:
//
//
// 2 <= n <= 9
// 0 <= k <= 9
//
//
// Related Topics 广度优先搜索 回溯 👍 106 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
package main

func numsSameConsecDiff(n int, k int) []int {
	rlt := make([]int, 0)
	var traverse func(last int, remain int, path int)

	traverse = func(last int, remain int, path int) {
		if remain == 0 {
			rlt = append(rlt, path)
			return
		}

		nums := getNextNum(last, k)

		for _, num := range nums {
			traverse(num, remain-1, path*10+num)
		}
	}

	nums := getFirstNum(k)
	for _, num := range nums {
		traverse(num, n-1, num)
	}

	return rlt
}

func getFirstNum(k int) []int {
	if k == 0 {
		return []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	}

	rlt := make([]int, 0)
	for i := 0; i < k; i++ {
		num := i
		for num < 10 {
			if num != 0 {
				rlt = append(rlt, num)
			}
			num += k
		}
	}
	return rlt
}

func getNextNum(num int, k int) []int {
	if k == 0 {
		return []int{num}
	}

	rlt := make([]int, 0)

	if num+k < 10 {
		rlt = append(rlt, num+k)
	}

	if num-k >= 0 {
		rlt = append(rlt, num-k)
	}
	return rlt
}

//leetcode submit region end(Prohibit modification and deletion)
