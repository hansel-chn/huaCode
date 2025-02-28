//Design and implement a data structure for a Least Frequently Used (LFU) cache.
//
//
// Implement the LFUCache class:
//
//
// LFUCache(int capacity) Initializes the object with the capacity of the data
//structure.
// int get(int key) Gets the value of the key if the key exists in the cache.
//Otherwise, returns -1.
// void put(int key, int value) Update the value of the key if present, or
//inserts the key if not already present. When the cache reaches its capacity, it
//should invalidate and remove the least frequently used key before inserting a new
//item. For this problem, when there is a tie (i.e., two or more keys with the same
//frequency), the least recently used key would be invalidated.
//
//
// To determine the least frequently used key, a use counter is maintained for
//each key in the cache. The key with the smallest use counter is the least
//frequently used key.
//
// When a key is first inserted into the cache, its use counter is set to 1 (
//due to the put operation). The use counter for a key in the cache is incremented
//either a get or put operation is called on it.
//
// The functions get and put must each run in O(1) average time complexity.
//
//
// Example 1:
//
//
//Input
//["LFUCache", "put", "put", "get", "put", "get", "get", "put", "get", "get",
//"get"]
//[[2], [1, 1], [2, 2], [1], [3, 3], [2], [3], [4, 4], [1], [3], [4]]
//Output
//[null, null, null, 1, null, -1, 3, null, -1, 3, 4]
//
//Explanation
//// cnt(x) = the use counter for key x
//// cache=[] will show the last used order for tiebreakers (leftmost element
//is  most recent)
//LFUCache lfu = new LFUCache(2);
//lfu.put(1, 1);   // cache=[1,_], cnt(1)=1
//lfu.put(2, 2);   // cache=[2,1], cnt(2)=1, cnt(1)=1
//lfu.get(1);      // return 1
//                 // cache=[1,2], cnt(2)=1, cnt(1)=2
//lfu.put(3, 3);   // 2 is the LFU key because cnt(2)=1 is the smallest,
//invalidate 2.
//                 // cache=[3,1], cnt(3)=1, cnt(1)=2
//lfu.get(2);      // return -1 (not found)
//lfu.get(3);      // return 3
//                 // cache=[3,1], cnt(3)=2, cnt(1)=2
//lfu.put(4, 4);   // Both 1 and 3 have the same cnt, but 1 is LRU, invalidate 1
//.
//                 // cache=[4,3], cnt(4)=1, cnt(3)=2
//lfu.get(1);      // return -1 (not found)
//lfu.get(3);      // return 3
//                 // cache=[3,4], cnt(4)=1, cnt(3)=3
//lfu.get(4);      // return 4
//                 // cache=[4,3], cnt(4)=2, cnt(3)=3
//
//
//
// Constraints:
//
//
// 1 <= capacity <= 10⁴
// 0 <= key <= 10⁵
// 0 <= value <= 10⁹
// At most 2 * 10⁵ calls will be made to get and put.
//
//
//
//
//
// Related Topics 设计 哈希表 链表 双向链表 👍 869 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
package main

import (
	"container/list"
)

type Node460 struct {
	key, value, freq int
}

type LFUCache struct {
	key2Elem  map[int]*list.Element
	freq2List map[int]*list.List
	minFreq   int
	capacity  int
}

func Constructor460(capacity int) LFUCache {
	return LFUCache{
		key2Elem:  make(map[int]*list.Element),
		freq2List: make(map[int]*list.List),
		minFreq:   0,
		capacity:  capacity,
	}
}

func (this *LFUCache) Get(key int) int {
	//if key == 12 {
	//	fmt.Println(len(this.key2Elem))
	//	fmt.Println("===============")
	//	for i, element := range this.key2Elem {
	//		fmt.Println(i, element.Value.(*Node460).value)
	//	}
	//	fmt.Println("===============")
	//}
	if elem, ok1 := this.key2Elem[key]; ok1 {
		node := elem.Value.(*Node460)

		this.increaseFreq(elem)

		return node.value
	}
	return -1
}

func (this *LFUCache) Put(key int, value int) {
	if elem, ok1 := this.key2Elem[key]; ok1 {
		node := elem.Value.(*Node460)
		node.value = value

		this.increaseFreq(elem)
	} else {
		newNode := &Node460{key: key, value: value, freq: 1}
		if len(this.key2Elem) >= this.capacity {
			minList := this.freq2List[this.minFreq]

			rmElem := minList.Back()
			rmKey := rmElem.Value.(*Node460).key
			minList.Remove(rmElem)
			delete(this.key2Elem, rmKey)

			if minList.Len() == 0 {
				delete(this.freq2List, this.minFreq)
			}
		}

		newElem := this.addElemInList(newNode)
		this.key2Elem[key] = newElem
		this.minFreq = 1
	}
}

func (this *LFUCache) increaseFreq(elem *list.Element) {
	node := elem.Value.(*Node460)

	ifDelLst := this.delElemInList(elem)

	if ifDelLst && this.minFreq == node.freq {
		this.minFreq = node.freq + 1
	}

	node.freq++

	newElem := this.addElemInList(node)
	this.key2Elem[node.key] = newElem
}

func (this *LFUCache) delElemInList(elem *list.Element) (delList bool) {
	node := elem.Value.(*Node460)

	listForDel := this.freq2List[node.freq]
	listForDel.Remove(elem)

	if listForDel.Len() == 0 {
		delete(this.freq2List, node.freq)
		return true
	}

	return false
}

func (this *LFUCache) addElemInList(node *Node460) *list.Element {
	if listForAdd, ok2 := this.freq2List[node.freq]; ok2 {
		return listForAdd.PushFront(node)
	} else {
		lst := list.New()
		this.freq2List[node.freq] = lst
		return lst.PushFront(node)
	}
}

/**
 * Your LFUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
//leetcode submit region end(Prohibit modification and deletion)
