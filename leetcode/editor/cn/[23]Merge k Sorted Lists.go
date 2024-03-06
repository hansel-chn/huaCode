/**
You are given an array of k linked-lists lists, each linked-list is sorted in
ascending order.

 Merge all the linked-lists into one sorted linked-list and return it.


 Example 1:


Input: lists = [[1,4,5],[1,3,4],[2,6]]
Output: [1,1,2,3,4,4,5,6]
Explanation: The linked-lists are:
[
  1->4->5,
  1->3->4,
  2->6
]
merging them into one sorted list:
1->1->2->3->4->4->5->6


 Example 2:


Input: lists = []
Output: []


 Example 3:


Input: lists = [[]]
Output: []



 Constraints:


 k == lists.length
 0 <= k <= 10⁴
 0 <= lists[i].length <= 500
 -10⁴ <= lists[i][j] <= 10⁴
 lists[i] is sorted in ascending order.
 The sum of lists[i].length will not exceed 10⁴.


 Related Topics 链表 分治 堆（优先队列） 归并排序 👍 2768 👎 0

*/

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
package main

//import "fmt"
//
//type ListNode struct {
//	Val  int
//	Next *ListNode
//}
//
//func main() {
//	node1 := ListNode{Val: 1}
//	node1.Next = &ListNode{Val: 4}
//	node1.Next.Next = &ListNode{Val: 5}
//	node2 := ListNode{Val: 1}
//	node2.Next = &ListNode{Val: 3}
//	node2.Next.Next = &ListNode{Val: 4}
//	node3 := ListNode{Val: 2}
//	node3.Next = &ListNode{Val: 6}
//	testCase := []*ListNode{&node1, &node2, &node3}
//
//	fmt.Println(mergeKLists(testCase))
//}

func mergeKLists(lists []*ListNode) *ListNode {
	queue := make(priorityQueue, 1)
	for _, list := range lists {
		if list != nil {
			queue = append(queue, list)
		}
	}

	if len(queue) == 0 {
		return nil
	}

	queue.maxHeapify(1)

	dummy := &ListNode{
		Val:  0,
		Next: nil,
	}
	node := dummy

	for len(queue) > 1 {
		rm := queue.removeMin()
		node.Next = rm
		node = node.Next
		if rm.Next != nil {
			queue.insert(rm.Next)
		}
	}
	return dummy.Next
}

type priorityQueue []*ListNode

func (p priorityQueue) maxHeapify(i int) {
	if i >= len(p) {
		return
	}

	p.maxHeapify(getLeftChildIdx(i))
	p.maxHeapify(getRightChildIdx(i))
	p.sink(i)
	return
}

func (p *priorityQueue) removeMin() *ListNode {
	delNode := (*p)[1]
	p.swap(1, len(*p)-1)
	*p = (*p)[:len(*p)-1]

	if len(*p) > 1 {
		p.sink(1)
	}
	return delNode
}

func (p *priorityQueue) insert(node *ListNode) {
	*p = append(*p, node)

	p.swim(len(*p) - 1)
}

func (p priorityQueue) sink(i int) {
	for getLeftChildIdx(i) < len(p) {
		minIdx := i
		if p.less(getLeftChildIdx(i), i) {
			minIdx = getLeftChildIdx(i)
		}

		if getRightChildIdx(i) < len(p) && p.less(getRightChildIdx(i), minIdx) {
			minIdx = getRightChildIdx(i)
		}

		if minIdx == i {
			break
		}

		p.swap(minIdx, i)
		i = minIdx
	}
}

func (p priorityQueue) swim(i int) {
	for i > 1 && p.less(i, getParentIdx(i)) {
		p.swap(getParentIdx(i), i)
		i = getParentIdx(i)
	}
}

func (p priorityQueue) less(i, j int) bool {
	if p[i].Val <= p[j].Val {
		return true
	} else {
		return false
	}
}

func (p priorityQueue) swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func getParentIdx(i int) int {
	return i / 2
}

func getLeftChildIdx(i int) int {
	return i * 2
}

func getRightChildIdx(i int) int {
	return i*2 + 1
}

//leetcode submit region end(Prohibit modification and deletion)
