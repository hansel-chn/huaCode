//You are given a string s, a string chars of distinct characters and an
//integer array vals of the same length as chars.
//
// The cost of the substring is the sum of the values of each character in the
//substring. The cost of an empty string is considered 0.
//
// The value of the character is defined in the following way:
//
//
// If the character is not in the string chars, then its value is its
//corresponding position (1-indexed) in the alphabet.
//
//
//
// For example, the value of 'a' is 1, the value of 'b' is 2, and so on. The
//value of 'z' is 26.
//
//
// Otherwise, assuming i is the index where the character occurs in the string
//chars, then its value is vals[i].
//
//
// Return the maximum cost among all substrings of the string s.
//
//
// Example 1:
//
//
//Input: s = "adaa", chars = "d", vals = [-1000]
//Output: 2
//Explanation: The value of the characters "a" and "d" is 1 and -1000
//respectively.
//The substring with the maximum cost is "aa" and its cost is 1 + 1 = 2.
//It can be proven that 2 is the maximum cost.
//
//
// Example 2:
//
//
//Input: s = "abc", chars = "abc", vals = [-1,-1,-1]
//Output: 0
//Explanation: The value of the characters "a", "b" and "c" is -1, -1, and -1
//respectively.
//The substring with the maximum cost is the empty substring "" and its cost is
//0.
//It can be proven that 0 is the maximum cost.
//
//
//
// Constraints:
//
//
// 1 <= s.length <= 10⁵
// s consist of lowercase English letters.
// 1 <= chars.length <= 26
// chars consist of distinct lowercase English letters.
// vals.length == chars.length
// -1000 <= vals[i] <= 1000
//
//
// Related Topics 数组 哈希表 字符串 动态规划 👍 26 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
package main

func maximumCostSubstring(s string, chars string, vals []int) int {
	char2val := make(map[byte]int)
	for i := 0; i < len(chars); i++ {
		char2val[chars[i]] = vals[i]
	}

	preSums := make([]int, len(s)+1)
	for i := 0; i < len(s); i++ {
		preSums[i+1] = preSums[i] + getChar2606(s[i], char2val)
	}

	minPreSum := 0
	rlt := 0
	for _, preSum := range preSums {
		rlt = max(rlt, preSum-minPreSum)
		minPreSum = min(minPreSum, preSum)
	}
	return rlt
}

//func maximumCostSubstring(s string, chars string, vals []int) int {
//	char2val := make(map[byte]int)
//	for i := 0; i < len(chars); i++ {
//		char2val[chars[i]] = vals[i]
//	}
//
//	dp := make([]int, len(s))
//
//	dp[0] = max(0, getChar2606(s[0], char2val))
//
//	for i := 1; i < len(s); i++ {
//		//dp[i] = max(dp[i-1]+getChar(s[i], char2val), getChar(s[i], char2val))
//		if dp[i-1] > 0 {
//			dp[i] = dp[i-1] + getChar2606(s[i], char2val)
//		} else {
//			dp[i] = getChar2606(s[i], char2val)
//		}
//	}
//
//	rlt := 0
//
//	for _, v := range dp {
//		rlt = max(rlt, v)
//	}
//	return rlt
//}

func getChar2606(c byte, char2val map[byte]int) int {
	if v, ok := char2val[c]; ok {
		return v
	} else {
		return int(c-'a') + 1
	}
}

//leetcode submit region end(Prohibit modification and deletion)
