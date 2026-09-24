// 121 的表驱动测试：实现之前它会红，实现对了就全绿。
//
// 跑：go test ./playground/p121_max_profit -v
package main

import "testing"

func TestMaxProfit(t *testing.T) {
	cases := []struct {
		name   string
		prices []int
		want   int
	}{
		{"示例1 先跌后涨", []int{7, 1, 5, 3, 6, 4}, 5},
		{"示例2 一路下跌", []int{7, 6, 4, 3, 1}, 0},
		{"只有一个元素", []int{5}, 0},
		{"两天上涨", []int{1, 2}, 1},
		{"最低价出现在最后", []int{2, 4, 1}, 2}, // 别写成「全局最低价 → 之后最高价」
		{"一路涨", []int{1, 2, 3, 4, 5}, 4},
		{"中间有平盘", []int{3, 3, 5, 0, 0, 3, 1, 4}, 4},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if got := maxProfit(c.prices); got != c.want {
				t.Errorf("maxProfit(%v) = %d, want %d", c.prices, got, c.want)
			}
		})
	}
}
