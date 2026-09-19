package problems

import "testing"

func TestMaxArea(t *testing.T) {
	tests := []struct {
		name   string
		height []int
		want   int
	}{
		{name: "题目示例1", height: []int{1, 8, 6, 2, 5, 4, 8, 3, 7}, want: 49},
		{name: "题目示例2", height: []int{1, 1}, want: 1},
		{name: "空数组", height: []int{}, want: 0},
		{name: "nil切片", height: nil, want: 0},
		{name: "单元素_无容器", height: []int{5}, want: 0},
		{name: "两元素_含零", height: []int{0, 5}, want: 0},
		{name: "两元素_相等", height: []int{4, 4}, want: 4},
		{name: "全相同", height: []int{3, 3, 3, 3}, want: 9},
		{name: "全零", height: []int{0, 0, 0}, want: 0},
		{name: "全零加一个高柱", height: []int{0, 0, 10, 0}, want: 0},
		{name: "递增_取两端", height: []int{1, 2, 3, 4, 5}, want: 6},
		{name: "递减_取两端", height: []int{5, 4, 3, 2, 1}, want: 6},
		{name: "最高柱在中间", height: []int{1, 100, 1}, want: 2},
		{name: "两侧最高且最宽", height: []int{100, 1, 1, 1, 100}, want: 400},
		{name: "单侧最高_另一边为零", height: []int{100, 0, 0, 0}, want: 0},
		{name: "含零的常规用例", height: []int{2, 0, 2}, want: 4},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxArea(tt.height); got != tt.want {
				t.Fatalf("MaxArea(%v) = %d，期望 %d", tt.height, got, tt.want)
			}
		})
	}
}

// 用 O(n^2) 暴力解做交叉验证，确保双指针没漏解（小规模随机 + 固定数据）。
func TestMaxAreaAgainstBruteForce(t *testing.T) {
	cases := [][]int{
		{1, 8, 6, 2, 5, 4, 8, 3, 7},
		{1, 2, 1},
		{2, 3, 4, 5, 18, 17, 6},
		{1, 1, 1, 1, 1},
		{0, 1, 0, 2, 1, 0, 1, 3},
		{9, 1, 9, 1, 9},
	}

	brute := func(h []int) int {
		best := 0
		for i := 0; i < len(h); i++ {
			for j := i + 1; j < len(h); j++ {
				m := h[i]
				if h[j] < m {
					m = h[j]
				}
				if a := m * (j - i); a > best {
					best = a
				}
			}
		}
		return best
	}

	for _, h := range cases {
		t.Run("", func(t *testing.T) {
			if got, want := MaxArea(h), brute(h); got != want {
				t.Fatalf("MaxArea(%v) = %d，暴力解为 %d", h, got, want)
			}
		})
	}
}
