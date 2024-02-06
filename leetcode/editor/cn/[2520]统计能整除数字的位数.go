/**
给你一个整数 num ，返回 num 中能整除 num 的数位的数目。

 如果满足 nums % val == 0 ，则认为整数 val 可以整除 nums 。



 示例 1：

 输入：num = 7
输出：1
解释：7 被自己整除，因此答案是 1 。


 示例 2：

 输入：num = 121
输出：2
解释：121 可以被 1 整除，但无法被 2 整除。由于 1 出现两次，所以返回 2 。


 示例 3：

 输入：num = 1248
输出：4
解释：1248 可以被它每一位上的数字整除，因此答案是 4 。




 提示：


 1 <= num <= 10⁹
 num 的数位中不含 0


 Related Topics 数学 👍 30 👎 0

*/

// leetcode submit region begin(Prohibit modification and deletion)
package main

func countDigits(num int) int {
	count := 0
	temp := num
	for temp > 0 {
		reminder := temp % 10
		// reminder not exist 0
		if num%reminder == 0 {
			count++
		}
		temp = temp / 10
	}
	return count
}

//func countDigits(num int) int {
//	count := 0
//	hash := make(map[int]struct{})
//	temp:=num
//	for temp> 0 {
//		reminder := temp% 10
//		if _, ok := hash[reminder]; ok {
//			count++
//		} else {
//			// reminder not exist 0
//			if num%reminder == 0 {
//				hash[reminder] = struct{}{}
//				count++
//			}
//		}
//		temp = temp / 10
//	}
//	return count
//}

//leetcode submit region end(Prohibit modification and deletion)
