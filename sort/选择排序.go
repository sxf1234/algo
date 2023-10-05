package main

import "fmt"

func selectionSort(nums []int) {
	n := len(nums)
	for i := 0; i < n-1; i++ {
		min := nums[i]
		for j := i + 1; j < n; j++ {
			if nums[j] < min {
				min = nums[j]
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
}

func main() {
	s := []int{1, 2, 4, 5, 3, 2, 1, 2, 8, 1}
	selectionSort(s)
	fmt.Println(s)
}
