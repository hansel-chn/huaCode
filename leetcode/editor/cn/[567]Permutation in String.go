/**
Given two strings s1 and s2, return true if s2 contains a permutation of s1, or
false otherwise.

 In other words, return true if one of s1's permutations is the substring of s2.



 Example 1:


Input: s1 = "ab", s2 = "eidbaooo"
Output: true
Explanation: s2 contains one permutation of s1 ("ba").


 Example 2:


Input: s1 = "ab", s2 = "eidboaoo"
Output: false



 Constraints:


 1 <= s1.length, s2.length <= 10⁴
 s1 and s2 consist of lowercase English letters.


 Related Topics 哈希表 双指针 字符串 滑动窗口 👍 1006 👎 0

*/

// leetcode submit region begin(Prohibit modification and deletion)
package main

func checkInclusion(s1 string, s2 string) bool {
	s1Map := make(map[byte]int)
	s2Map := make(map[byte]int)
	for i := 0; i < len(s1); i++ {
		s1Map[s1[i]]++
	}

	var left, right, valid int
	for right < len(s2) {
		addV := s2[right]
		right++
		if _, ok := s1Map[addV]; ok {
			s2Map[addV]++
			if s2Map[addV] == s1Map[addV] {
				valid++
			}
		}

		for right-left >= len(s1) {
			if valid == len(s1Map) {
				return true
			}

			if _, ok := s1Map[s2[left]]; ok {
				if s2Map[s2[left]] == s1Map[s2[left]] {
					valid--
				}
				s2Map[s2[left]]--
			}
			left++
		}
	}
	return false
}

//leetcode submit region end(Prohibit modification and deletion)
