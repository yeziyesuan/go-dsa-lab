package problems

import "testing"

func TestMaxProfit(t *testing.T) {
	tests := []struct {
		name   string
		prices []int
		want   int
	}{
		{name: "题目示例1", prices: []int{7, 1, 5, 3, 6, 4}, want: 5},
		{name: "题目示例2_单边下跌", prices: []int{7, 6, 4, 3, 1}, want: 0},
		{name: "空数组", prices: []int{}, want: 0},
		{name: "nil切片", prices: nil, want: 0},
		{name: "单元素", prices: []int{5}, want: 0},
		{name: "全相同", prices: []int{3, 3, 3}, want: 0},
		{name: "全零", prices: []int{0, 0, 0}, want: 0},
		{name: "递增_买第一天", prices: []int{1, 2, 3, 4, 5}, want: 4},
		{name: "最低价在最后", prices: []int{5, 4, 3, 2, 1}, want: 0},
		{name: "最低价在中间", prices: []int{5, 1, 5}, want: 4},
		{name: "最大价在最低价之前", prices: []int{10, 1, 2}, want: 1},
		{name: "含零价格", prices: []int{2, 0, 5}, want: 5},
		{name: "两天_持平", prices: []int{2, 2}, want: 0},
		{name: "两天_上涨", prices: []int{2, 9}, want: 7},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxProfit(tt.prices); got != tt.want {
				t.Fatalf("MaxProfit(%v) = %d，期望 %d", tt.prices, got, tt.want)
			}
		})
	}
}
