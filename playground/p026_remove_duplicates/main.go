package main

import "fmt"

// 删除有序数组中的重复项
func removeDuplicates(nums []int) int {
	var (
		k = 1
	)
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] { //这里对于是用if还是for循环拿捏不准，使用for循环感觉会让运行更复杂，还是用if，同时，对于是相等还是不等拿捏不准
			nums[k] = nums[i]
			k++
		}
	}
	return k
}

func main() {
	nums1 := []int{1, 1, 2}
	k1 := removeDuplicates(nums1)
	fmt.Println(k1, nums1[:k1]) // 期望 2 [1 2]

	nums2 := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	k2 := removeDuplicates(nums2)
	fmt.Println(k2, nums2[:k2]) // 期望 5 [0 1 2 3 4]

	nums3 := []int{1}
	k3 := removeDuplicates(nums3)
	fmt.Println(k3, nums3[:k3]) // 期望 1 [1]
}
