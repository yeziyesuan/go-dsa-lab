package problems

import (
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
		head []int // 期望的 nums[:k] 内容
	}{
		{name: "题目示例1", nums: []int{1, 1, 2}, want: 2, head: []int{1, 2}},
		{name: "题目示例2", nums: []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, want: 5, head: []int{0, 1, 2, 3, 4}},
		{name: "空数组", nums: []int{}, want: 0, head: []int{}},
		{name: "nil切片", nums: nil, want: 0, head: []int{}},
		{name: "单元素", nums: []int{7}, want: 1, head: []int{7}},
		{name: "全相同", nums: []int{5, 5, 5, 5}, want: 1, head: []int{5}},
		{name: "无重复", nums: []int{1, 2, 3}, want: 3, head: []int{1, 2, 3}},
		{name: "含负数", nums: []int{-3, -3, -1, 0, 0}, want: 3, head: []int{-3, -1, 0}},
		{name: "全负数且全相同", nums: []int{-2, -2}, want: 1, head: []int{-2}},
		{name: "重复集中在尾部", nums: []int{1, 2, 2, 2}, want: 2, head: []int{1, 2}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nums := append([]int(nil), tt.nums...)

			if got := RemoveDuplicates(nums); got != tt.want {
				t.Fatalf("RemoveDuplicates(%v) = %d，期望 %d", tt.nums, got, tt.want)
			}

			// 返回值 k 必须同时是「有效前缀长度」和「前缀内容正确」。
			// 注意：空输入时 nums[:0] 是 nil 切片，和 []int{} 用 DeepEqual 比会被判不等，
			// 所以这里按长度 + 逐元素比较，不区分 nil 与空切片。
			head := nums[:tt.want]
			if len(head) != len(tt.head) {
				t.Fatalf("RemoveDuplicates(%v) 之后 nums[:%d] 长度 = %d，期望 %d",
					tt.nums, tt.want, len(head), len(tt.head))
			}
			for i := range head {
				if head[i] != tt.head[i] {
					t.Fatalf("RemoveDuplicates(%v) 之后 nums[:%d] = %v，期望 %v",
						tt.nums, tt.want, head, tt.head)
				}
			}
		})
	}
}

// 单独盯一下空输入：切片不能越界，也不能依赖调用方传入非 nil 切片。
func TestRemoveDuplicatesEmpty(t *testing.T) {
	if got := RemoveDuplicates(nil); got != 0 {
		t.Fatalf("RemoveDuplicates(nil) = %d，期望 0", got)
	}
	if got := RemoveDuplicates([]int{}); got != 0 {
		t.Fatalf("RemoveDuplicates([]int{}) = %d，期望 0", got)
	}
}
