package problems

import "testing"

func TestMaxSubArray(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{name: "题目示例1", nums: []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, want: 6},
		{name: "题目示例2_单元素", nums: []int{1}, want: 1},
		{name: "题目示例3_全正", nums: []int{5, 4, -1, 7, 8}, want: 23},
		{name: "空数组", nums: []int{}, want: 0},
		{name: "nil切片", nums: nil, want: 0},
		{name: "单元素负数", nums: []int{-1}, want: -1},
		{name: "全负数", nums: []int{-3, -2, -5}, want: -2},
		{name: "全负数_两个", nums: []int{-2, -1}, want: -1},
		{name: "全相同正数", nums: []int{3, 3, 3}, want: 9},
		{name: "全相同负数", nums: []int{-4, -4, -4}, want: -4},
		{name: "全零", nums: []int{0, 0, 0}, want: 0},
		{name: "含零的负数数组", nums: []int{-2, 0, -3}, want: 0},
		{name: "最大子数组在中间", nums: []int{-1, -1, 10, -1, -1}, want: 10},
		{name: "最大子数组在前缀", nums: []int{1, 2, 3, -10}, want: 6},
		{name: "最大子数组在后缀", nums: []int{-10, 1, 2, 3}, want: 6},
		{name: "首尾各一段_取较大", nums: []int{4, -1, -1, -1, 3}, want: 4},
		{name: "正负交替", nums: []int{-2, 1, -3, 4, -1, 2, 1, -5, 4, 100}, want: 105},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxSubArray(tt.nums); got != tt.want {
				t.Fatalf("MaxSubArray(%v) = %d，期望 %d", tt.nums, got, tt.want)
			}
		})
	}
}
