package problems

// 42. 接雨水（Trapping Rain Water）
//
// 题目：
//
//	给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，
//	下雨之后能接多少雨水。
//
// 思路（本题用双指针 + leftMax/rightMax）：
//
//	某个位置 i 能接的水量 = min(它左边最高的柱子, 它右边最高的柱子) - height[i]，
//	前提是这个差值大于 0。也就是被「两侧最高墙里较矮的那面」夹出来的凹槽。
//
//	朴素做法是预处理出 leftMax[i] / rightMax[i] 两个数组，再累加，空间 O(n)。
//	双指针把这 O(n) 空间压成 O(1)：维护 l、r 两个指针和 leftMax、rightMax
//	两个变量（分别表示 [0, l] 和 [r, n-1] 区间的最大值）。
//
//	关键在于：当 height[l] < height[r] 时，虽然 leftMax 只统计了左边的真实
//	最大值，但右侧一定存在一个比 height[l] 更高的柱子（当前的 r，而且
//	rightMax >= height[r] > height[l] >= min(leftMax, 真实右侧最大值)）。
//	于是 min(leftMax, 真实右侧最高) == leftMax，位置 l 的水位就完全由
//	leftMax 决定，可以安全结算：leftMax - height[l]。然后 l++。
//	另一边对称。每一步都结算掉一个「水位已经确定」的位置，指针相遇即收工。
//
// 复杂度：
//
//	时间 O(n)：l 和 r 每轮必有一个移动，最多移动 n-1 次。
//	空间 O(1)：两个指针 + 两个最大高度变量（对比预处理法的 O(n) 数组）。
//
// 卡点：
//
//  1. 结算位置的顺序写反：必须在判断 height[l] < height[r] **之后**，
//     用对应那边的 leftMax/rightMax 去结算。如果先无条件累加
//     leftMax - height[l]，在左侧还没确定水位时就会多加（甚至加出负数）。
//  2. 忘记 `if w > 0` 或忘记 leftMax 单调不减：leftMax 要先 `if height[l] > leftMax`
//     更新，再算差值。空桶（height[l] >= leftMax）时差值 <= 0，加进去就是负数。
//     常见写法是 `leftMax = max(leftMax, height[l])` 然后 `res += leftMax - height[l]`，
//     因为 leftMax 已包含 height[l]，差值必然 >= 0，天然不用判断。
//  3. 指针移动条件写成 `height[l] <= height[r]` 时到底动哪边——两种写法都可以，
//     但必须保证「结算谁就动谁」，不能结算 l 却动了 r。
//  4. 空数组 / 单元素 / 全相同高度：结果都应为 0。前两个用例循环不进入；
//     全相同时 leftMax 始终等于 height[l]，差值为 0。
//  5. 高度里有 0：合法（0 就是空槽），比如 [0,1,0,2] 能接 1 单位。
//
// 单调栈解法（文字说明，W1 做不出来很正常）：
//
//	思路是「按行（横向）算水，而不是按列」。从左到右扫，维护一个**单调递减**
//	的栈，栈里存下标，对应的 height 从栈底到栈顶严格递减。
//
//	扫到下标 i 时，只要 height[i] 比栈顶对应的高度大，就说明栈顶位置是一个
//	「凹槽底部」：把栈顶 pop 出来记为 mid，此时新的栈顶（如果存在）就是凹槽
//	的**左墙**，而 i 是**右墙**。这个槽能接的水是
//	    (min(height[left], height[i]) - height[mid]) * (i - left - 1)
//	也就是「两侧墙里较矮的那面 减去 槽底高度」乘上「槽的宽度」。
//	高度差是横向那条水带的高度，宽度是左右墙之间的下标距离减一。
//	把所有 pop 出来的槽的水量累加就是答案。
//
//	两个解法的对照：双指针是「按列 / 竖向」累加，每个位置算一次；
//	单调栈是「按行 / 横向」累加，每个下标最多进出栈一次。
//	两者都是 O(n) 时间，但单调栈需要额外 O(n) 的栈空间，而且
//	「pop 之后新栈顶是左墙」这个不变量比双指针更绕，是典型的
//	「知道模板也想不起来怎么写」的题。第一周先把双指针版写熟，
//	单调栈留到后面刷「下一个更大元素」那一族题时一起补。
//
// Go 注意点：
//
//   - 用显式 `if a > b` 手写最大值，不引 math 包（Go 1.21+ 的内置 max
//     在这里可用，但 W1 统一手写，保证不同版本都能编译）。
//   - 双指针版零分配：没有 make、没有 append，只有两个局部变量，
//     不会给 GC 添任何活。
//   - 如果真写单调栈版，`stack = append(stack, i)` / `stack = stack[:len(stack)-1]`
//     是 Go 里最惯用的栈写法（用切片当栈），比引 container/list 轻得多。
//   - 累加用 int 就够：本题数据规模下不会溢出 int64 的范围。
func Trap(height []int) int {
	res := 0
	l, r := 0, len(height)-1
	leftMax, rightMax := 0, 0

	for l < r {
		if height[l] < height[r] {
			// 左侧较短：位置 l 的水位由 leftMax 唯一决定。
			if height[l] > leftMax {
				leftMax = height[l]
			}
			res += leftMax - height[l]
			l++
		} else {
			// 右侧较短（含相等）：位置 r 的水位由 rightMax 唯一决定。
			if height[r] > rightMax {
				rightMax = height[r]
			}
			res += rightMax - height[r]
			r--
		}
	}

	return res
}
