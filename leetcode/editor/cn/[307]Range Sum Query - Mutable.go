//Given an integer array nums, handle multiple queries of the following types:
//
//
// Update the value of an element in nums.
// Calculate the sum of the elements of nums between indices left and right
//inclusive where left <= right.
//
//
// Implement the NumArray class:
//
//
// NumArray(int[] nums) Initializes the object with the integer array nums.
// void update(int index, int val) Updates the value of nums[index] to be val.
// int sumRange(int left, int right) Returns the sum of the elements of nums
//between indices left and right inclusive (i.e. nums[left] + nums[left + 1] + ... +
//nums[right]).
//
//
//
// Example 1:
//
//
//Input
//["NumArray", "sumRange", "update", "sumRange"]
//[[[1, 3, 5]], [0, 2], [1, 2], [0, 2]]
//Output
//[null, 9, null, 8]
//
//Explanation
//NumArray numArray = new NumArray([1, 3, 5]);
//numArray.sumRange(0, 2); // return 1 + 3 + 5 = 9
//numArray.update(1, 2);   // nums = [1, 2, 5]
//numArray.sumRange(0, 2); // return 1 + 2 + 5 = 8
//
//
//
// Constraints:
//
//
// 1 <= nums.length <= 3 * 10⁴
// -100 <= nums[i] <= 100
// 0 <= index < nums.length
// -100 <= val <= 100
// 0 <= left <= right < nums.length
// At most 3 * 10⁴ calls will be made to update and sumRange.
//
//
// Related Topics 设计 树状数组 线段树 数组 👍 761 👎 0

// leetcode submit region begin(Prohibit modification and deletion)

package main

type SegmentNode struct {
	lNode, rNode *SegmentNode
	l, r         int
	v            int
}

func (s *SegmentNode) Update(index int, val int, mergeFunc func(a, b int) int) bool {
	if index < s.l || index > s.r {
		return false
	}

	if s.l == s.r && index == s.l {
		s.v = val
		return true
	}

	if s.lNode.Update(index, val, mergeFunc) || s.rNode.Update(index, val, mergeFunc) {
		s.v = mergeFunc(s.lNode.v, s.rNode.v)
		return true
	}

	return false
}

//func (s *SegmentNode) Query(l int, r int, mergeFunc func(a, b int) int) int {
//	if l == s.l && r == s.r {
//		return s.v
//	}
//
//	mid := s.l + (s.r-s.l)/2
//	if r <= mid {
//		return s.lNode.Query(l, r, mergeFunc)
//	} else if l > mid {
//		return s.rNode.Query(l, r, mergeFunc)
//	} else {
//		return mergeFunc(s.lNode.Query(l, mid, mergeFunc), s.rNode.Query(mid+1, r, mergeFunc))
//	}
//}

//func (s *SegmentNode) Query(l int, r int, mergeFunc func(a, b int) int) int {
//	if s.l > r || s.r < l {
//		return 0
//	}
//
//	if l <= s.l && r >= s.r {
//		return s.v
//	}
//
//	return mergeFunc(s.lNode.Query(l, r, mergeFunc), s.rNode.Query(l, r, mergeFunc))
//}

func query(s *SegmentNode, l int, r int, mergeFunc func(a, b int) int) int {
	if s.l > r || s.r < l {
		return 0
	}

	if l <= s.l && r >= s.r {
		return s.v
	}

	return mergeFunc(query(s.lNode, l, r, mergeFunc), query(s.rNode, l, r, mergeFunc))
}

func update(s *SegmentNode, index int, val int, mergeFunc func(a, b int) int) bool {
	if index < s.l || index > s.r {
		return false
	}

	if s.l == s.r && index == s.l {
		s.v = val
		return true
	}

	if s.lNode.Update(index, val, mergeFunc) || s.rNode.Update(index, val, mergeFunc) {
		s.v = mergeFunc(s.lNode.v, s.rNode.v)
		return true
	}

	return false
}

func buildNode(nums []int, l, r int, mergeFunc func(a, b int) int) *SegmentNode {
	if l == r {
		return &SegmentNode{
			l: l,
			r: r,
			v: nums[l],
		}
	}

	mid := l + (r-l)/2
	lNode := buildNode(nums, l, mid, mergeFunc)
	rNode := buildNode(nums, mid+1, r, mergeFunc)

	return &SegmentNode{
		lNode: lNode,
		rNode: rNode,
		l:     l,
		r:     r,
		v:     mergeFunc(lNode.v, rNode.v),
	}
}

type NumArray struct {
	root      *SegmentNode
	MergeFunc func(a, b int) int
}

func Constructor(nums []int) NumArray {
	segTree := NumArray{
		MergeFunc: func(a, b int) int {
			return a + b
		},
	}

	segTree.root = buildNode(nums, 0, len(nums)-1, segTree.MergeFunc)
	return segTree
}

func (n *NumArray) Update(index int, val int) {
	update(n.root, index, val,n.MergeFunc)
	//n.root.Update(index, val, n.MergeFunc)
}

func (n *NumArray) SumRange(left int, right int) int {
	return query(n.root, left, right, n.MergeFunc)
	//return n.root.Query(left, right, n.MergeFunc)
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * obj.Update(index,val);
 * param_2 := obj.SumRange(left,right);
 */
//leetcode submit region end(Prohibit modification and deletion)
