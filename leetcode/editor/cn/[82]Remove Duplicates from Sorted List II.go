/**
Given the head of a sorted linked list, delete all nodes that have duplicate
numbers, leaving only distinct numbers from the original list. Return the linked
list sorted as well.


 Example 1:


Input: head = [1,2,3,3,4,4,5]
Output: [1,2,5]


 Example 2:


Input: head = [1,1,1,2,3]
Output: [2,3]



 Constraints:


 The number of nodes in the list is in the range [0, 300].
 -100 <= Node.val <= 100
 The list is guaranteed to be sorted in ascending order.


 Related Topics 链表 双指针 👍 1295 👎 0

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

func deleteDuplicates(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}
	pre, cur := dummy, head
	for cur != nil && cur.Next != nil {
		if pre.Next.Val == cur.Next.Val {
			cur = cur.Next
			continue
		}

		if pre.Next == cur {
			cur = cur.Next
			pre = pre.Next
		} else {
			cur = cur.Next
			pre.Next = cur
		}
	}

	if pre.Next != cur {
		cur = cur.Next
		pre.Next = cur
	}
	return dummy.Next
}

//func deleteDuplicates(head *ListNode) *ListNode {
//	dummy := &ListNode{Next: head}
//	left, right := dummy, head
//
//	for right != nil {
//		flag := 0
//		for right.Next != nil && right.Next.Val == right.Val {
//			right = right.Next
//			flag = 1
//		}
//
//		if flag == 0 {
//			left.Next = right
//			left = left.Next
//		}
//		right = right.Next
//	}
//
//	left.Next = nil
//	return dummy.Next
//}

//leetcode submit region end(Prohibit modification and deletion)
