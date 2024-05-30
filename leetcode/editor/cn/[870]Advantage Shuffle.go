/**
You are given two integer arrays nums1 and nums2 both of the same length. The
advantage of nums1 with respect to nums2 is the number of indices i for which
nums1[i] > nums2[i].

 Return any permutation of nums1 that maximizes its advantage with respect to
nums2.


 Example 1:
 Input: nums1 = [2,7,11,15], nums2 = [1,10,4,11]
Output: [2,11,7,15]

 Example 2:
 Input: nums1 = [12,24,8,32], nums2 = [13,25,32,11]
Output: [24,32,8,12]


 Constraints:


 1 <= nums1.length <= 10⁵
 nums2.length == nums1.length
 0 <= nums1[i], nums2[i] <= 10⁹


 Related Topics 贪心 数组 双指针 排序 👍 420 👎 0

*/

// leetcode submit region begin(Prohibit modification and deletion)
package main

import (
	"container/heap"
	"sort"
)

func advantageCount(nums1 []int, nums2 []int) []int {

	// sort
	rlt := make([]int, len(nums1))
	sort.Ints(nums1)

	//sortNums2 := make([][2]int, len(nums2))
	//for i, v := range nums2 {
	//	sortNums2[i] = [2]int{i, v}
	//}
	//sort.Slice(sortNums2, func(i, j int) bool {
	//	return sortNums2[i][1] < sortNums2[j][1]
	//})

	// time limit exceeded
	//sortNums2 := make([][2]int, 0)
	//for i, v := range nums2 {
	//	pos := sort.Search(len(sortNums2), func(i int) bool {
	//		return sortNums2[i][1] > v
	//	})
	//
	//	sortNums2 = append(sortNums2, [2]int{})
	//	copy(sortNums2[pos+1:], sortNums2[pos:])
	//	sortNums2[pos] = [2]int{i, v}
	//}

	//num2Cur := len(nums2) - 1
	//left, right := 0, len(nums1)-1
	//
	//for num2Cur >= 0 {
	//	if nums1[right] > sortNums2[num2Cur][1] {
	//		rlt[sortNums2[num2Cur][0]] = nums1[right]
	//		right--
	//	} else {
	//		rlt[sortNums2[num2Cur][0]] = nums1[left]
	//		left++
	//	}
	//	num2Cur--
	//}

	// heapify
	sortNums2 := make(SliceHeap, len(nums2))
	for i, v := range nums2 {
		sortNums2[i] = []int{i, v}
	}
	heap.Init(&sortNums2)

	left, right := 0, len(nums1)-1

	for sortNums2.Len() > 0 {
		num2 := heap.Pop(&sortNums2).([]int)
		if nums1[right] > num2[1] {
			rlt[num2[0]] = nums1[right]
			right--
		} else {
			rlt[num2[0]] = nums1[left]
			left++
		}
	}
	return rlt
}

type SliceHeap [][]int

func (s SliceHeap) Len() int {
	return len(s)
}

func (s SliceHeap) Less(i, j int) bool {
	return s[i][1] > s[j][1]
}

func (s SliceHeap) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s *SliceHeap) Push(x any) {
	*s = append(*s, x.([]int))
}

func (s *SliceHeap) Pop() any {
	rlt := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return rlt
}

//leetcode submit region end(Prohibit modification and deletion)
