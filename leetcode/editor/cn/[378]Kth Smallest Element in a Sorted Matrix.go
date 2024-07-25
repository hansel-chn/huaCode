/**
Given an n x n matrix where each of the rows and columns is sorted in ascending
order, return the kᵗʰ smallest element in the matrix.

 Note that it is the kᵗʰ smallest element in the sorted order, not the kᵗʰ
distinct element.

 You must find a solution with a memory complexity better than O(n²).


 Example 1:


Input: matrix = [[1,5,9],[10,11,13],[12,13,15]], k = 8
Output: 13
Explanation: The elements in the matrix are [1,5,9,10,11,12,13,13,15], and the 8
ᵗʰ smallest number is 13


 Example 2:


Input: matrix = [[-5]], k = 1
Output: -5



 Constraints:


 n == matrix.length == matrix[i].length
 1 <= n <= 300
 -10⁹ <= matrix[i][j] <= 10⁹
 All the rows and columns of matrix are guaranteed to be sorted in non-
decreasing order.
 1 <= k <= n²



 Follow up:


 Could you solve the problem with a constant memory (i.e., O(1) memory
complexity)?
 Could you solve the problem in O(n) time complexity? The solution may be too
advanced for an interview but you may find reading this paper fun.


 Related Topics 数组 二分查找 矩阵 排序 堆（优先队列） 👍 1059 👎 0

*/

// leetcode submit region begin(Prohibit modification and deletion)
package main

import (
	"container/heap"
)

func kthSmallest(matrix [][]int, k int) int {
	pq := priorityQueue378{}
	for i := 0; i < len(matrix); i++ {
		heap.Push(&pq, [3]int{matrix[i][0], i, 0})
	}

	for i := 0; i < k-1; i++ {
		p := heap.Pop(&pq).([3]int)
		nxtX := p[1]
		nxtY := p[2] + 1
		if nxtY < len(matrix[0]) {
			heap.Push(&pq, [3]int{matrix[nxtX][nxtY], nxtX, nxtY})
		}
	}

	return heap.Pop(&pq).([3]int)[0]
}

type priorityQueue378 [][3]int

func (pq priorityQueue378) Less(i, j int) bool {
	return pq[i][0] <= pq[j][0]
}

func (pq priorityQueue378) Len() int {
	return len(pq)
}

func (pq *priorityQueue378) Swap(i, j int) {
	(*pq)[i], (*pq)[j] = (*pq)[j], (*pq)[i]
}

func (pq *priorityQueue378) Push(i any) {
	*pq = append(*pq, i.([3]int))
}

func (pq *priorityQueue378) Pop() any {
	l := pq.Len()
	v := (*pq)[l-1]
	*pq = (*pq)[:l-1]
	return v
}

//leetcode submit region end(Prohibit modification and deletion)
