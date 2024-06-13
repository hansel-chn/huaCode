/**
Koko loves to eat bananas. There are n piles of bananas, the iᵗʰ pile has piles[
i] bananas. The guards have gone and will come back in h hours.

 Koko can decide her bananas-per-hour eating speed of k. Each hour, she chooses
some pile of bananas and eats k bananas from that pile. If the pile has less
than k bananas, she eats all of them instead and will not eat any more bananas
during this hour.

 Koko likes to eat slowly but still wants to finish eating all the bananas
before the guards return.

 Return the minimum integer k such that she can eat all the bananas within h
hours.


 Example 1:


Input: piles = [3,6,7,11], h = 8
Output: 4


 Example 2:


Input: piles = [30,11,23,4,20], h = 5
Output: 30


 Example 3:


Input: piles = [30,11,23,4,20], h = 6
Output: 23



 Constraints:


 1 <= piles.length <= 10⁴
 piles.length <= h <= 10⁹
 1 <= piles[i] <= 10⁹


 Related Topics 数组 二分查找 👍 615 👎 0

*/

// leetcode submit region begin(Prohibit modification and deletion)
package main

func minEatingSpeed(piles []int, h int) int {
	maxPile := 0
	for _, pile := range piles {
		maxPile = max(maxPile, pile)
	}

	var left, right int = 1, maxPile + 1

	for left < right {
		mid := left + (right-left)/2
		if EatingTime(piles, mid) == h {
			right = mid
		} else if EatingTime(piles, mid) < h {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

func EatingTime(piles []int, speed int) (h int) {
	for i := 0; i < len(piles); i++ {
		if piles[i]%speed != 0 {
			h = h + piles[i]/speed + 1
		} else {
			h = h + piles[i]/speed
		}
	}
	return h
}

//leetcode submit region end(Prohibit modification and deletion)
