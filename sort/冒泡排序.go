package main

import "fmt"

/* 冒泡排序 */
func bubbleSort(nums []int) {
	n := len(nums)
	for i := n - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if nums[j] > nums[j+1] {
				nums[j+1], nums[j] = nums[j], nums[j+1]
			}
		}
	}
}
func main() {
	s := []int{1, 2, 4, 5, 3, 2, 1, 2, 8, 1}
	bubbleSort(s)
	fmt.Println(s)
}
