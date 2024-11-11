//You are given an integer array cookies, where cookies[i] denotes the number
//of cookies in the iᵗʰ bag. You are also given an integer k that denotes the
//number of children to distribute all the bags of cookies to. All the cookies in the
//same bag must go to the same child and cannot be split up.
//
// The unfairness of a distribution is defined as the maximum total cookies
//obtained by a single child in the distribution.
//
// Return the minimum unfairness of all distributions.
//
//
// Example 1:
//
//
//Input: cookies = [8,15,10,20,8], k = 2
//Output: 31
//Explanation: One optimal distribution is [8,15,8] and [10,20]
//- The 1ˢᵗ child receives [8,15,8] which has a total of 8 + 15 + 8 = 31
//cookies.
//- The 2ⁿᵈ child receives [10,20] which has a total of 10 + 20 = 30 cookies.
//The unfairness of the distribution is max(31,30) = 31.
//It can be shown that there is no distribution with an unfairness less than 31.
//
//
//
// Example 2:
//
//
//Input: cookies = [6,1,3,2,2,4,1,2], k = 3
//Output: 7
//Explanation: One optimal distribution is [6,1], [3,2,2], and [4,1,2]
//- The 1ˢᵗ child receives [6,1] which has a total of 6 + 1 = 7 cookies.
//- The 2ⁿᵈ child receives [3,2,2] which has a total of 3 + 2 + 2 = 7 cookies.
//- The 3ʳᵈ child receives [4,1,2] which has a total of 4 + 1 + 2 = 7 cookies.
//The unfairness of the distribution is max(7,7,7) = 7.
//It can be shown that there is no distribution with an unfairness less than 7.
//
//
//
// Constraints:
//
//
// 2 <= cookies.length <= 8
// 1 <= cookies[i] <= 10⁵
// 2 <= k <= cookies.length
//
//
// Related Topics 位运算 数组 动态规划 回溯 状态压缩 👍 99 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
package main

import "math"

func distributeCookies(cookies []int, k int) int {
	childCookies := make([]int, k)
	minUnfair:=math.MaxInt32
	var backtrack func(pos int)
	backtrack = func(pos int) {
		if pos == len(cookies) {
			minUnfair=min(minUnfair,maxCookieChild2305(childCookies))
			return
		}

		existMap := make(map[int]struct{})
		for i := range childCookies {
			if _, ok := existMap[childCookies[i]]; ok {
				continue
			}

			existMap[childCookies[i]] = struct{}{}
			childCookies[i] += cookies[pos]
			backtrack(pos + 1)
			childCookies[i] -= cookies[pos]
		}
	}
	backtrack(0)
	return minUnfair
}

func maxCookieChild2305(childCookies []int) int {
	maxCookie:=0
	for _, cookie := range childCookies {
		maxCookie=max(maxCookie,cookie)
	}
	return maxCookie
}

//leetcode submit region end(Prohibit modification and deletion)
