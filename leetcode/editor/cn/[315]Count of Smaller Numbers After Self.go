/**
Given an integer array nums, return an integer array counts where counts[i] is
the number of smaller elements to the right of nums[i].


 Example 1:


Input: nums = [5,2,6,1]
Output: [2,1,1,0]
Explanation:
To the right of 5 there are 2 smaller elements (2 and 1).
To the right of 2 there is only 1 smaller element (1).
To the right of 6 there is 1 smaller element (1).
To the right of 1 there is 0 smaller element.


 Example 2:


Input: nums = [-1]
Output: [0]


 Example 3:


Input: nums = [-1,-1]
Output: [0,0]



 Constraints:


 1 <= nums.length <= 10⁵
 -10⁴ <= nums[i] <= 10⁴


 Related Topics 树状数组 线段树 数组 二分查找 分治 有序集合 归并排序 👍 1078 👎 0

*/

// leetcode submit region begin(Prohibit modification and deletion)
package main

func countSmaller(nums []int) []int {
	count := make([]int, len(nums))
	newNums := make([][2]int, 0)
	for i, num := range nums {
		newNums = append(newNums, [2]int{num, i})
	}

	var merge func(left, mid, right int)
	merge = func(left, mid, right int) {
		leftNums := make([][2]int, mid-left+1)
		leftNumsLen := mid - left + 1
		copy(leftNums, newNums[left:mid+1])

		var pos, i, j = left, 0, mid + 1

		for i < leftNumsLen && j <= right {
			if leftNums[i][0] <= newNums[j][0] {
				count[leftNums[i][1]] += j - (mid + 1)
				newNums[pos] = leftNums[i]
				i++
				pos++
			} else {
				newNums[pos] = newNums[j]
				j++
				pos++
			}
		}

		if j == right+1 {
			for i < leftNumsLen {
				count[leftNums[i][1]] += j - (mid + 1)
				newNums[pos] = leftNums[i]
				i++
				pos++
			}
		}
	}

	var sort func(i, j int)
	sort = func(i, j int) {
		if i == j {
			return
		}

		mid := i + (j-i)/2
		sort(i, mid)
		sort(mid+1, j)
		merge(i, mid, j)
	}

	sort(0, len(newNums)-1)
	return count
}

// // low performance
//var count []int
//
//func countSmaller(nums []int) []int {
//	count = make([]int, len(nums))
//	newNums := make([][2]int, 0)
//	for i, num := range nums {
//		newNums = append(newNums, [2]int{num, i})
//	}
//
//	sort(newNums)
//	return count
//}
//
//func sort(newNums [][2]int) [][2]int {
//	if len(newNums) == 1 {
//		return newNums
//	}
//
//	mid := len(newNums) / 2
//
//	left := sort(newNums[:mid])
//	right := sort(newNums[mid:])
//	return merge(left, right)
//}
//
//func merge(left, right [][2]int) [][2]int {
//	var leftIdx, rightIdx int
//	var rlt [][2]int
//	for leftIdx < len(left) && rightIdx < len(right) {
//		if left[leftIdx][0] <= right[rightIdx][0] {
//			count[left[leftIdx][1]] += rightIdx
//			rlt = append(rlt, left[leftIdx])
//			leftIdx++
//		} else {
//			rlt = append(rlt, right[rightIdx])
//			rightIdx++
//		}
//	}
//
//	if leftIdx == len(left) {
//		rlt = append(rlt, right[rightIdx:]...)
//	}
//
//	if rightIdx == len(right) {
//		for leftIdx < len(left) {
//			count[left[leftIdx][1]] += rightIdx
//			rlt = append(rlt, left[leftIdx])
//			leftIdx++
//		}
//	}
//	return rlt
//}

//leetcode submit region end(Prohibit modification and deletion)
