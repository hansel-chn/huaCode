/**
Design your implementation of the circular double-ended queue (deque).

 Implement the MyCircularDeque class:


 MyCircularDeque(int k) Initializes the deque with a maximum size of k.
 boolean insertFront() Adds an item at the front of Deque. Returns true if the
operation is successful, or false otherwise.
 boolean insertLast() Adds an item at the rear of Deque. Returns true if the
operation is successful, or false otherwise.
 boolean deleteFront() Deletes an item from the front of Deque. Returns true if
the operation is successful, or false otherwise.
 boolean deleteLast() Deletes an item from the rear of Deque. Returns true if
the operation is successful, or false otherwise.
 int getFront() Returns the front item from the Deque. Returns -1 if the deque
is empty.
 int getRear() Returns the last item from Deque. Returns -1 if the deque is
empty.
 boolean isEmpty() Returns true if the deque is empty, or false otherwise.
 boolean isFull() Returns true if the deque is full, or false otherwise.



 Example 1:


Input
["MyCircularDeque", "insertLast", "insertLast", "insertFront", "insertFront",
"getRear", "isFull", "deleteLast", "insertFront", "getFront"]
[[3], [1], [2], [3], [4], [], [], [], [4], []]
Output
[null, true, true, true, false, 2, true, true, true, 4]

Explanation
MyCircularDeque myCircularDeque = new MyCircularDeque(3);
myCircularDeque.insertLast(1);  // return True
myCircularDeque.insertLast(2);  // return True
myCircularDeque.insertFront(3); // return True
myCircularDeque.insertFront(4); // return False, the queue is full.
myCircularDeque.getRear();      // return 2
myCircularDeque.isFull();       // return True
myCircularDeque.deleteLast();   // return True
myCircularDeque.insertFront(4); // return True
myCircularDeque.getFront();     // return 4



 Constraints:


 1 <= k <= 1000
 0 <= value <= 1000
 At most 2000 calls will be made to insertFront, insertLast, deleteFront,
deleteLast, getFront, getRear, isEmpty, isFull.


 Related Topics 设计 队列 数组 链表 👍 234 👎 0

*/

// leetcode submit region begin(Prohibit modification and deletion)
package main

type MyCircularDeque struct {
	Front, Rear int
	Size, Cap   int
	queue       []int
}

//func Constructor(k int) MyCircularDeque {
//	return MyCircularDeque{
//		Front: 0,
//		Rear:  0,
//		Size:  0,
//		Cap:   k,
//		queue: make([]int, k),
//	}
//}

func (this *MyCircularDeque) InsertFront(value int) bool {
	if this.IsFull() {
		return false
	}

	if this.Size == 0 {
		this.Front = 0
		this.Rear = 0
		this.queue[this.Front] = value
		this.Size++
	} else {
		this.Front = (this.Front - 1 + this.Cap) % this.Cap
		this.queue[this.Front] = value
		this.Size++
	}

	return true
}

func (this *MyCircularDeque) InsertLast(value int) bool {
	if this.IsFull() {
		return false
	}

	if this.Size == 0 {
		this.Front = 0
		this.Rear = 0
		this.queue[this.Rear] = value
		this.Size++
	} else {
		this.Rear = (this.Rear + 1) % this.Cap
		this.queue[this.Rear] = value
		this.Size++
	}

	return true
}

func (this *MyCircularDeque) DeleteFront() bool {
	if this.IsEmpty() {
		return false
	}

	this.Front = (this.Front + 1) % this.Cap
	this.Size--
	return true
}

func (this *MyCircularDeque) DeleteLast() bool {
	if this.IsEmpty() {
		return false
	}

	this.Rear = (this.Rear - 1 + this.Cap) % this.Cap
	this.Size--
	return true
}

func (this *MyCircularDeque) GetFront() int {
	if this.IsEmpty() {
		return -1
	}
	return this.queue[this.Front]
}

func (this *MyCircularDeque) GetRear() int {
	if this.IsEmpty() {
		return -1
	}
	return this.queue[this.Rear]
}

func (this *MyCircularDeque) IsEmpty() bool {
	if this.Size == 0 {
		return true
	} else {
		return false
	}
}

func (this *MyCircularDeque) IsFull() bool {
	if this.Size == len(this.queue) {
		return true
	} else {
		return false
	}
}

/**
 * Your MyCircularDeque object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.InsertFront(value);
 * param_2 := obj.InsertLast(value);
 * param_3 := obj.DeleteFront();
 * param_4 := obj.DeleteLast();
 * param_5 := obj.GetFront();
 * param_6 := obj.GetRear();
 * param_7 := obj.IsEmpty();
 * param_8 := obj.IsFull();
 */
//leetcode submit region end(Prohibit modification and deletion)
