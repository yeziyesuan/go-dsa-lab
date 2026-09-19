package problems

import "testing"

func TestTrap(t *testing.T) {
	tests := []struct {
		name   string
		height []int
		want   int
	}{
		{name: "题目示例1", height: []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}, want: 6},
		{name: "题目示例2", height: []int{4, 2, 0, 3, 2, 5}, want: 9},
		{name: "空数组", height: []int{}, want: 0},
		{name: "nil切片", height: nil, want: 0},
		{name: "单元素", height: []int{5}, want: 0},
		{name: "两元素_接不到水", height: []int{1, 2}, want: 0},
		{name: "全相同", height: []int{3, 3, 3, 3}, want: 0},
		{name: "全零", height: []int{0, 0, 0}, want: 0},
		{name: "单调递增_接不到水", height: []int{1, 2, 3, 4}, want: 0},
		{name: "单调递减_接不到水", height: []int{4, 3, 2, 1}, want: 0},
		{name: "单个凹槽", height: []int{2, 0, 2}, want: 2},
		{name: "深凹槽", height: []int{5, 0, 0, 0, 5}, want: 15},
		{name: "等高三段中间凹陷", height: []int{3, 0, 3, 0, 3}, want: 6},
		{name: "含零起止", height: []int{0, 2, 0, 2, 0}, want: 2},
		{name: "阶梯状", height: []int{0, 1, 2, 1, 0, 1, 2, 1, 0}, want: 4},
		{name: "左侧高右侧低_只有矮墙决定", height: []int{5, 1, 1, 1, 2}, want: 3},
		{name: "右侧高左侧低", height: []int{2, 1, 1, 1, 5}, want: 3},
		{name: "大坑套小坑", height: []int{6, 4, 2, 0, 3, 2, 0, 3, 1, 5}, want: 25},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Trap(tt.height); got != tt.want {
				t.Fatalf("Trap(%v) = %d，期望 %d", tt.height, got, tt.want)
			}
		})
	}
}

// 交叉验证：用「预处理 leftMax/rightMax 数组」的 O(n) 空间朴素解对拍，
// 确认双指针版没算错任何一个位置。
func TestTrapAgainstPrefixArrays(t *testing.T) {
	cases := [][]int{
		{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1},
		{4, 2, 0, 3, 2, 5},
		{2, 0, 2},
		{5, 0, 0, 0, 5},
		{6, 4, 2, 0, 3, 2, 0, 3, 1, 5},
		{1, 0, 1},
		{0, 0, 0},
		{3, 1, 2, 1, 3},
	}

	naive := func(h []int) int {
		n := len(h)
		if n == 0 {
			return 0
		}
		left := make([]int, n)
		right := make([]int, n)
		left[0] = h[0]
		for i := 1; i < n; i++ {
			left[i] = left[i-1]
			if h[i] > left[i] {
				left[i] = h[i]
			}
		}
		right[n-1] = h[n-1]
		for i := n - 2; i >= 0; i-- {
			right[i] = right[i+1]
			if h[i] > right[i] {
				right[i] = h[i]
			}
		}
		sum := 0
		for i := 0; i < n; i++ {
			m := left[i]
			if right[i] < m {
				m = right[i]
			}
			if m-h[i] > 0 {
				sum += m - h[i]
			}
		}
		return sum
	}

	for _, h := range cases {
		t.Run("", func(t *testing.T) {
			if got, want := Trap(h), naive(h); got != want {
				t.Fatalf("Trap(%v) = %d，朴素解为 %d", h, got, want)
			}
		})
	}
}
