/**
Given the root of a binary tree, invert the tree, and return its root.


 Example 1:


Input: root = [4,2,7,1,3,6,9]
Output: [4,7,2,9,6,3,1]


 Example 2:


Input: root = [2,1,3]
Output: [2,3,1]


 Example 3:


Input: root = []
Output: []



 Constraints:


 The number of nodes in the tree is in the range [0, 100].
 -100 <= Node.val <= 100


 Related Topics 树 深度优先搜索 广度优先搜索 二叉树 👍 1785 👎 0

*/

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
package main

//	func invertTree(root *TreeNode) *TreeNode {
//		if root == nil {
//			return nil
//		}
//
//		root.Left, root.Right = invertTree(root.Right), invertTree(root.Left)
//		return root
//	}
//
// 主函数
func invertTree(root *TreeNode) *TreeNode {
	// 遍历二叉树，交换每个节点的子节点
	traverse(root)
	return root
}

// 二叉树遍历函数
func traverse(root *TreeNode) {
	if root == nil {
		return
	}

	/**** 中序遍历 ****/
	// 每一个节点需要做的事就是交换它的左右子节点

	// 遍历框架，去遍历左右子树的节点
	traverse(root.Left)
	tmp := root.Left
	root.Left = root.Right
	root.Right = tmp
	traverse(root.Left)
}

//leetcode submit region end(Prohibit modification and deletion)
