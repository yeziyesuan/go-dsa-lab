# 42. 接雨水

- 难度 / 类型：困难 · 双指针 / 单调栈
- 思路（一句话）：位置 `i` 的水位 = `min(左侧最高, 右侧最高) - height[i]`；用双指针 + `leftMax`/`rightMax` 把预处理数组的 O(n) 空间压成 O(1)。
- 复杂度：时间 O(n)，空间 O(1)
- 卡点：核心是**先判断、后结算，且结算谁就动谁**。当 `height[l] < height[r]` 时右侧一定存在比 `height[l]` 更高的柱子，所以 `min(leftMax, 真实右侧最高) == leftMax`，位置 `l` 的水位已经确定，可以放心 `res += leftMax - height[l]`。顺序写反（无条件先累加）会在水位未定时多加，甚至加出负数。第二个坑是忘了水位差可能为 0：先 `leftMax = max(leftMax, height[l])` 再减，差值天然 `>= 0`，不用额外判断。第三个坑是空数组/单元素，循环不进入、返回 0，别在外面先访问 `height[0]`。
- 单调栈解法（文字版）：从左到右扫，维护**单调递减**的下标栈。遇到 `height[i]` 大于栈顶高度时，栈顶就是槽底 `mid`，pop 之后的新栈顶 `left` 是左墙、`i` 是右墙，这一层水带的水量 = `(min(height[left], height[i]) - height[mid]) * (i - left - 1)`。累加所有 pop 出来的层就是答案。它是「按行/横向」累加，双指针是「按列/竖向」累加，都是 O(n)，但栈版要 O(n) 额外空间，且「pop 后新栈顶是左墙」这个不变量更绕。第一周先把双指针写熟。
- Go 注意点：手写 `if a > b` 而不引 `math` 包。双指针版零分配——没有 `make`、没有 `append`，不给 GC 添活。真要写单调栈版，`stack = append(stack, i)` / `stack = stack[:len(stack)-1]` 就是 Go 里最惯用的栈写法，不必引 `container/list`。
- 关键代码片段：

```go
for l < r {
	if height[l] < height[r] {
		if height[l] > leftMax {
			leftMax = height[l]
		}
		res += leftMax - height[l] // 左侧水位已确定
		l++
	} else {
		if height[r] > rightMax {
			rightMax = height[r]
		}
		res += rightMax - height[r]
		r--
	}
}
```

- 复做日期：9/26 ⬜ / 10/3 ⬜
