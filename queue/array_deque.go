package main

type arrayDeque struct {
	nums        []int // 用于存储双向队列元素的数组
	front       int   // 队首指针，指向队首元素
	queSize     int   // 双向队列长度
	queCapacity int   // 队列容量（即最大容纳元素数量）
}

/* 初始化队列 */
func newArrayDeque(queCapacity int) *arrayDeque {
	return &arrayDeque{queCapacity: queCapacity, front: 0, queSize: 0, nums: make([]int, queCapacity)}
}

/* 获取双向队列的长度 */
func (q *arrayDeque) size() int {
	return q.queSize
}

/* 判断双向队列是否为空 */
func (q *arrayDeque) isEmpty() bool {
	return q.queSize == 0
}

func (q *arrayDeque) isFill() bool {
	return q.queSize == q.queCapacity
}

/* 计算环形数组索引 */
func (q *arrayDeque) index(i int) int {
	// 通过取余操作实现数组首尾相连
	// 当 i 越过数组尾部后，回到头部
	// 当 i 越过数组头部后，回到尾部
	return (i + q.queCapacity) % q.queCapacity
}

/* 队首入队 */
func (q *arrayDeque) pushFirst(num int) {
	if q.isFill() {
		return
	}
	q.front = q.index(q.front - 1)
	q.nums[q.front] = num
	q.queSize++
}

/* 队尾入队 */
func (q *arrayDeque) pushLast(num int) {
	if q.isFill() {
		return
	}
	real := (q.front + q.queSize) % q.queCapacity
	q.nums[real] = num
	q.queSize++
}

/* 访问队首元素 */
func (q *arrayDeque) peekFirst() any {
	if q.isEmpty() {
		return nil
	}
	return q.nums[q.front]
}

/* 队首出队 */
func (q *arrayDeque) popFirst() any {
	num := q.peekFirst()
	q.front = q.index(q.front + 1)
	return num
}

/* 访问队尾元素 */
func (q *arrayDeque) peekLast() any {
	if q.isEmpty() {
		return nil
	}
	real := q.index(q.front + q.queSize - 1)
	return q.nums[real]
}

/* 队尾出队 */
func (q *arrayDeque) popLast() any {
	num := q.peekLast()
	q.queSize--
	return num
}

/* 获取 Slice 用于打印 */
func (q *arrayDeque) toSlice() []int {
	// 仅转换有效长度范围内的列表元素
	res := make([]int, q.queSize)
	for i, j := 0, q.front; i < q.queSize; i++ {
		res[i] = q.nums[q.index(j)]
		j++
	}
	return res
}
