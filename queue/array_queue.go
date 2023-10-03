package main

import "fmt"

type arrayQueue struct {
	nums        []int
	front       int
	queSize     int
	queCapacity int
}

func newArrayQueue(queCapacity int) *arrayQueue {
	return &arrayQueue{front: 0, queSize: 0, queCapacity: queCapacity, nums: make([]int, queCapacity)}
}

func (q *arrayQueue) size() int {
	return q.queSize
}

/* 判断队列是否为空 */
func (q *arrayQueue) isEmpty() bool {
	return q.queSize == 0
}

func (q *arrayQueue) push(num int) {
	if q.queSize == q.queCapacity {
		return
	}
	real := (q.front + q.queSize) % q.queCapacity
	q.nums[real] = num
	q.queSize++
}

/* 访问队首元素 */
func (q *arrayQueue) peek() any {
	if q.isEmpty() {
		return nil
	}
	return q.nums[q.front]
}

/* 出队 */
func (q *arrayQueue) pop() any {
	num := q.peek()
	q.front = (q.front + 1) % q.queCapacity
	q.queSize--

	return num
}

/* 获取 Slice 用于打印 */
func (q *arrayQueue) toSlice() []int {
	rear := (q.front + q.queSize)
	if rear >= q.queCapacity {
		rear %= q.queCapacity
		return append(q.nums[q.front:], q.nums[:rear]...)
	}
	return q.nums[q.front:rear]
}

func main() {
	que := newArrayQueue(4)
	que.push(1)
	que.push(2)
	que.push(3)
	que.push(4)
	que.pop()
	que.push(5)

	fmt.Println(que.toSlice())
}
