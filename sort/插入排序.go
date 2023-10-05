package main

import "fmt"

/* 插入排序 */
func insertionSort(nums []int) {
	n := len(nums)
	for i := 1; i < n; i++ {
		base := nums[i]
		for j := i - 1; j >= 0; j-- {
			if nums[j] < base {
				nums[j+1] = base
				break
			}
			nums[j+1] = nums[j]
		}
	}
}

func main() {
	s := []int{1, 2, 4, 5, 3, 2, 1, 2, 8, 1}
	insertionSort(s)
	fmt.Println(s)
}
