/**
给你一个正整数 n ，请你返回 n 的 惩罚数 。

 n 的 惩罚数 定义为所有满足以下条件 i 的数的平方和：


 1 <= i <= n
 i * i 的十进制表示的字符串可以分割成若干连续子字符串，且这些子字符串对应的整数值之和等于 i 。




 示例 1：


输入：n = 10
输出：182
解释：总共有 3 个整数 i 满足要求：
- 1 ，因为 1 * 1 = 1
- 9 ，因为 9 * 9 = 81 ，且 81 可以分割成 8 + 1 。
- 10 ，因为 10 * 10 = 100 ，且 100 可以分割成 10 + 0 。
因此，10 的惩罚数为 1 + 81 + 100 = 182


 示例 2：


输入：n = 37
输出：1478
解释：总共有 4 个整数 i 满足要求：
- 1 ，因为 1 * 1 = 1
- 9 ，因为 9 * 9 = 81 ，且 81 可以分割成 8 + 1 。
- 10 ，因为 10 * 10 = 100 ，且 100 可以分割成 10 + 0 。
- 36 ，因为 36 * 36 = 1296 ，且 1296 可以分割成 1 + 29 + 6 。
因此，37 的惩罚数为 1 + 81 + 100 + 1296 = 1478




 提示：


 1 <= n <= 1000


 Related Topics 数学 回溯 👍 97 👎 0

*/

// leetcode submit region begin(Prohibit modification and deletion)
package main

//func punishmentNumber(n int) int {
//	if n == 0 {
//		return 0
//	}
//	result := 0
//	for i := 1; i <= n; i++ {
//		num := strconv.Itoa(i * i)
//		if ok := CheckSumOfSplitNums(num, 0, i); ok {
//			result += i*i
//		}
//	}
//	return result
//}
//
//func CheckSumOfSplitNums(num string, sum int, target int) bool {
//	if len(num) == 0 {
//		if sum == target {
//			return true
//		} else {
//			return false
//		}
//	}
//
//	for i := 0; i < len(num); i++ {
//		intNum, _ := strconv.Atoi(num[:i+1])
//		if sum+intNum > target {
//			break
//		}
//		rlt := CheckSumOfSplitNums(num[i+1:], sum+intNum, target)
//		if rlt == true {
//			return true
//		}
//	}
//	return false
//}

func punishmentNumber(n int) int {
	if n == 0 {
		return 0
	}
	result := 0
	for i := 1; i <= n; i++ {
		if ok := CheckSumOfSplitNums(i*i, 0, i); ok {
			result += i * i
		}
	}
	return result
}

func CheckSumOfSplitNums(num int, sum int, target int) bool {
	if num == 0 {
		if sum == target {
			return true
		} else {
			return false
		}
	}
	divisor := 10
	for num > 0 {
		check := CheckSumOfSplitNums(num/divisor, sum+num%divisor, target)
		if check {
			return true
		}
		if sum+num%divisor > target {
			return false
		}
		if num/divisor == 0 {
			break
		}
		divisor *= 10
	}
	return false
}

//leetcode submit region end(Prohibit modification and deletion)
