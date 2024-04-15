/**
Given an array of integers temperatures represents the daily temperatures,
return an array answer such that answer[i] is the number of days you have to wait
after the iᵗʰ day to get a warmer temperature. If there is no future day for which
this is possible, keep answer[i] == 0 instead.


 Example 1:
 Input: temperatures = [73,74,75,71,69,72,76,73]
Output: [1,1,4,2,1,1,0,0]

 Example 2:
 Input: temperatures = [30,40,50,60]
Output: [1,1,1,0]

 Example 3:
 Input: temperatures = [30,60,90]
Output: [1,1,0]


 Constraints:


 1 <= temperatures.length <= 10⁵
 30 <= temperatures[i] <= 100


 Related Topics 栈 数组 单调栈 👍 1746 👎 0

*/

// leetcode submit region begin(Prohibit modification and deletion)
package main

func dailyTemperatures(temperatures []int) []int {
	// 0 represent index; 1represent value
	monoStack := make([][2]int, 0)
	rlt := make([]int, len(temperatures))

	for i := len(temperatures) - 1; i >= 0; i-- {
		for len(monoStack) > 0 && temperatures[i] >= monoStack[len(monoStack)-1][1] {
			monoStack = monoStack[:len(monoStack)-1]
		}

		if len(monoStack) > 0 {
			rlt[i] = monoStack[len(monoStack)-1][0] - i
		} else {

			rlt[i] = 0
		}
		monoStack = append(monoStack, [2]int{i, temperatures[i]})
	}
	return rlt
}

//leetcode submit region end(Prohibit modification and deletion)
