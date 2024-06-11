/**
You are given two integer arrays nums1 and nums2 sorted in non-decreasing order
and an integer k.

 Define a pair (u, v) which consists of one element from the first array and
one element from the second array.

 Return the k pairs (u1, v1), (u2, v2), ..., (uk, vk) with the smallest sums.


 Example 1:


Input: nums1 = [1,7,11], nums2 = [2,4,6], k = 3
Output: [[1,2],[1,4],[1,6]]
Explanation: The first 3 pairs are returned from the sequence: [1,2],[1,4],[1,6]
,[7,2],[7,4],[11,2],[7,6],[11,4],[11,6]


 Example 2:


Input: nums1 = [1,1,2], nums2 = [1,2,3], k = 2
Output: [[1,1],[1,1]]
Explanation: The first 2 pairs are returned from the sequence: [1,1],[1,1],[1,2]
,[2,1],[1,2],[2,2],[1,3],[1,3],[2,3]



 Constraints:


 1 <= nums1.length, nums2.length <= 10⁵
 -10⁹ <= nums1[i], nums2[i] <= 10⁹
 nums1 and nums2 both are sorted in non-decreasing order.
 1 <= k <= 10⁴
 k <= nums1.length * nums2.length


 Related Topics 数组 堆（优先队列） 👍 595 👎 0

*/

// leetcode submit region begin(Prohibit modification and deletion)
package main

//func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
//	priorQueue := heapNodes{
//		num1:  nums1,
//		num2:  nums2,
//		items: make([]item, 0),
//	}
//	for i := range nums1 {
//		priorQueue.items = append(priorQueue.items, item{
//			idx1: i,
//			idx2: 0,
//		})
//	}
//
//	heap.Init(&priorQueue)
//
//	rlt := make([][]int, 0)
//	for i := 0; i < k; i++ {
//		minV := heap.Pop(&priorQueue).(item)
//		rlt = append(rlt, []int{nums1[minV.idx1], nums2[minV.idx2]})
//		if minV.idx2+1 < len(nums2) {
//			heap.Push(&priorQueue, item{
//				idx1: minV.idx1,
//				idx2: minV.idx2 + 1,
//			})
//		}
//	}
//	return rlt
//}
//
//type heapNodes struct {
//	num1  []int
//	num2  []int
//	items []item
//}
//
//type item struct {
//	idx1 int
//	idx2 int
//}
//
//func (h heapNodes) Len() int {
//	return len(h.items)
//}
//
//func (h heapNodes) Less(i, j int) bool {
//	return h.num1[h.items[i].idx1]+h.num2[h.items[i].idx2] < h.num1[h.items[j].idx1]+h.num2[h.items[j].idx2]
//}
//
//func (h heapNodes) Swap(i, j int) {
//	h.items[i], h.items[j] = h.items[j], h.items[i]
//}
//
//func (h *heapNodes) Push(x any) {
//	h.items = append(h.items, x.(item))
//}
//
//func (h *heapNodes) Pop() any {
//	old := h.items
//	length := len(old)
//	pop := old[length-1]
//	h.items = old[:length-1]
//	return pop
//}

//leetcode submit region end(Prohibit modification and deletion)
