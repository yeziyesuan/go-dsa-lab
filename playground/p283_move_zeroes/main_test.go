// 283 的表驱动测试：实现 moveZeroes 之前它是红的。
//
// 跑：go test ./playground/p283_move_zeroes -v
package main

import (
	"reflect"
	"testing"
)

func TestMoveZeroes(t *testing.T) {
	cases := []struct {
		name string
		nums []int
		want []int
	}{
		{"示例1", []int{0, 1, 0, 3, 12}, []int{1, 3, 12, 0, 0}},
		{"示例2 单元素零", []int{0}, []int{0}},
		{"单元素非零", []int{7}, []int{7}},
		{"全零", []int{0, 0, 0, 0}, []int{0, 0, 0, 0}},
		{"无零", []int{1, 2, 3}, []int{1, 2, 3}},
		{"零全在前", []int{0, 0, 1, 2}, []int{1, 2, 0, 0}},
		{"零全在后", []int{1, 2, 0, 0}, []int{1, 2, 0, 0}},
		{"含负数", []int{-1, 0, -2, 0, 3}, []int{-1, -2, 3, 0, 0}},
		{"交替零与非零", []int{1, 0, 2, 0, 3, 0}, []int{1, 2, 3, 0, 0, 0}},
		{"空数组", []int{}, []int{}},
		{"nil 切片", nil, []int{}},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			nums := append([]int(nil), c.nums...) // 传副本：moveZeroes 会原地改
			moveZeroes(nums)

			// nil 和空切片统一按空切片比较（append 出来可能是 nil）。
			want := c.want
			if want == nil {
				want = []int{}
			}
			if nums == nil {
				nums = []int{}
			}
			if !reflect.DeepEqual(nums, want) {
				t.Fatalf("moveZeroes(%v) 之后得到 %v，期望 %v", c.nums, nums, want)
			}
		})
	}
}

// 单独盯「相对顺序」这个最容易写坏的约束：非零元素的先后顺序不能变。
func TestMoveZeroesKeepsOrder(t *testing.T) {
	nums := []int{0, 5, 0, 1, 0, -3, 0, 0, 4}
	moveZeroes(nums)

	got := make([]int, 0, len(nums))
	for _, v := range nums {
		if v != 0 {
			got = append(got, v)
		}
	}
	want := []int{5, 1, -3, 4}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("非零元素顺序为 %v，期望 %v", got, want)
	}
}
