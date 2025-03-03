//A phrase is a palindrome if, after converting all uppercase letters into
//lowercase letters and removing all non-alphanumeric characters, it reads the same
//forward and backward. Alphanumeric characters include letters and numbers.
//
// Given a string s, return true if it is a palindrome, or false otherwise.
//
//
// Example 1:
//
//
//Input: s = "A man, a plan, a canal: Panama"
//Output: true
//Explanation: "amanaplanacanalpanama" is a palindrome.
//
//
// Example 2:
//
//
//Input: s = "race a car"
//Output: false
//Explanation: "raceacar" is not a palindrome.
//
//
// Example 3:
//
//
//Input: s = " "
//Output: true
//Explanation: s is an empty string "" after removing non-alphanumeric
//characters.
//Since an empty string reads the same forward and backward, it is a palindrome.
//
//
//
//
// Constraints:
//
//
// 1 <= s.length <= 2 * 10⁵
// s consists only of printable ASCII characters.
//
//
// Related Topics 双指针 字符串 👍 799 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
package main

import (
	"strings"
)

func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	bs := []byte(s)
	var slow, fast int
	for fast < len(bs) {
		if checkValid125(bs[fast]) {
			bs[slow] = bs[fast]
			slow++
			fast++
		} else {
			fast++
		}
	}

	var l, r = 0, slow - 1
	for l < r {
		if bs[l] != bs[r] {
			//fmt.Println(bs[l], bs[r])
			return false
		}
		l++
		r--
	}

	return true
}

func checkValid125(c byte) bool {
	if ('0' <= c && c <= '9') || ('a' <= c && c <= 'z') {
		return true
	} else {
		return false
	}
}

//leetcode submit region end(Prohibit modification and deletion)
