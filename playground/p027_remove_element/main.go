package main

import "fmt"

// removeElement 原地删除 nums 中所有等于 val 的元素，返回剩余元素个数 k（27. 移除元素）。
//
// 思路演变（我自己想到的几种）：
//
//	① 双层 for + 整体左移：O(n²)。不是最优解——注意本题 n ≤ 100，n² = 10⁴ 其实能过，
//	   真正的理由是它写起来更容易出错，且有更简单的 O(n) 写法。
//	② map 过滤：能做到 O(n)，但要 O(n) 额外空间、还得把结果写回 nums，绕路；而且 map 不保序。
//	③ nums[:0] + append（本文件实现）：把 nums 的前面几格当成"写缓冲区"。
//	   nums[:0] 与 nums 共用底层数组 ⇒ append 就是原地写回原数组；
//	   写入位置（len(num1)）永远 ≤ 当前读位置 ⇒ 不会覆盖还没读到的元素；
//	   保留元素数 ≤ len(nums) ≤ cap(nums) ⇒ append 永远不需要扩容 ⇒ O(n) 时间 / O(1) 空间。
//	④ 显式下标的快慢双指针（与 ③ 等价，面试更好讲）：
//	   w := 0; for i := range nums { if nums[i] != val { nums[w] = nums[i]; w++ } }; return w
func removeElement(nums []int, val int) int {
	num1 := nums[:0]
	for _, v := range nums {
		if v != val {
			num1 = append(num1, v)
		}
	}
	return len(num1)
}

func main() {
	nums1 := []int{3, 2, 2, 3}
	k1 := removeElement(nums1, 3)
	fmt.Println(k1, nums1[:k1]) // 期望 2 [2 2]

	nums2 := []int{0, 1, 2, 2, 3, 0, 4, 2}
	k2 := removeElement(nums2, 2)
	fmt.Println(k2, nums2[:k2]) // 期望 5，前 5 个是 {0,0,1,3,4}（顺序任意，判题会排序）

	nums3 := []int{}
	k3 := removeElement(nums3, 2)
	fmt.Println(k3, nums3[:k3]) // 期望 0 []（题目允许空数组，这里不会 panic）

	nums4 := []int{2, 2, 2}
	k4 := removeElement(nums4, 2)
	fmt.Println(k4, nums4[:k4]) // 期望 0 []
}
