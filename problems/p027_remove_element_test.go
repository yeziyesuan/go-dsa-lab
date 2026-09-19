package problems

import (
	"reflect"
	"sort"
	"testing"
)

func TestRemoveElement(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		val  int
		want int
		head []int // 期望的 nums[:k]，顺序不限，比较前排序
	}{
		{name: "题目示例1", nums: []int{3, 2, 2, 3}, val: 3, want: 2, head: []int{2, 2}},
		{name: "题目示例2", nums: []int{0, 1, 2, 2, 3, 0, 4, 2}, val: 2, want: 5, head: []int{0, 0, 1, 3, 4}},
		{name: "空数组", nums: []int{}, val: 1, want: 0, head: []int{}},
		{name: "nil切片", nums: nil, val: 1, want: 0, head: []int{}},
		{name: "单元素_命中", nums: []int{5}, val: 5, want: 0, head: []int{}},
		{name: "单元素_未命中", nums: []int{5}, val: 3, want: 1, head: []int{5}},
		{name: "全等于val", nums: []int{2, 2, 2, 2}, val: 2, want: 0, head: []int{}},
		{name: "全不等于val", nums: []int{1, 2, 3}, val: 9, want: 3, head: []int{1, 2, 3}},
		{name: "含负数", nums: []int{-1, -2, -1, 0}, val: -1, want: 2, head: []int{-2, 0}},
		{name: "全负数全命中", nums: []int{-7, -7}, val: -7, want: 0, head: []int{}},
		{name: "val为0且数组全零", nums: []int{0, 0, 0}, val: 0, want: 0, head: []int{}},
		{name: "目标在首尾", nums: []int{4, 1, 2, 3, 4}, val: 4, want: 3, head: []int{1, 2, 3}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nums := append([]int(nil), tt.nums...)

			if got := RemoveElement(nums, tt.val); got != tt.want {
				t.Fatalf("RemoveElement(%v, %d) = %d，期望 %d", tt.nums, tt.val, got, tt.want)
			}

			head := append([]int(nil), nums[:tt.want]...)
			sort.Ints(head)
			want := append([]int(nil), tt.head...)
			sort.Ints(want)

			if !reflect.DeepEqual(head, want) {
				t.Fatalf("RemoveElement(%v, %d) 之后 nums[:%d] 排序 = %v，期望 %v",
					tt.nums, tt.val, tt.want, head, want)
			}

			// 前缀里不允许残留 val。
			for i := 0; i < tt.want; i++ {
				if nums[i] == tt.val {
					t.Fatalf("RemoveElement(%v, %d) 之后 nums[%d] 仍等于 val", tt.nums, tt.val, i)
				}
			}
		})
	}
}
