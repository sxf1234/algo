package main

import "fmt"

/*
是的，当我们以最左端元素为基准数时，必须先“从右往左”再”从左往右“。这个结论有些反直觉，我来详细剖析一下原因。

哨兵划分 partition() 的最后一步是 nums[left] 与 nums[i] 交换，完成交换后基准数左边的元素都小于等于基准数，这也就要求 交换前 nums[left] >= nums[i] 必须要成立。

假设我们先“从左往右寻找第一个比基准数小的元素”，那么如果找不到，则会在 i == j 时跳出循环，此时 nums[j] == nums[i] > nums[left] ；也就是说，这种情况下执行最后一步交换就会出问题了，会把一个比基准数更大的元素交换至数组最左端，导致哨兵划分失败。
而如果我们先“从右往左”，就不会发生这个问题了。
*/
func partition(nums []int, left, right int) int {
	i := left
	j := right
	baseNum := nums[left]
	for i < j {
		//要先从右向左
		for nums[j] >= baseNum && i < j {
			j--
		}
		for nums[i] <= baseNum && i < j {
			i++
		}
		nums[i], nums[j] = nums[j], nums[i]
		fmt.Println(i, j)
		fmt.Println(nums)
	}
	nums[i], nums[left] = nums[left], nums[i]
	fmt.Println(nums)
	return i
}

/* 快速排序 */
func quickSort(nums []int, left, right int) {
	// 子数组长度为 1 时终止递归
	if left >= right {
		return
	}
	mid := newpartition(nums, left, right)
	quickSort(nums, left, mid-1)
	quickSort(nums, mid+1, right)
}

/* 选取三个元素的中位数 */
func medianThree(nums []int, left, mid, right int) int {
	// 此处使用异或运算来简化代码（!= 在这里起到异或的作用）
	// 异或规则为 0 ^ 0 = 1 ^ 1 = 0, 0 ^ 1 = 1 ^ 0 = 1
	if (nums[left] < nums[mid]) != (nums[left] < nums[right]) {
		return left
	} else if (nums[mid] < nums[left]) != (nums[mid] < nums[right]) {
		return mid
	}
	return right
}

/* 哨兵划分（三数取中值）*/
func newpartition(nums []int, left, right int) int {
	// 以 nums[left] 作为基准数
	med := medianThree(nums, left, (left+right)/2, right)
	// 将中位数交换至数组最左端
	nums[left], nums[med] = nums[med], nums[left]
	// 以 nums[left] 作为基准数
	i, j := left, right
	for i < j {
		for i < j && nums[j] >= nums[left] {
			j-- //从右向左找首个小于基准数的元素
		}
		for i < j && nums[i] <= nums[left] {
			i++ //从左向右找首个大于基准数的元素
		}
		//元素交换
		nums[i], nums[j] = nums[j], nums[i]
	}
	//将基准数交换至两子数组的分界线
	nums[i], nums[left] = nums[left], nums[i]
	return i //返回基准数的索引
}

/* 快速排序（尾递归优化）*/
func quickSortnew(nums []int, left, right int) {
	// 子数组长度为 1 时终止
	for left < right {
		// 哨兵划分操作
		pivot := partition(nums, left, right)
		// 对两个子数组中较短的那个执行快排
		if pivot-left < right-pivot {
			quickSortnew(nums, left, pivot-1) // 递归排序左子数组
			left = pivot + 1                  // 剩余未排序区间为 [pivot + 1, right]
		} else {
			quickSortnew(nums, pivot+1, right) // 递归排序右子数组
			right = pivot - 1                  // 剩余未排序区间为 [left, pivot - 1]
		}
	}
}
func main() {
	s := []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
	quickSortnew(s, 0, 8)
	fmt.Println(s)
}
