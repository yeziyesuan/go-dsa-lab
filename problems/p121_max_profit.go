package problems

// 121. 买卖股票的最佳时机（Best Time to Buy and Sell Stock）
//
// 题目：
//
//	给定数组 prices，prices[i] 表示第 i 天股票的价格。你只能选择某一天买入，
//	并在未来的某一天卖出，返回能获得的最大利润；如果无法获利，返回 0。
//
// 思路：
//
//	一次遍历维护「历史最低价」。利润 = 卖出价 - 买入价，且买入必须在卖出之前，
//	所以顺序扫一遍：先用今天价格和历史最低价算一次「今天卖能赚多少」并更新答案，
//	再用今天价格去更新历史最低价。先算利润后更新最低价，就天然保证了
//	「买入日 <= 卖出日」（同一天买入卖出利润为 0，也不影响）。
//
// 复杂度：
//
//	时间 O(n)：只扫一趟，每天做两次常数比较。
//	空间 O(1)：只有 low 和 profit 两个变量。
//
// 卡点：
//
//  1. 空数组（和 nil 切片）没有挡掉：low 初始化成 prices[0] 会直接越界 panic。
//     题目要求空数组返回 0，先 `if len(prices) == 0 { return 0 }`。
//  2. 两句的顺序写反：如果先更新 low 再算利润，当天就会出现「用今天价格买入、
//     又用今天价格卖出」甚至「用今天的最低价倒推利润」的错误结果。
//  3. low 的初值给成 0 或 math.MaxInt 都不对——0 会让价格全为正时利润恒为
//     price 本身；正确写法是初始化为 prices[0]。
//  4. 只想赚最少 0：当 maxProfit 是负数（价格一路下跌）时要返回 0，
//     因为「不交易」也是允许的选择。用 profit 初值 0 即可。
//
// Go 注意点：
//
//   - 这题的 O(1) 空间靠的是「遍历中维护状态」，不需要任何辅助切片/数组；
//     不要为了可读性去开一个 minPrices []int 把空间写成 O(n)。
//   - 不要用 math.MaxInt 之类的哨兵值初始化 low：题面保证 0 <= prices[i]，
//     但直接拿 prices[0] 更省心，也少一个 import。
//   - 如果 prices 可能为空且函数不能提前 return，就该用 `prices[0]` 之外的
//     写法配合 ok 判断；这里显式 return 0 最直白。
func MaxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	low := prices[0]
	profit := 0
	for _, p := range prices {
		if p-low > profit {
			profit = p - low
		}
		if p < low {
			low = p
		}
	}
	return profit
}
