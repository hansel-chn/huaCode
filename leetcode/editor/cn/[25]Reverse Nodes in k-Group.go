/**
Given the head of a linked list, reverse the nodes of the list k at a time, and
return the modified list.

 k is a positive integer and is less than or equal to the length of the linked
list. If the number of nodes is not a multiple of k then left-out nodes, in the
end, should remain as it is.

 You may not alter the values in the list's nodes, only nodes themselves may be
changed.


 Example 1:


Input: head = [1,2,3,4,5], k = 2
Output: [2,1,4,3,5]


 Example 2:


Input: head = [1,2,3,4,5], k = 3
Output: [3,2,1,4,5]



 Constraints:


 The number of nodes in the list is n.
 1 <= k <= n <= 5000
 0 <= Node.val <= 1000



 Follow-up: Can you solve the problem in O(1) extra memory space?

 Related Topics 递归 链表 👍 2309 👎 0

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

type ListNode struct {
	Val  int
	Next *ListNode
}

//// iterate method
//func reverseKGroup(head *ListNode, k int) *ListNode {
//
//	oldHead := head
//	for i := 0; i < k; i++ {
//		if head == nil {
//			return oldHead
//		}
//		head = head.Next
//	}
//
//	newHead := reverse(oldHead, k)
//	oldHead.Next = reverseKGroup(head, k)
//	return newHead
//}
//func reverse(node *ListNode, k int) (newHead *ListNode) {
//	if k == 1 {
//		return node
//	}
//
//	newHead = reverse(node.Next, k-1)
//	node.Next.Next = node
//	node.Next = nil
//	return newHead
//}

// 不对，应该控制两个节点之间，之前reverse未控制，导致反转到应该暂停的地方继续向后反转了
func reverseKGroup(head *ListNode, k int) *ListNode {
	var oldHead *ListNode
	oldHead = head
	for i := 0; i < k; i++ {
		if head == nil {
			return oldHead
		}
		head = head.Next
	}

	newHead := reverse(oldHead, head)
	oldHead.Next = reverseKGroup(head, k)
	return newHead
}
func reverse(head *ListNode, tail *ListNode) (newHead *ListNode) {
	var cur, next *ListNode
	cur = head
	next = head.Next

	for next != tail {
		newNext := next.Next
		next.Next = cur
		cur = next
		next = newNext
	}
	return cur
}

//// https://labuladong.online/algo/data-structure/reverse-nodes-in-k-group/#%E4%B8%80%E3%80%81%E5%88%86%E6%9E%90%E9%97%AE%E9%A2%98
//// 区别是，对当前节点cur的认识，我写的方法当前节点cur为链的新头部，对应链接中的pre节点；而链接中的cur节点为即将加入新链的节点，pre为新链的头节点
//func reverse(head *ListNode, tail *ListNode) (newHead *ListNode) {
//	var prev, cur, next *ListNode
//	prev = nil
//	cur = head
//	next = head.Next
//
//	head.Next = nil
//	for next != tail {
//		prev = cur
//		cur = next
//		next = next.Next
//		cur.Next = prev
//	}
//	return cur
//}

// 这样反转链表连同最后长度未达到k的也反转了，与题意不符合
//func reverseKGroup(head *ListNode, k int) *ListNode {
//	if head == nil {
//		return nil
//	}
//
//	newTail := head
//	var prev, cur, next *ListNode
//	prev = nil
//	cur = head
//	next = head.Next
//
//	for i := 1; i < k; i++ {
//		if next==nil{
//			newTail.Next = nil
//			return cur
//		}
//		prev = cur
//		cur = next
//		next = next.Next
//		cur.Next = prev
//	}
//	newTail.Next = reverseKGroup(next, k)
//	return cur
//}

//leetcode submit region end(Prohibit modification and deletion)
