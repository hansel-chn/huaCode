/**
Given an integer array nums and an integer k, return the kᵗʰ largest element in
the array.

 Note that it is the kᵗʰ largest element in the sorted order, not the kᵗʰ
distinct element.

 Can you solve it without sorting?


 Example 1:
 Input: nums = [3,2,1,5,6,4], k = 2
Output: 5

 Example 2:
 Input: nums = [3,2,3,1,2,4,5,5,6], k = 4
Output: 4


 Constraints:


 1 <= k <= nums.length <= 10⁵
 -10⁴ <= nums[i] <= 10⁴


 Related Topics 数组 分治 快速选择 排序 堆（优先队列） 👍 2478 👎 0

*/

// leetcode submit region begin(Prohibit modification and deletion)
package main

import (
	"container/heap"
)

func findKthLargest(nums []int, k int) int {
	pq := findKthLargestHeap(nums)
	heap.Init(&pq)
	for i := 0; i < k-1; i++ {
		heap.Pop(&pq)
	}
	//fmt.Println(reflect.TypeOf(heap.Pop(&pq)).Kind())
	return heap.Pop(&pq).(int)
}

type findKthLargestHeap []int

func (f findKthLargestHeap) Less(i, j int) bool {
	return f[i] > f[j]
}

func (f findKthLargestHeap) Len() int {
	return len(f)
}

func (f findKthLargestHeap) Swap(i, j int) {
	f[i], f[j] = f[j], f[i]
}

func (f *findKthLargestHeap) Push(v any) {
	*f = append(*f, v.(int))
}

func (f *findKthLargestHeap) Pop() any {
	old := *f
	length := old.Len()
	v := old[length-1]
	*f = (*f)[:length-1]
	return v
}

// // build heap by myself
//func findKthLargest(nums []int, k int) int {
//	pq := buildFindKthLargestHeap()
//	for _, num := range nums {
//		pq.Push(num)
//	}
//
//	//fmt.Println(pq)
//	for i := 0; i < k-1; i++ {
//		pq.Pop()
//		//fmt.Println(pq.Pop())
//	}
//
//	return pq.Pop()
//}
//
//func buildFindKthLargestHeap() findKthLargestHeap {
//	return make(findKthLargestHeap, 1)
//}
//
//type findKthLargestHeap []int
//
//func (f *findKthLargestHeap) IfEmpty() bool {
//	return !(len(*f) <= 1)
//}
//
//func (f *findKthLargestHeap) Pop() int {
//	maxV := (*f)[1]
//	f.swap(1, f.maxIdx())
//	*f = (*f)[:len(*f)-1]
//	f.sink(1)
//	return maxV
//}
//
//func (f *findKthLargestHeap) Push(v int) {
//	*f = append(*f, v)
//	f.swim(f.maxIdx())
//}
//
//func (f *findKthLargestHeap) swim(i int) {
//	parent := f.Parent(i)
//	if parent > 0 {
//		if (*f)[i] > (*f)[parent] {
//			f.swap(i, parent)
//			f.swim(parent)
//		}
//	}
//}
//
//func (f *findKthLargestHeap) sink(i int) {
//	left, right := f.Child(i)
//	if left > f.maxIdx() {
//		return
//	}
//
//	hi := left
//	if right < f.maxIdx() && (*f)[hi] < (*f)[right] {
//		hi = right
//	}
//
//	if (*f)[i] < (*f)[hi] {
//		f.swap(hi, i)
//		f.sink(hi)
//	}
//}
//
//func (f findKthLargestHeap) maxIdx() int {
//	return len(f) - 1
//}
//
//func (f *findKthLargestHeap) Parent(idx int) int {
//	return idx / 2
//}
//
//func (f *findKthLargestHeap) Child(idx int) (left, right int) {
//	return idx * 2, idx*2 + 1
//}
//
//func (f *findKthLargestHeap) swap(i, j int) {
//	(*f)[i], (*f)[j] = (*f)[j], (*f)[i]
//}

//func findKthLargest(nums []int, k int) int {
//	return quickSort(nums, 0, len(nums)-1, len(nums)-k)
//}
//
//func quickSort(nums []int, lo, hi, k int) int {
//	pivot := lo
//	i := lo + 1
//	j := hi
//	for i <= j {
//		// for i <= j && nums[i] < nums[pivot] {
//		// situation2: for i < hi && nums[i] < nums[pivot] {
//		// situation1: for i <= hi && nums[i] < nums[pivot] {
//		for i <= hi && nums[i] < nums[pivot] {
//			i++
//		}
//		//for i <= j && nums[j] > nums[pivot] {
//		for lo < j && nums[j] > nums[pivot] {
//			j--
//		}
//		// 大于或者大于等于都可以，因为上面的条件限制了j位置, i <= j失效说明当前位置j一定小于pivot
//		// situation1, 2: if i >= j {
//		// situation1: if i > j {
//		if i > j {
//			break
//		}
//		findKthLargestSwap(nums, i, j)
//		i++
//		j--
//	}
//
//	findKthLargestSwap(nums, pivot, j)
//
//	if j == k {
//		return nums[j]
//	} else if j < k {
//		return quickSort(nums, j+1, hi, k)
//	} else {
//		return quickSort(nums, lo, j-1, k)
//	}
//}
//func findKthLargestSwap(nums []int, i, j int) {
//	nums[i], nums[j] = nums[j], nums[i]
//}
//
//func findKthLargestShuffle(nums []int) {
//	for i := 0; i < len(nums); i++ {
//		findKthLargestSwap(nums, i, i+rand.Intn(len(nums)-i))
//	}
//}

//leetcode submit region end(Prohibit modification and deletion)
