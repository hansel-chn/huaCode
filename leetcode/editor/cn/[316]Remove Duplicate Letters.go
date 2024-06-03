/**
Given a string s, remove duplicate letters so that every letter appears once
and only once. You must make sure your result is the smallest in lexicographical
order among all possible results.


 Example 1:


Input: s = "bcabc"
Output: "abc"


 Example 2:


Input: s = "cbacdcbc"
Output: "acdb"



 Constraints:


 1 <= s.length <= 10⁴
 s consists of lowercase English letters.



 Note: This question is the same as 1081: https://leetcode.com/problems/
smallest-subsequence-of-distinct-characters/

 Related Topics 栈 贪心 字符串 单调栈 👍 1084 👎 0

*/

// leetcode submit region begin(Prohibit modification and deletion)
package main

func removeDuplicateLetters(s string) string {
	var countMap [26]int
	for i := 0; i < len(s); i++ {
		countMap[s[i]-'a']++
	}

	var existMap [26]int

	stack := make([]byte, 0)

	for i := 0; i < len(s); i++ {
		idx := s[i] - 'a'
		if existMap[idx] == 1 {
			countMap[idx]--
			continue
		}

		for len(stack) > 0 && countMap[stack[len(stack)-1]-'a'] > 1 && stack[len(stack)-1] > s[i] {
			countMap[stack[len(stack)-1]-'a']--
			existMap[stack[len(stack)-1]-'a'] = 0
			stack = stack[:len(stack)-1]
		}
		existMap[idx] = 1
		stack = append(stack, s[i])
	}
	return string(stack)
}

//leetcode submit region end(Prohibit modification and deletion)
