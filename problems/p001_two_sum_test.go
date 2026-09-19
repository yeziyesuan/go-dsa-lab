package problems

import (
	"reflect"
	"sort"
	"testing"
)

// 两数之和的下标顺序不唯一，统一排序后再比较。
func sortPair(p []int) []int {
	q := append([]int(nil), p...)
	sort.Ints(q)
	return q
}

func TestTwoSum(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   []int // 期望下标对，比较前会排序
	}{
		{name: "题目示例1", nums: []int{2, 7, 11, 15}, target: 9, want: []int{0, 1}},
		{name: "题目示例2", nums: []int{3, 2, 4}, target: 6, want: []int{1, 2}},
		{name: "题目示例3_相同元素下自成一对", nums: []int{3, 3}, target: 6, want: []int{0, 1}},
		{name: "含负数_目标为零", nums: []int{-3, 4, 3, 90}, target: 0, want: []int{0, 2}},
		{name: "答案落在数组末尾", nums: []int{1, 2, 3, 4, 5}, target: 9, want: []int{3, 4}},
		{name: "答案含下标零", nums: []int{0, 4, 3, 0}, target: 0, want: []int{0, 3}},
		{name: "单元素_无解", nums: []int{5}, target: 5, want: nil},
		{name: "空数组_无解", nums: []int{}, target: 0, want: nil},
		{name: "nil切片_无解", nums: nil, target: 0, want: nil},
		{name: "全相同_无解", nums: []int{1, 1, 1}, target: 7, want: nil},
		{name: "全负数", nums: []int{-1, -2, -3, -4}, target: -7, want: []int{2, 3}},
		{name: "全是零_取前两个零", nums: []int{0, 0, 0}, target: 0, want: []int{0, 1}},
		{name: "只有一个零_不能用两次", nums: []int{0}, target: 0, want: nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := TwoSum(tt.nums, tt.target)

			if tt.want == nil {
				// 题目保证有解，无解用例只用于检查不 panic、返回长度不为 2。
				if len(got) == 2 {
					t.Fatalf("TwoSum(%v, %d) = %v，期望无解（长度 != 2）", tt.nums, tt.target, got)
				}
				return
			}

			if len(got) != 2 {
				t.Fatalf("TwoSum(%v, %d) = %v，期望长度 2 的结果 %v", tt.nums, tt.target, got, tt.want)
			}
			if !reflect.DeepEqual(sortPair(got), sortPair(tt.want)) {
				t.Fatalf("TwoSum(%v, %d) = %v，期望（顺序无关）%v", tt.nums, tt.target, got, tt.want)
			}
			// 顺手校验返回下标真的能凑出 target，防止只对比形状。
			if sum := tt.nums[got[0]] + tt.nums[got[1]]; sum != tt.target {
				t.Fatalf("TwoSum(%v, %d) = %v，但两数之和为 %d", tt.nums, tt.target, got, sum)
			}
		})
	}
}
