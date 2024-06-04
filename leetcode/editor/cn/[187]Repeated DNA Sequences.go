/**
The DNA sequence is composed of a series of nucleotides abbreviated as 'A', 'C',
 'G', and 'T'.


 For example, "ACGAATTCCG" is a DNA sequence.


 When studying DNA, it is useful to identify repeated sequences within the DNA.


 Given a string s that represents a DNA sequence, return all the 10-letter-long
sequences (substrings) that occur more than once in a DNA molecule. You may
return the answer in any order.


 Example 1:
 Input: s = "AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT"
Output: ["AAAAACCCCC","CCCCCAAAAA"]

 Example 2:
 Input: s = "AAAAAAAAAAAAA"
Output: ["AAAAAAAAAA"]


 Constraints:


 1 <= s.length <= 10⁵
 s[i] is either 'A', 'C', 'G', or 'T'.


 Related Topics 位运算 哈希表 字符串 滑动窗口 哈希函数 滚动哈希 👍 586 👎 0

*/

// leetcode submit region begin(Prohibit modification and deletion)
package main

import "math"

func findRepeatedDnaSequences(s string) []string {
	existMap := make(map[int]struct{})
	dnaMap := map[byte]int{'A': 0, 'C': 1, 'G': 2, 'T': 3}

	var sys int = 4
	var top int = 10
	var left, right int = 0, 0
	var remove int = int(math.Pow(float64(sys), float64(top-1)))
	var sValue int
	rltMap := make(map[string]struct{})
	var rlt []string

	for right < len(s) {
		sValue = sValue*sys + dnaMap[s[right]]
		if right-left+1 == top {
			if _, ok := existMap[sValue]; ok {
				if _, rltMapOk := rltMap[s[left:right+1]]; !rltMapOk {
					rltMap[s[left:right+1]] = struct{}{}
				}
			} else {
				existMap[sValue] = struct{}{}
			}

			sValue -= dnaMap[s[left]] * remove
			left++
		}
		right++
	}

	for k := range rltMap {
		rlt = append(rlt, k)
	}

	return rlt
}

//func findRepeatedDnaSequences(s string) []string {
//	existMap := make(map[int][]string)
//	dnaMap := map[byte]int{'A': 0, 'C': 1, 'G': 2, 'T': 3}
//	var prime int = 1000000007
//
//	var sys int = 4
//	var top int = 10
//	var remove int = int(math.Pow(float64(sys), float64(top-1)))
//	var left, right int = 0, 0
//	var sValue int
//	rltMap:=make(map[string]struct{})
//	var rlt []string
//
//	for right < len(s) {
//		sValue = (sValue*sys%prime + dnaMap[s[right]]) % prime
//		if right-left+1 == top {
//			flag := 0
//			if vSlice, ok := existMap[sValue]; ok {
//				for _, v := range vSlice {
//					if v == s[left:right+1] {
//						if _, rltMapOk := rltMap[v]; !rltMapOk {
//							rltMap[v] = struct{}{}
//						}
//						flag = 1
//					}
//				}
//			}
//
//			if flag == 0 {
//				existMap[sValue] = append(existMap[sValue], s[left:right+1])
//			}
//
//			sValue = (sValue - (dnaMap[s[left]]*remove)%prime + prime) % prime
//			left++
//		}
//		right++
//	}
//
//	for k := range rltMap {
//		rlt = append(rlt, k)
//	}
//
//	return rlt
//}

//leetcode submit region end(Prohibit modification and deletion)
