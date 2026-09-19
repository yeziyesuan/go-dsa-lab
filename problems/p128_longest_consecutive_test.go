package problems

import "testing"

func TestLongestConsecutive(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{name: "示例1_两段序列", nums: []int{100, 4, 200, 1, 3, 2}, want: 4},
		{name: "示例2_含重复元素", nums: []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}, want: 9},
		{name: "空切片", nums: []int{}, want: 0},
		{name: "nil切片", nums: nil, want: 0},
		{name: "单元素", nums: []int{5}, want: 1},
		{name: "全相同", nums: []int{5, 5, 5, 5}, want: 1},
		{name: "两元素连续", nums: []int{2, 1}, want: 2},
		{name: "两元素不连续", nums: []int{2, 4}, want: 1},
		{name: "负数连续到正数", nums: []int{-2, -1, 0, 1}, want: 4},
		{name: "负数离散", nums: []int{-5, -1, 3}, want: 1},
		{name: "重复打散顺序", nums: []int{1, 2, 0, 1}, want: 3},
		{name: "两段等长取相等", nums: []int{1, 2, 3, 10, 11, 12}, want: 3},
		{name: "乱序长段", nums: []int{9, 1, 4, 7, 3, -1, 0, 5, 8, -1, 6}, want: 7},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LongestConsecutive(tt.nums); got != tt.want {
				t.Errorf("LongestConsecutive(%v) = %d, want %d", tt.nums, got, tt.want)
			}
		})
	}
}
