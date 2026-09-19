package problems

// LongestPalindrome 求最长回文子串（LeetCode 5. 最长回文子串）。
//
// 题目：给定字符串 s，找到 s 中最长的回文子串（回文即正着读和反着读一样）。
//
// 思路：中心扩展法。回文串的中心只有两种：单字符中心（奇数长度，如 "aba"）
// 和双字符之间的空隙（偶数长度，如 "abba"）。一共 2n-1 个中心，
// 枚举每个中心向两侧扩展，直到两侧字符不同或越界；记录最长的那一段。
// 偶数中心用 (i, i+1) 表示，奇数中心用 (i, i) 表示。
//
// 复杂度：时间 O(n²)（n 个中心，每个中心最多扩 n/2 步，n 为字符数）；
// 空间 O(n)（[]rune(s) 拷贝了一份码点）。
//
// 卡点：
//  1. 中心有 2n-1 个。只枚举「单个字符」中心会漏掉全部偶数长度回文（"bb"、"abba"）。
//  2. 扩展循环退出时两侧已经不相等了，真正的回文区间是 [lo+1, hi-1]。
//     直接用 [lo, hi] 会多带一个不匹配的字符 —— 这是中心扩展法最常写错的一行。
//  3. 相等长度时的取舍要固定（本实现用 > 比较，保留先出现的那个）。
//     否则「多解」用例（"babad" 既可以是 "bab" 也可以是 "aba"）会让测试结果飘。
//  4. 多字节字符必须按 rune 处理。若照搬 s[i] 的字节下标写法，"中文" 的最长回文
//     会退化成一个孤立的 0xE4 字节（非法 UTF-8），因为单个字节永远「等于自己」。
//
// Go 注意点：
//   - 先 []rune(s) 再操作：rune 是 int32 码点，切片下标就是「第几个字符」，多字节安全。
//     代价是一次 O(n) 拷贝，而且不能再拿字节下标去切原串，要用 string(rs[lo:hi+1]) 还原。
//   - 对照：纯 byte 版零拷贝、ASCII 下更快，但遇到 UTF-8 多字节字符会把字符切碎，
//     所以通用实现选 rune 版。两者取舍见文件末尾的 longestPalindromeBrute 注释。
func LongestPalindrome(s string) string {
	rs := []rune(s)
	if len(rs) < 2 {
		return s
	}
	bestLo, bestHi := 0, 0
	for i := range rs {
		if lo, hi := expandRunes(rs, i, i); hi-lo > bestHi-bestLo { // 奇数长度中心
			bestLo, bestHi = lo, hi
		}
		if lo, hi := expandRunes(rs, i, i+1); hi-lo > bestHi-bestLo { // 偶数长度中心
			bestLo, bestHi = lo, hi
		}
	}
	return string(rs[bestLo : bestHi+1])
}

// expandRunes 以 rs[lo]、rs[hi] 为中心向两侧扩展，返回最终回文区间的闭区间下标。
// 退出循环时两侧字符已经不相等（或越界），所以要把多走的那一步退回来。
func expandRunes(rs []rune, lo, hi int) (int, int) {
	for lo >= 0 && hi < len(rs) && rs[lo] == rs[hi] {
		lo--
		hi++
	}
	return lo + 1, hi - 1
}

// longestPalindromeBrute 是暴力对照实现（小写不导出，只在测试里交叉验证）：
// 枚举全部 O(n²) 个子串、每个子串用 O(n) 判断回文，总时间 O(n³)，空间 O(n)。
// 它比中心扩展法慢一个数量级，但逻辑直白、不容易写错，适合当「标准答案」对照。
//
// 两者差别一句话：暴力法是「枚举区间 + 验证回文」，中心扩展法是「枚举中心 + 生长」。
// 暴力法把 O(n) 的回文验证重复做了 O(n²) 次，中心扩展法把验证融进了生长过程里，
// 于是省掉一个 n。两者最终结果的长度一定相同（相等长度时都保留先出现的那个）。
func longestPalindromeBrute(s string) string {
	rs := []rune(s)
	bestLo, bestHi := 0, -1
	for i := 0; i < len(rs); i++ {
		for j := i; j < len(rs); j++ {
			if j-i <= bestHi-bestLo {
				continue // 不可能更长，省掉一次回文检查
			}
			if isPalindromeRunes(rs, i, j) {
				bestLo, bestHi = i, j
			}
		}
	}
	if bestHi < bestLo {
		return ""
	}
	return string(rs[bestLo : bestHi+1])
}

// isPalindromeRunes 判断 rs[lo..hi]（闭区间）是否为回文。
func isPalindromeRunes(rs []rune, lo, hi int) bool {
	for lo < hi {
		if rs[lo] != rs[hi] {
			return false
		}
		lo++
		hi--
	}
	return true
}
