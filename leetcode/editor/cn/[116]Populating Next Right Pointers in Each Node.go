/**
You are given a perfect binary tree where all leaves are on the same level, and
every parent has two children. The binary tree has the following definition:


struct Node {
  int val;
  Node *left;
  Node *right;
  Node *next;
}


 Populate each next pointer to point to its next right node. If there is no
next right node, the next pointer should be set to NULL.

 Initially, all next pointers are set to NULL.


 Example 1:


Input: root = [1,2,3,4,5,6,7]
Output: [1,#,2,3,#,4,5,6,7,#]
Explanation: Given the above perfect binary tree (Figure A), your function
should populate each next pointer to point to its next right node, just like in
Figure B. The serialized output is in level order as connected by the next pointers,
with '#' signifying the end of each level.


 Example 2:


Input: root = []
Output: []



 Constraints:


 The number of nodes in the tree is in the range [0, 2¹² - 1].
 -1000 <= Node.val <= 1000



 Follow-up:


 You may only use constant extra space.
 The recursive approach is fine. You may assume implicit stack space does not
count as extra space for this problem.


 Related Topics 树 深度优先搜索 广度优先搜索 链表 二叉树 👍 1111 👎 0

*/

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

package main

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}

	var traverse func(node1, node2 *Node)
	traverse = func(node1, node2 *Node) {
		if node1 == nil {
			return
		}

		node1.Next = node2
		traverse(node1.Left, node1.Right)
		traverse(node1.Right, node2.Left)
		traverse(node2.Left, node2.Right)
	}

	traverse(root.Left, root.Right)
	return root
}

//func connect(root *Node) *Node {
//	if root == nil {
//		return nil
//	}
//	queue := make([]*Node, 0)
//	var bfs func()
//	bfs = func() {
//		for len(queue) != 0 {
//			sz := len(queue)
//			for i := 0; i < sz; i++ {
//				node := queue[0]
//				queue = queue[1:]
//				if i < sz-1 {
//					node.Next = queue[0]
//				}
//				if node.Left != nil && node.Right != nil {
//					queue = append(queue, node.Left, node.Right)
//				}
//			}
//		}
//	}
//
//	queue = append(queue, root)
//	bfs()
//	return root
//}
//leetcode submit region end(Prohibit modification and deletion)
