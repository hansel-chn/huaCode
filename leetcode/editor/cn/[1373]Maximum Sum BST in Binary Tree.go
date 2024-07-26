/**
Given a binary tree root, return the maximum sum of all keys of any sub-tree
which is also a Binary Search Tree (BST).

 Assume a BST is defined as follows:


 The left subtree of a node contains only nodes with keys less than the node's
key.
 The right subtree of a node contains only nodes with keys greater than the
node's key.
 Both the left and right subtrees must also be binary search trees.



 Example 1:




Input: root = [1,4,3,2,4,2,5,null,null,null,null,null,null,4,6]
Output: 20
Explanation: Maximum sum in a valid Binary search tree is obtained in root node
with key equal to 3.


 Example 2:




Input: root = [4,3,null,1,2]
Output: 2
Explanation: Maximum sum in a valid Binary search tree is obtained in a single
root node with key equal to 2.


 Example 3:


Input: root = [-4,-2,-5]
Output: 0
Explanation: All values are negatives. Return an empty BST.



 Constraints:


 The number of nodes in the tree is in the range [1, 4 * 10⁴].
 -4 * 10⁴ <= Node.val <= 4 * 10⁴


 Related Topics 树 深度优先搜索 二叉搜索树 动态规划 二叉树 👍 202 👎 0

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

func maxSumBST(root *TreeNode) int {
	var maxSum int
	var postTraverse func(node *TreeNode) (maxV, minV, sum int, isBST bool)
	postTraverse = func(node *TreeNode) (maxV, minV, sum int, isBST bool) {
		if node == nil {
			return int(4e5), int(-4e5), 0, true
		}

		lMin, lMax, lSum, lIsBST := postTraverse(node.Left)
		rMin, rMax, rSum, rIsBST := postTraverse(node.Right)

		if !(lIsBST && rIsBST) || lMax >= node.Val || rMin <= node.Val {
			return 0, 0, 0, false
		}

		maxSum = max(maxSum, lSum+rSum+node.Val)
		return min(node.Val, lMin), max(rMax, node.Val), lSum + rSum + node.Val, true
	}
	postTraverse(root)
	return maxSum
}

//leetcode submit region end(Prohibit modification and deletion)
