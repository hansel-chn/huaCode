/**
给你一个整数数组 nums，请你将该数组升序排列。






 示例 1：


输入：nums = [5,2,3,1]
输出：[1,2,3,5]


 示例 2：


输入：nums = [5,1,1,2,0,0]
输出：[0,0,1,1,2,5]




 提示：


 1 <= nums.length <= 5 * 10⁴
 -5 * 10⁴ <= nums[i] <= 5 * 10⁴


 Related Topics 数组 分治 桶排序 计数排序 基数排序 排序 堆（优先队列） 归并排序 👍 930 👎 0

*/

// leetcode submit region begin(Prohibit modification and deletion)
package main

import "math"

//func swap(nums []int, a, b int) {
//	temp := nums[a]
//	nums[a] = nums[b]
//	nums[b] = temp
//}
//
//// based on two max subheap
//func maxHeapify(nums []int, startPos int, endPos int) {
//	large := startPos
//	if startPos*2+1 <= endPos && nums[large] < nums[startPos*2+1] {
//		large = startPos*2 + 1
//	}
//
//	if startPos*2+2 <= endPos && nums[large] < nums[startPos*2+2] {
//		large = startPos*2 + 2
//	}
//
//	if large != startPos {
//		swap(nums, startPos, large)
//		maxHeapify(nums, large, endPos)
//	}
//}
//
//func buildMaxHeap(nums []int, topPos int, bottomPos int) {
//	for i := (bottomPos - 1) >> 1; i >= topPos; i-- {
//		maxHeapify(nums, i, bottomPos)
//	}
//}
//
//// heap sort
//func heapSortArray(nums []int) []int {
//	buildMaxHeap(nums, 0, len(nums)-1)
//	for i := len(nums) - 1; i >= 1; i-- {
//		swap(nums, 0, i)
//		maxHeapify(nums, 0, i-1)
//	}
//	return nums
//}
////================================================================================
//// quick sort
//func swap(nums []int, a, b int) {
//	nums[a], nums[b] = nums[b], nums[a]
//}

//func quickSort(nums []int, left, right int) {
//	//rand.Seed(time.Now().UnixNano())
//	randNum := rand.Intn(right-left+1) + left
//	swap(nums, randNum, right)
//	flag := 0
//	pos := right
//	originLeft := left
//	originRight := right
//	for left < right {
//		if flag == 0 {
//			if nums[left] > nums[pos] {
//				swap(nums, left, pos)
//				pos = left
//				flag = 1
//				right--
//			} else {
//				left++
//			}
//		} else {
//			if nums[right] < nums[pos] {
//				swap(nums, right, pos)
//				pos = right
//				flag = 0
//				left++
//			} else {
//				right--
//			}
//		}
//	}
//	if originLeft < pos {
//		quickSort(nums, originLeft, pos-1)
//	}
//	if originRight > pos {
//		quickSort(nums, pos+1, originRight)
//	}
//	return
//}

////-------------------------------------------------------------------------
//// mid: val to be detected, lt pivot: [0,left], gt pivot: [right,length(nums)]
//// rand solve problem of fully reversed data
//// dual pivot and three-way partition quicksort solve problem of large amounts of duplicate data
//func quickSort(nums []int, originLeft, originRight int) {
//	if originLeft >= originRight {
//		return
//	}
//
//	//rand.Seed(time.Now().UnixNano())
//	randNum := rand.Intn(originRight-originLeft+1) + originLeft
//	pivot := nums[randNum]
//	swap(nums, randNum, originRight)
//	left := originLeft - 1
//	mid := originLeft
//	right := originRight
//	for mid < right {
//		if nums[mid] == pivot {
//			mid++
//		} else if nums[mid] < pivot {
//			left++
//			swap(nums, left, mid)
//			mid++
//		} else {
//			right--
//			swap(nums, right, mid)
//		}
//	}
//	swap(nums, right, originRight)
//
//	quickSort(nums, originLeft, left)
//	quickSort(nums, right+1, originRight)
//	return
//}
//
//func sortArray(nums []int) []int {
//	if len(nums) < 1 {
//		return nums
//	}
//	quickSort(nums, 0, len(nums)-1)
//	return nums
//}

//// ==============================================================================
//// bucketSort
//func sortArray(nums []int) []int {
//	if len(nums) < 2 {
//		return nums
//	}
//
//	minNum := 0
//	maxNum := 0
//	for i := 0; i < len(nums); i++ {
//		if nums[i] < minNum {
//			minNum = nums[i]
//		}
//		if nums[i] > maxNum {
//			maxNum = nums[i]
//		}
//	}
//
//	bucketNum := 5
//	bucketSize := (maxNum-minNum)/bucketNum + 1
//	buckets := make([][]int, 5)
//	for i := 0; i < len(nums); i++ {
//		index := (nums[i] - minNum) / bucketSize
//		buckets[index] = append(buckets[index], nums[i])
//	}
//
//	for _, bucket := range buckets {
//		sort.Ints(bucket)
//	}
//
//	key:=0
//	for _, bucket := range buckets {
//		for _, v := range bucket {
//			nums[key]=v
//			key++
//		}
//	}
//	return nums
//}

//// ====================================================================
//// mergeSort
//func mergeSort(arr []int) []int {
//	if len(arr) < 2 {
//		return arr
//	}
//	mid := (len(arr) - 1) / 2
//	sortedArr1 := mergeSort(arr[0 : mid+1])
//	sortedArr2 := mergeSort(arr[mid+1:])
//	arr1Len := len(sortedArr1)
//	arr2Len := len(sortedArr2)
//
//	newArr := make([]int, arr2Len+arr1Len)
//	idx, arr1Idx, arr2Idx := 0, 0, 0
//	for arr1Idx < arr1Len && arr2Idx < arr2Len {
//		if sortedArr1[arr1Idx] < sortedArr2[arr2Idx] {
//			newArr[idx] = sortedArr1[arr1Idx]
//			arr1Idx++
//			idx++
//		} else {
//			newArr[idx] = sortedArr2[arr2Idx]
//			arr2Idx++
//			idx++
//		}
//	}
//
//	if arr1Idx == arr1Len {
//		copy(newArr[idx:], sortedArr2[arr2Idx:])
//	} else {
//		copy(newArr[idx:], sortedArr1[arr1Idx:])
//	}
//	return newArr
//}
//func sortArray(nums []int) []int {
//	newArr := mergeSort(nums)
//	return newArr
//}

// ==============================================================
// count sort
// The code shows the case when data in int slice are all non-negative integers.
func countSort(nums []int) []int {
	max := math.MinInt
	for _, num := range nums {
		if max < num {
			max = num
		}
	}

	countArr := make([]int, max+1)
	for _, num := range nums {
		countArr[num]++
	}

	for i := 1; i < max+1; i++ {
		countArr[i] += countArr[i-1]
	}

	newArr := make([]int, len(nums))
	for i := len(nums) - 1; i >= 0; i-- {
		newArr[countArr[nums[i]]-1] = nums[i]
		countArr[nums[i]]--
	}
	return newArr
}

func sortArray(nums []int) []int {
	return countSort(nums)

}

//leetcode submit region end(Prohibit modification and deletion)
