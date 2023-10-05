package main

import "fmt"

/* 计数排序 */
// 简单实现，无法用于排序对象
func countingSortNaive(nums []int) {
	// 1. 统计数组最大元素 m
	maxVal := nums[0]
	for _, v := range nums {
		if v > maxVal {
			maxVal = v
		}
	}
	// 2. 统计各数字的出现次数
	// counter[num] 代表 num 的出现次数
	counter := make([]int, maxVal+1)
	for _, v := range nums {
		counter[v]++
	}
	// 3. 遍历 counter ，将各元素填入原数组 nums
	index := 0 //结果数组
	for i, v := range counter {
		for j := 0; j < v; j++ {
			nums[index] = i
			index++
		}
	}
}

/* 计数排序 */
// 完整实现，可排序对象，并且是稳定排序
func newcountingSort(nums []int) {
	// 1. 统计数组最大元素 m
	maxVal := nums[0]
	for _, v := range nums {
		if v > maxVal {
			maxVal = v
		}
	}
	// 2. 统计各数字的出现次数
	// counter[num] 代表 num 的出现次数
	counter := make([]int, maxVal+1)
	for _, v := range nums {
		counter[v]++
	}
	// 3. 求 counter 的前缀和，将“出现次数”转换为“尾索引”
	// 即 counter[num]-1 是 num 在 res 中最后一次出现的索引
	for i := 0; i < maxVal; i++ {
		counter[i+1] += counter[i]
	}
	res := make([]int, len(nums))
	for i := len(nums) - 1; i >= 0; i-- {
		num := nums[i]
		res[counter[num]-1] = num
		counter[num]--
	}
	// 使用结果数组 res 覆盖原数组 nums
	copy(nums, res)
}

func main() {
	s := []int{1, 2, 4, 5, 3, 2, 1, 2, 8, 1}
	newcountingSort(s)
	fmt.Println(s)
}
