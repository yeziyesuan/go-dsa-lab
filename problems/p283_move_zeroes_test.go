package problems

import (
	"reflect"
	"testing"
)

func TestMoveZeroes(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{name: "题目示例1", nums: []int{0, 1, 0, 3, 12}, want: []int{1, 3, 12, 0, 0}},
		{name: "题目示例2_单元素零", nums: []int{0}, want: []int{0}},
		{name: "空数组", nums: []int{}, want: []int{}},
		{name: "nil切片", nums: nil, want: []int{}},
		{name: "单元素非零", nums: []int{7}, want: []int{7}},
		{name: "全零", nums: []int{0, 0, 0, 0}, want: []int{0, 0, 0, 0}},
		{name: "无零", nums: []int{1, 2, 3}, want: []int{1, 2, 3}},
		{name: "零全在前", nums: []int{0, 0, 1, 2}, want: []int{1, 2, 0, 0}},
		{name: "零全在后_应保持不变", nums: []int{1, 2, 0, 0}, want: []int{1, 2, 0, 0}},
		{name: "含负数", nums: []int{-1, 0, -2, 0, 3}, want: []int{-1, -2, 3, 0, 0}},
		{name: "全负数无零", nums: []int{-2, -1}, want: []int{-2, -1}},
		{name: "交替零与非零", nums: []int{1, 0, 2, 0, 3, 0}, want: []int{1, 2, 3, 0, 0, 0}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nums := append([]int(nil), tt.nums...)
			MoveZeroes(nums)

			// want 为 nil 时统一按空切片比较。
			want := tt.want
			if want == nil {
				want = []int{}
			}
			if nums == nil {
				nums = []int{}
			}
			if !reflect.DeepEqual(nums, want) {
				t.Fatalf("MoveZeroes(%v) 之后得到 %v，期望 %v", tt.nums, nums, want)
			}
		})
	}
}

// 单独校验「相对顺序」这个容易写坏的约束：非零元素必须按原顺序出现。
func TestMoveZeroesKeepsOrder(t *testing.T) {
	nums := []int{0, 5, 0, 1, 0, -3, 0, 0, 4}
	MoveZeroes(nums)

	nonzero := make([]int, 0, len(nums))
	for _, v := range nums {
		if v != 0 {
			nonzero = append(nonzero, v)
		}
	}
	want := []int{5, 1, -3, 4}
	if !reflect.DeepEqual(nonzero, want) {
		t.Fatalf("非零元素顺序为 %v，期望 %v", nonzero, want)
	}
}
