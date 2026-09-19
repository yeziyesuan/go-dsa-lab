package problems

import "testing"

func TestContainsDuplicate(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want bool
	}{
		{name: "示例1_有重复", nums: []int{1, 2, 3, 1}, want: true},
		{name: "示例2_互不相同", nums: []int{1, 2, 3, 4}, want: false},
		{name: "示例3_多组重复", nums: []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}, want: true},
		{name: "空切片", nums: []int{}, want: false},
		{name: "nil切片", nums: nil, want: false},
		{name: "单元素", nums: []int{7}, want: false},
		{name: "两元素相同", nums: []int{0, 0}, want: true},
		{name: "全相同", nums: []int{2, 2, 2, 2}, want: true},
		{name: "含负数的重复", nums: []int{-3, 0, 5, -3}, want: true},
		{name: "含负数不重复", nums: []int{-3, 0, 5, -4}, want: false},
		{name: "零与负数掺在一起", nums: []int{0, -1, 1, 0}, want: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ContainsDuplicate(tt.nums); got != tt.want {
				t.Errorf("ContainsDuplicate(%v) = %v, want %v", tt.nums, got, tt.want)
			}
		})
	}
}
