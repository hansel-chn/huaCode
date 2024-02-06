/**
给定一个单链表 L 的头节点 head ，单链表 L 表示为：


L0 → L1 → … → Ln - 1 → Ln


 请将其重新排列后变为：


L0 → Ln → L1 → Ln - 1 → L2 → Ln - 2 → …

 不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。



 示例 1：




输入：head = [1,2,3,4]
输出：[1,4,2,3]

 示例 2：




输入：head = [1,2,3,4,5]
输出：[1,5,2,4,3]



 提示：


 链表的长度范围为 [1, 5 * 10⁴]
 1 <= node.val <= 1000


 Related Topics 栈 递归 链表 双指针 👍 1390 👎 0

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

//type ListNode struct {
//	Val  int
//	Next *ListNode
//}

type stack143 struct {
	buf   []*ListNode
	count int
}

func (s *stack143) pop() *ListNode {
	var node *ListNode
	if s.count > 0 {
		node = s.buf[len(s.buf)-1]
		s.buf = s.buf[:len(s.buf)-1]
		s.count--
	}
	return node
}

func (s *stack143) push(node *ListNode) {
	s.buf = append(s.buf, node)
	s.count++
}

//type queue143 struct {
//	buf   []*ListNode
//	count int
//}
//
//func (q *queue143) add(node *ListNode) {
//	q.buf = append(q.buf, node)
//	q.count++
//}
//
//func (q *queue143) remove() *ListNode {
//	var node *ListNode
//	if q.count > 0 {
//		node = q.buf[0]
//		q.buf = q.buf[1:]
//		q.count--
//	}
//	return node
//}

func reorderList(head *ListNode) {
	preHead := &ListNode{
		Next: head,
	}
	stack := stack143{}
	count := 0
	for head != nil {
		stack.push(head)
		count++
		head = head.Next
	}

	head = preHead.Next
	for i := 1; i <= (count-1)/2; i++ {
		temp := head.Next
		head.Next = stack.pop()
		head.Next.Next = temp
		head = head.Next.Next
	}
	if count%2 == 0 {
		head.Next.Next = nil

	} else {
		head.Next = nil

	}
}

//leetcode submit region end(Prohibit modification and deletion)
