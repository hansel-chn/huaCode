/**
Given two integer arrays preorder and inorder where preorder is the preorder
traversal of a binary tree and inorder is the inorder traversal of the same tree,
construct and return the binary tree.


 Example 1:


Input: preorder = [3,9,20,15,7], inorder = [9,3,15,20,7]
Output: [3,9,20,null,null,15,7]


 Example 2:


Input: preorder = [-1], inorder = [-1]
Output: [-1]



 Constraints:


 1 <= preorder.length <= 3000
 inorder.length == preorder.length
 -3000 <= preorder[i], inorder[i] <= 3000
 preorder and inorder consist of unique values.
 Each value of inorder also appears in preorder.
 preorder is guaranteed to be the preorder traversal of the tree.
 inorder is guaranteed to be the inorder traversal of the tree.


 Related Topics 树 数组 哈希表 分治 二叉树 👍 2247 👎 0

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

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) < 1 {
		return nil
	}

	root := &TreeNode{}
	root.Val = preorder[0]

	var idx, v int
	for idx, v = range inorder {
		if root.Val == v {
			break
		}
	}

	root.Left = buildTree(preorder[1:idx+1], inorder[:idx])
	root.Right = buildTree(preorder[idx+1:], inorder[idx+1:])
	return root
}

//leetcode submit region end(Prohibit modification and deletion)
