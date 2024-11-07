//In LeetCode Store, there are n items to sell. Each item has a price. However,
//there are some special offers, and a special offer consists of one or more
//different kinds of items with a sale price.
//
// You are given an integer array price where price[i] is the price of the iᵗʰ
//item, and an integer array needs where needs[i] is the number of pieces of the iᵗ
//ʰ item you want to buy.
//
// You are also given an array special where special[i] is of size n + 1 where
//special[i][j] is the number of pieces of the jᵗʰ item in the iᵗʰ offer and
//special[i][n] (i.e., the last integer in the array) is the price of the iᵗʰ offer.
//
// Return the lowest price you have to pay for exactly certain items as given,
//where you could make optimal use of the special offers. You are not allowed to
//buy more items than you want, even if that would lower the overall price. You
//could use any of the special offers as many times as you want.
//
//
// Example 1:
//
//
//Input: price = [2,5], special = [[3,0,5],[1,2,10]], needs = [3,2]
//Output: 14
//Explanation: There are two kinds of items, A and B. Their prices are $2 and $5
// respectively.
//In special offer 1, you can pay $5 for 3A and 0B
//In special offer 2, you can pay $10 for 1A and 2B.
//You need to buy 3A and 2B, so you may pay $10 for 1A and 2B (special offer #2)
//, and $4 for 2A.
//
//
// Example 2:
//
//
//Input: price = [2,3,4], special = [[1,1,0,4],[2,2,1,9]], needs = [1,2,1]
//Output: 11
//Explanation: The price of A is $2, and $3 for B, $4 for C.
//You may pay $4 for 1A and 1B, and $9 for 2A ,2B and 1C.
//You need to buy 1A ,2B and 1C, so you may pay $4 for 1A and 1B (special offer
//#1), and $3 for 1B, $4 for 1C.
//You cannot add more items, though only $9 for 2A ,2B and 1C.
//
//
//
// Constraints:
//
//
// n == price.length == needs.length
// 1 <= n <= 6
// 0 <= price[i], needs[i] <= 10
// 1 <= special.length <= 100
// special[i].length == n + 1
// 0 <= special[i][j] <= 50
// The input is generated that at least one of special[i][j] is non-zero for 0 <
//= j <= n - 1.
//
//
// Related Topics 位运算 记忆化搜索 数组 动态规划 回溯 状态压缩 👍 396 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
package main

import (
	"math"
)

func shoppingOffers(price []int, special [][]int, needs []int) int {
	special = filterSpecial(price, special)
	nonSpecial := 0
	for i := 0; i < len(needs); i++ {
		nonSpecial += needs[i] * price[i]
	}

	if len(special) == 0 {
		return nonSpecial
	}

	speLen := len(special)
	priceIdx := len(special[0]) - 1

	minCost := math.MaxInt64
	cost := 0
	var traverse func(pos int)
	traverse = func(pos int) {
		for i := pos; i < speLen; i++ {
			if !checkMinusSpe(special[i], needs) {
				restCost := costReminder(price, needs)
				cost += restCost
				minCost = min(minCost, cost)
				cost -= restCost
				continue
			}

			minusSpe(special[i], needs)
			cost += special[i][priceIdx]
			traverse(pos)
			cost -= special[i][priceIdx]
			plusSpe(special[i], needs)

		}
	}
	traverse(0)

	minCost = min(minCost, nonSpecial)
	return minCost
}

func filterSpecial(price []int, special [][]int) (newSpecial [][]int) {
	for i := 0; i < len(special); i++ {
		cost := 0
		for j := 0; j < len(special[0])-1; j++ {
			cost += special[i][j] * price[j]
		}
		if cost > special[i][len(special[i])-1] {
			newSpecial = append(newSpecial, special[i])
		}
	}
	return newSpecial
}

func costReminder(price []int, needs []int) int {
	restCost := 0
	for i, need := range needs {
		restCost += need * price[i]
	}
	return restCost
}

func checkMinusSpe(discount []int, needs []int) bool {
	for i := 0; i < len(needs); i++ {
		if needs[i]-discount[i] < 0 {
			return false
		}
	}
	return true
}

func minusSpe(discount []int, needs []int) {
	for i := 0; i < len(needs); i++ {
		needs[i] -= discount[i]
	}
}

func plusSpe(discount []int, needs []int) {
	for i := 0; i < len(needs); i++ {
		needs[i] += discount[i]
	}
}

//leetcode submit region end(Prohibit modification and deletion)
