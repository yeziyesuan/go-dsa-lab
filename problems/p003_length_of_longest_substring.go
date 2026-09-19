package problems

// LengthOfLongestSubstring 求不含重复字符的最长子串长度（LeetCode 3）。
//
// 题目：给定字符串 s，找出其中不含重复字符的最长子串，返回该子串的长度。
//
// 思路：滑动窗口 [l, r]。r 从左往右扫，用 map[byte]int 记住每个字符「最近一次出现的下标」。
// 若 s[r] 之前在窗口内出现过（下标 >= l），就把左边界挪到「上次出现位置 + 1」；
// 每轮用当前窗口长度 r-l+1 更新答案。l 和 r 都只向右走，所以整体是线性的。
//
// 复杂度：时间 O(n)（每个字符最多被 r 扫到一次、被 l 跨过一次）；
// 空间 O(min(n, 字符集大小))，即 map[byte]int 的大小。
//
// 卡点（本题最典型的坑）：
//  1. 左边界必须取 max：l = max(l, i+1)，不能直接写 l = i+1。
//     因为 map 里存的 i 可能落在窗口左边之外（是「很久以前」出现过的字符），
//     直接赋值会让 l 向左回退，窗口里混进重复字符，答案偏大。
//     复现用例 "abba"：扫到最后一个 'a' 时 last['a'] = 0，而此时 l 已经是 2，
//     写成 l = 0+1 = 1 就回退了一名，窗口变成 "bba"，算出 3（正确答案是 2）。
//  2. map 必须存「下标」而不是「是否出现过」，否则无法知道左边界该跳到哪。
//  3. 更新答案要放在移动左边界之后，用当前窗口长度 r-l+1（不是全局长度）。
//
// Go 注意点：
//   - 用 map[byte]int 而不是 map[rune]int：题目字符集是 ASCII（字母、数字、符号、空格），
//     一字符即一字节，byte 当 key 更省内存也更快；用 rune 会多一次 UTF-8 解码。
//   - 但按 byte 滑窗对多字节字符是「按字节去重」，长度也是字节数而不是字符数
//     （见测试里的中文用例："中中" 返回 3）。要按「字符」处理必须 range s 或 []rune(s)。
//   - Go 1.21 起内置了 max/min，可以直接写 max(l, i+1)，不用自己定义辅助函数；
//     等价的显式写法是 if i+1 > l { l = i + 1 }。
func LengthOfLongestSubstring(s string) int {
	last := make(map[byte]int, len(s)) // 字符 → 最近一次出现的下标
	best, l := 0, 0
	for r := 0; r < len(s); r++ {
		c := s[r]
		if i, ok := last[c]; ok {
			l = max(l, i+1) // 关键：取 max，防止左边界回退（坑："abba"）
		}
		last[c] = r
		if r-l+1 > best {
			best = r - l + 1
		}
	}
	return best
}
