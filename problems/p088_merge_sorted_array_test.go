package problems

import (
	"reflect"
	"testing"
)

func TestMerge(t *testing.T) {
	tests := []struct {
		name  string
		nums1 []int // 长度必须为 m+n
		m     int
		nums2 []int
		n     int
		want  []int // nums1 合并后的整体内容
	}{
		{
			name:  "题目示例1",
			nums1: []int{1, 2, 3, 0, 0, 0}, m: 3,
			nums2: []int{2, 5, 6}, n: 3,
			want: []int{1, 2, 2, 3, 5, 6},
		},
		{
			name:  "题目示例2_nums2为空",
			nums1: []int{1}, m: 1,
			nums2: []int{}, n: 0,
			want: []int{1},
		},
		{
			name:  "题目示例3_nums1有效段为空",
			nums1: []int{0}, m: 0,
			nums2: []int{1}, n: 1,
			want: []int{1},
		},
		{
			name:  "两个都空",
			nums1: []int{}, m: 0,
			nums2: []int{}, n: 0,
			want: []int{},
		},
		{
			name:  "单元素_各一个",
			nums1: []int{1, 0}, m: 1,
			nums2: []int{2}, n: 1,
			want: []int{1, 2},
		},
		{
			name:  "单元素_nums2更小",
			nums1: []int{2, 0}, m: 1,
			nums2: []int{1}, n: 1,
			want: []int{1, 2},
		},
		{
			name:  "含负数",
			nums1: []int{-3, -1, 0, 0}, m: 2,
			nums2: []int{-2, 0}, n: 2,
			want: []int{-3, -2, -1, 0},
		},
		{
			name:  "全负数",
			nums1: []int{-5, -1, 0}, m: 2,
			nums2: []int{-4}, n: 1,
			want: []int{-5, -4, -1},
		},
		{
			name:  "全相同",
			nums1: []int{2, 2, 0, 0}, m: 2,
			nums2: []int{2, 2}, n: 2,
			want: []int{2, 2, 2, 2},
		},
		{
			name:  "nums1全部大于nums2",
			nums1: []int{4, 5, 6, 0, 0, 0}, m: 3,
			nums2: []int{1, 2, 3}, n: 3,
			want: []int{1, 2, 3, 4, 5, 6},
		},
		{
			name:  "nums1全部小于nums2",
			nums1: []int{1, 2, 3, 0, 0, 0}, m: 3,
			nums2: []int{4, 5, 6}, n: 3,
			want: []int{1, 2, 3, 4, 5, 6},
		},
		{
			name:  "含零值元素不与占位符混淆",
			nums1: []int{0, 0, 0, 0}, m: 2,
			nums2: []int{0, 0}, n: 2,
			want: []int{0, 0, 0, 0},
		},
		{
			name:  "规模不对称_m远大于n",
			nums1: []int{1, 3, 5, 7, 9, 0}, m: 5,
			nums2: []int{2}, n: 1,
			want: []int{1, 2, 3, 5, 7, 9},
		},
		{
			name:  "规模不对称_n远大于m",
			nums1: []int{5, 0, 0, 0, 0}, m: 1,
			nums2: []int{1, 2, 3, 4}, n: 4,
			want: []int{1, 2, 3, 4, 5},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nums1 := append([]int(nil), tt.nums1...)
			Merge(nums1, tt.m, tt.nums2, tt.n)

			want := tt.want
			if want == nil {
				want = []int{}
			}
			if nums1 == nil {
				nums1 = []int{}
			}
			// 断言 nums1 整体（长度 m+n），不只是有效段。
			if !reflect.DeepEqual(nums1, want) {
				t.Fatalf("Merge(%v, %d, %v, %d) 之后 nums1 = %v，期望 %v",
					tt.nums1, tt.m, tt.nums2, tt.n, nums1, want)
			}
		})
	}
}

// 越界参数不应 panic：函数内部有防线，直接原样返回。
func TestMergeInvalidArgs(t *testing.T) {
	nums1 := []int{1, 2, 3}
	Merge(nums1, 3, []int{4}, 1) // m+n > len(nums1)
	if !reflect.DeepEqual(nums1, []int{1, 2, 3}) {
		t.Fatalf("非法参数下 nums1 被修改为 %v，期望保持不变", nums1)
	}

	nums2 := []int{9}
	Merge(nums1, 1, nums2, 5) // n > len(nums2)
	if !reflect.DeepEqual(nums1, []int{1, 2, 3}) {
		t.Fatalf("非法参数下 nums1 被修改为 %v，期望保持不变", nums1)
	}
}
