/**
A super ugly number is a positive integer whose prime factors are in the array
primes.

 Given an integer n and an array of integers primes, return the nᵗʰ super ugly
number.

 The nᵗʰ super ugly number is guaranteed to fit in a 32-bit signed integer.


 Example 1:


Input: n = 12, primes = [2,7,13,19]
Output: 32
Explanation: [1,2,4,7,8,13,14,16,19,26,28,32] is the sequence of the first 12
super ugly numbers given primes = [2,7,13,19].


 Example 2:


Input: n = 1, primes = [2,3,5]
Output: 1
Explanation: 1 has no prime factors, therefore all of its prime factors are in
the array primes = [2,3,5].



 Constraints:


 1 <= n <= 10⁵
 1 <= primes.length <= 100
 2 <= primes[i] <= 1000
 primes[i] is guaranteed to be a prime number.
 All the values of primes are unique and sorted in ascending order.


 Related Topics 数组 数学 动态规划 👍 411 👎 0

*/

// leetcode submit region begin(Prohibit modification and deletion)
package main

import "container/heap"

func nthSuperUglyNumber(n int, primes []int) int {
	pq := make(priorQueue, 0)
	ugly := []int{1}
	for _, prime := range primes {
		pq = append(pq, item{
			val:   prime,
			prime: prime,
			idx:   0,
		})
	}

	heap.Init(&pq)

	for n > 1 {
		pop := heap.Pop(&pq).(item)
		if pop.val != ugly[len(ugly)-1] {
			ugly = append(ugly, pop.val)
			n--
		}
		heap.Push(&pq, item{
			val:   ugly[pop.idx+1] * pop.prime,
			prime: pop.prime,
			idx:   pop.idx + 1,
		})
	}
	return ugly[len(ugly)-1]
}

type item struct {
	val   int
	prime int
	idx   int
}

type priorQueue []item

func (p priorQueue) Len() int {
	return len(p)
}

func (p priorQueue) Less(i, j int) bool {
	return p[i].val < p[j].val
}

func (p priorQueue) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p *priorQueue) Push(x any) {
	*p = append(*p, x.(item))
}

func (p *priorQueue) Pop() any {
	old := *p
	n := old.Len()
	v := old[n-1]
	*p = old[:n-1]
	return v
}

//leetcode submit region end(Prohibit modification and deletion)
