/**
Given two strings s and p, return an array of all the start indices of p's
anagrams in s. You may return the answer in any order.

 An Anagram is a word or phrase formed by rearranging the letters of a
different word or phrase, typically using all the original letters exactly once.


 Example 1:


Input: s = "cbaebabacd", p = "abc"
Output: [0,6]
Explanation:
The substring with start index = 0 is "cba", which is an anagram of "abc".
The substring with start index = 6 is "bac", which is an anagram of "abc".


 Example 2:


Input: s = "abab", p = "ab"
Output: [0,1,2]
Explanation:
The substring with start index = 0 is "ab", which is an anagram of "ab".
The substring with start index = 1 is "ba", which is an anagram of "ab".
The substring with start index = 2 is "ab", which is an anagram of "ab".



 Constraints:


 1 <= s.length, p.length <= 3 * 10⁴
 s and p consist of lowercase English letters.


 Related Topics 哈希表 字符串 滑动窗口 👍 1446 👎 0

*/

// leetcode submit region begin(Prohibit modification and deletion)
package main

func findAnagrams(s string, p string) []int {
	pMap := make(map[byte]int)
	winMap := make(map[byte]int)
	for i := 0; i < len(p); i++ {
		pMap[p[i]]++
	}

	var left, right, valid int
	rlt := make([]int, 0)

	for right < len(s) {
		addV := s[right]
		right++
		if _, ok := pMap[addV]; ok {
			winMap[addV]++
			if winMap[addV] == pMap[addV] {
				valid++
			}
		}

		for right-left >= len(p) {
			if valid == len(pMap) {
				rlt = append(rlt, left)
			}

			if _, ok := pMap[s[left]]; ok {
				if winMap[s[left]] == pMap[s[left]] {
					valid--
				}
				winMap[s[left]]--
			}
			left++
		}
	}
	return rlt
}

//leetcode submit region end(Prohibit modification and deletion)
