// 283. 移动零（LeetCode 283 · Easy）
//
// 题干：给定一个整数数组 nums，把所有 0 移到数组末尾，同时保持非零元素的相对顺序。
// 要求：必须原地操作，不能拷贝额外数组。
//
// 示例1：[0,1,0,3,12] → [1,3,12,0,0]
// 示例2：[0]           → [0]
//
// 运行：go run ./playground/p283_move_zeroes
// 测试：go test ./playground/p283_move_zeroes -v
// 和你写过的 27 移除元素不同的地方：**这题函数没有返回值**，
// 结果只能通过参数 nums 带出去——切片是引用语义，原地写回调用方看得见。
package main

import "fmt"

// moveZeroes 原地把所有 0 挪到末尾，保持非零元素的相对顺序。
//
// 思路（一句话）：不排序——把非零元素按原顺序原地覆盖到前面（nums[:0] 复用底层数组），再把剩下的位置清零
// 复杂度：时间 O(n)（扫一趟，每个元素最多写一次），空间 O(1)（只有下标和一个切片头，清零用 clear 不分配内存）
// ⟵ 助手修正：原来写的是「时间 O(1)，空间 O(n)」，正好写反；而且空间写成 O(n) 恰恰暴露了
//
//	原实现里 make([]int, zeroes) 真的分配了一块临时数组（见下方实现里的注释）。
//
// 提示（只给方向，不给答案）：往「快慢指针 + 覆盖」或「遇到非零就交换」两条路想，
// 两种都写出来，数一数各自往数组里写了几次。
func moveZeroes(nums []int) {
	// TODO 你的解法
	nal := nums[:0]
	for _, v := range nums {
		if v != 0 {
			nal = append(nal, v)
		}
	}
	// ⟵ 助手改：原来这里是
	//     zeroes := len(nums) - len(nal)
	//     nal = append(nal, make([]int, zeroes)...)
	//   make([]int, zeroes) 会分配一块 O(n) 的临时数组——题目要求「原地、不复制数组」，
	//   这块分配既违背要求，也把空间从 O(1) 抬成了 O(n)。Go 1.21+ 一行就够：
	clear(nums[len(nal):]) // nal 与 nums 共用底层数组，len(nal) 就是「非零区」的终点
}

func main() {
	// moveZeroes 会就地改切片，所以每次传副本进去，否则打印出来的是改完的结果看不出前后对比。
	run := func(in []int) {
		nums := append([]int(nil), in...)
		moveZeroes(nums)
		fmt.Printf("%v → %v\n", in, nums)
	}

	run([]int{0, 1, 0, 3, 12}) // 期望 [1 3 12 0 0]
	run([]int{0})              // 期望 [0]
	run([]int{0, 0, 1, 2})     // 期望 [1 2 0 0]
	run([]int{1, 2, 0, 0})     // 期望 [1 2 0 0]（本来就在后面，不该被改动）
	run([]int{0, 0, 0, 0})     // 期望 [0 0 0 0]（全零也要能走完，别提前返回）
}
