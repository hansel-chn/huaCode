/**
In the video game Fallout 4, the quest "Road to Freedom" requires players to
reach a metal dial called the "Freedom Trail Ring" and use the dial to spell a
specific keyword to open the door.

 Given a string ring that represents the code engraved on the outer ring and
another string key that represents the keyword that needs to be spelled, return
the minimum number of steps to spell all the characters in the keyword.

 Initially, the first character of the ring is aligned at the "12:00" direction.
 You should spell all the characters in key one by one by rotating ring
clockwise or anticlockwise to make each character of the string key aligned at the "12:0
0" direction and then by pressing the center button.

 At the stage of rotating the ring to spell the key character key[i]:


 You can rotate the ring clockwise or anticlockwise by one place, which counts
as one step. The final purpose of the rotation is to align one of ring's
characters at the "12:00" direction, where this character must equal key[i].
 If the character key[i] has been aligned at the "12:00" direction, press the
center button to spell, which also counts as one step. After the pressing, you
could begin to spell the next character in the key (next stage). Otherwise, you
have finished all the spelling.



 Example 1:


Input: ring = "godding", key = "gd"
Output: 4
Explanation:
For the first key character 'g', since it is already in place, we just need 1
step to spell this character.
For the second key character 'd', we need to rotate the ring "godding"
anticlockwise by two steps to make it become "ddinggo".
Also, we need 1 more step for spelling.
So the final output is 4.


 Example 2:


Input: ring = "godding", key = "godding"
Output: 13



 Constraints:


 1 <= ring.length, key.length <= 100
 ring and key consist of only lower case English letters.
 It is guaranteed that key could always be spelled by rotating ring.


 Related Topics 深度优先搜索 广度优先搜索 字符串 动态规划 👍 318 👎 0

*/

// leetcode submit region begin(Prohibit modification and deletion)
package main

import (
	"math"
)

func findRotateSteps(ring string, key string) int {
	ringMap := make(map[byte][]int)
	for i := 0; i < len(ring); i++ {
		if _, ok := ringMap[ring[i]]; !ok {
			ringMap[ring[i]] = []int{i}
		} else {
			ringMap[ring[i]] = append(ringMap[ring[i]], i)
		}
	}

	dp := make([][]int, len(ring))
	for i := range dp {
		dp[i] = make([]int, len(key))
	}

	for j := 0; j < len(key); j++ {
		//for i:=0;i< len(ring);i++ {
		//
		//}
		if j == 0 {
			idxes := ringMap[key[0]]
			for _, idx := range idxes {
				dp[idx][j] = getShortestDistance(len(ring), 0, idx) + 1
			}
			continue
		}

		preIdxes := ringMap[key[j-1]]
		idxes := ringMap[key[j]]
		for _, nowIdx := range idxes {
			minDistance := math.MaxInt
			for _, preIdx := range preIdxes {
				minDistance = min(minDistance, dp[preIdx][j-1]+getShortestDistance(len(ring), nowIdx, preIdx))
			}
			dp[nowIdx][j] = minDistance + 1
		}
	}

	rlt := math.MaxInt
	lastCharIdxes := ringMap[key[len(key)-1]]
	for _,idx := range lastCharIdxes {
		rlt = min(rlt, dp[idx][len(key)-1])
	}
	//fmt.Println(dp)
	//fmt.Println(ringMap)
	return rlt
}

func getShortestDistance(ringLen int, i, j int) int {
	abs := i - j
	if abs < 0 {
		abs = -abs
	}
	return min(abs, ringLen-abs)
}

//leetcode submit region end(Prohibit modification and deletion)
