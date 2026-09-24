// 121. 买卖股票的最佳时机（LeetCode 121 · Easy）
//
// 题干：给定数组 prices，prices[i] 表示第 i 天的股价。
// 你只能选「某一天买入、之后的某一天卖出」各一次（也可以不交易），
// 求能获得的最大利润；不获利就返回 0。
//
// 示例1：prices = [7,1,5,3,6,4] → 5   （第 2 天买 1，第 5 天卖 6）
// 示例2：prices = [7,6,4,3,1]   → 0   （一路下跌，不交易）
//
// 运行：go run ./playground/p121_max_profit
// 测试：go test ./playground/p121_max_profit -v
// 约定：先纸上写完整解法（package / import / 函数签名 / 返回值都写全），
// 再凭记忆敲进来；敲完按「纸面自检四问」过一遍：变量 → 边界 → 返回值 → 复杂度。
package main

import (
	"fmt"
)

// maxProfit 返回一次买卖能获得的最大利润。
//
// 思路（一句话）：一趟遍历，用 val 记住见过的最低价、用 num 记住最大利润；不额外开容器
// 复杂度：时间 O(n)（单层遍历 n 次），空间 O(1)（val / num / maxval 三个常数变量）
// ⟵ 助手修正了两处笔误：① 并没有"创建切片"，prices 是参数，直接遍历它；
//
//	② 空间不是 O(n)——你没有新开跟 n 同规模的容器，只有三个标量，所以是 O(1)
//
// TODO 先写暴力（O(n²)）跑通示例，再想能不能一趟扫完（O(n)）——两种都留着对比
func maxProfit(prices []int) int {
	// TODO 你的解法
	if len(prices) == 0 {
		return 0
	} //纠错
	val := prices[0]
	var num int
	for _, v := range prices {
		maxval := v - val
		if maxval > num {
			num = maxval
		}
		if v < val {
			val = v
		}
	}
	return num // ⟵ 助手改：原来是 `return 0`，把循环里辛苦算出来的 num 扔掉了（这题第一次检查就是这个错）
}

func main() {
	// 官方示例：先让这两个全对，再自己加边界用例（空数组 / 只有一个元素 / 一路涨）
	fmt.Println("示例1:", maxProfit([]int{7, 1, 5, 3, 6, 4}), "期望 5")
	fmt.Println("示例2:", maxProfit([]int{7, 6, 4, 3, 1}), "期望 0")

	// 自己加的边界（写完把期望值填对，再跑一次）
	fmt.Println("单元素:", maxProfit([]int{5}), "期望 0")
	fmt.Println("一路涨:", maxProfit([]int{1, 2, 3, 4, 5}), "期望 4")
}
