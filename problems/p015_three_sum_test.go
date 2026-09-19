package problems

import (
	"reflect"
	"sort"
	"testing"
)

// 归一化：[[]int] 的每个三元组内部排序，再把三元组之间按字典序排序，
// 这样结果顺序就不影响比较了。
func normalizeTriplets(xs [][]int) [][]int {
	out := make([][]int, 0, len(xs))
	for _, t := range xs {
		c := append([]int(nil), t...)
		sort.Ints(c)
		out = append(out, c)
	}
	sort.Slice(out, func(i, j int) bool {
		a, b := out[i], out[j]
		for k := 0; k < len(a) && k < len(b); k++ {
			if a[k] != b[k] {
				return a[k] < b[k]
			}
		}
		return len(a) < len(b)
	})
	return out
}

func TestThreeSum(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want [][]int
	}{
		{
			name: "题目示例1",
			nums: []int{-1, 0, 1, 2, -1, -4},
			want: [][]int{{-1, -1, 2}, {-1, 0, 1}},
		},
		{
			name: "题目示例2_无解",
			nums: []int{0, 1, 1},
			want: nil,
		},
		{
			name: "题目示例3_全零",
			nums: []int{0, 0, 0},
			want: [][]int{{0, 0, 0}},
		},
		{name: "空数组", nums: []int{}, want: nil},
		{name: "nil切片", nums: nil, want: nil},
		{name: "单元素", nums: []int{0}, want: nil},
		{name: "两元素", nums: []int{-1, 1}, want: nil},
		{
			name: "重复元素多组去重",
			nums: []int{-2, 0, 0, 2, 2},
			want: [][]int{{-2, 0, 2}},
		},
		{
			name: "全正数_无解",
			nums: []int{1, 2, 3, 4},
			want: nil,
		},
		{
			name: "全负数_无解",
			nums: []int{-1, -2, -3},
			want: nil,
		},
		{
			name: "多个零只出一组",
			nums: []int{0, 0, 0, 0},
			want: [][]int{{0, 0, 0}},
		},
		{
			name: "含大量重复",
			nums: []int{-1, -1, -1, 2, 2, 2},
			want: [][]int{{-1, -1, 2}},
		},
		{
			name: "多个不同解",
			nums: []int{-4, -2, -2, -2, 0, 1, 2, 2, 2, 3, 3, 4, 4, 6, 6},
			want: [][]int{{-4, -2, 6}, {-4, 0, 4}, {-4, 1, 3}, {-4, 2, 2}, {-2, -2, 4}, {-2, 0, 2}},
		},
		{
			name: "正负对称含零",
			nums: []int{-1, 0, 1, -1, 0, 1},
			want: [][]int{{-1, 0, 1}},
		},
		{
			name: "排序后才可能出现的边界",
			nums: []int{3, 0, -2, -1, 1, 2},
			want: [][]int{{-2, 0, 2}, {-2, -1, 3}, {-1, 0, 1}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nums := append([]int(nil), tt.nums...)

			got := normalizeTriplets(ThreeSum(nums))
			want := normalizeTriplets(tt.want)

			if len(got) == 0 && len(want) == 0 {
				return // nil 与空切片视为等价
			}
			if !reflect.DeepEqual(got, want) {
				t.Fatalf("ThreeSum(%v) = %v（归一化后 %v），期望 %v（归一化后 %v）",
					tt.nums, ThreeSum(append([]int(nil), tt.nums...)), got, tt.want, want)
			}

			// 额外校验：每个三元组确实和为 0，且没有重复三元组。
			for _, tp := range got {
				if len(tp) != 3 {
					t.Fatalf("三元组 %v 长度不为 3", tp)
				}
				if s := tp[0] + tp[1] + tp[2]; s != 0 {
					t.Fatalf("三元组 %v 的和为 %d，期望 0", tp, s)
				}
			}
			for i := 1; i < len(got); i++ {
				if reflect.DeepEqual(got[i], got[i-1]) {
					t.Fatalf("结果里出现重复三元组 %v", got[i])
				}
			}
		})
	}
}
