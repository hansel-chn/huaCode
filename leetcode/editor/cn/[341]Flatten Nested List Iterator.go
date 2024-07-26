/**
You are given a nested list of integers nestedList. Each element is either an
integer or a list whose elements may also be integers or other lists. Implement
an iterator to flatten it.

 Implement the NestedIterator class:


 NestedIterator(List<NestedInteger> nestedList) Initializes the iterator with
the nested list nestedList.
 int next() Returns the next integer in the nested list.
 boolean hasNext() Returns true if there are still some integers in the nested
list and false otherwise.


 Your code will be tested with the following pseudocode:


initialize iterator with nestedList
res = []
while iterator.hasNext()
    append iterator.next() to the end of res
return res


 If res matches the expected flattened list, then your code will be judged as
correct.


 Example 1:


Input: nestedList = [[1,1],2,[1,1]]
Output: [1,1,2,1,1]
Explanation: By calling next repeatedly until hasNext returns false, the order
of elements returned by next should be: [1,1,2,1,1].


 Example 2:


Input: nestedList = [1,[4,[6]]]
Output: [1,4,6]
Explanation: By calling next repeatedly until hasNext returns false, the order
of elements returned by next should be: [1,4,6].



 Constraints:


 1 <= nestedList.length <= 500
 The values of the integers in the nested list is in the range [-10⁶, 10⁶].


 Related Topics 栈 树 深度优先搜索 设计 队列 迭代器 👍 548 👎 0

*/

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * // This is the interface that allows for creating nested lists.
 * // You should not implement it, or speculate about its implementation
 * type NestedInteger struct {
 * }
 *
 * // Return true if this NestedInteger holds a single integer, rather than a nested list.
 * func (this NestedInteger) IsInteger() bool {}
 *
 * // Return the single integer that this NestedInteger holds, if it holds a single integer
 * // The result is undefined if this NestedInteger holds a nested list
 * // So before calling this method, you should have a check
 * func (this NestedInteger) GetInteger() int {}
 *
 * // Set this NestedInteger to hold a single integer.
 * func (n *NestedInteger) SetInteger(value int) {}
 *
 * // Set this NestedInteger to hold a nested list and adds a nested integer to it.
 * func (this *NestedInteger) Add(elem NestedInteger) {}
 *
 * // Return the nested list that this NestedInteger holds, if it holds a nested list
 * // The list length is zero if this NestedInteger holds a single integer
 * // You can access NestedInteger's List element directly if you want to modify it
 * func (this NestedInteger) GetList() []*NestedInteger {}
 */
package main

//type NestedInteger struct {
//}
//
//func (this NestedInteger) IsInteger() bool           {}
//func (this NestedInteger) GetInteger() int           {}
//func (this NestedInteger) GetList() []*NestedInteger {}

//type NestedIterator struct {
//	stack [][]*NestedInteger
//}
//
//func Constructor(nestedList []*NestedInteger) *NestedIterator {
//	return &NestedIterator{stack: [][]*NestedInteger{nestedList}}
//}
//
//func (this *NestedIterator) Next() int {
//	nxt := this.stack[len(this.stack)-1][0].GetInteger()
//	this.stack[len(this.stack)-1] = this.stack[len(this.stack)-1][1:]
//	return nxt
//}
//
//func (this *NestedIterator) HasNext() bool {
//	for len(this.stack) > 0 {
//		nestedList := this.stack[len(this.stack)-1]
//
//		if len(nestedList) == 0 {
//			this.stack = this.stack[:len(this.stack)-1]
//			continue
//		}
//
//		if nestedList[0].IsInteger() {
//			return true
//		}
//		this.stack[len(this.stack)-1] = this.stack[len(this.stack)-1][1:]
//		this.stack = append(this.stack, nestedList[0].GetList())
//	}
//	return false
//}

//leetcode submit region end(Prohibit modification and deletion)
