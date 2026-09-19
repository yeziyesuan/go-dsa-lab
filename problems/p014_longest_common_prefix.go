package problems

// LongestCommonPrefix 求字符串切片的最长公共前缀（LeetCode 14. 最长公共前缀）。
//
// 题目：给定字符串数组 strs，返回所有字符串的最长公共前缀；没有公共前缀返回 ""。
//
// 思路：纵向扫描。以 strs[0] 为基准，逐列（第 i 个字节）比较所有字符串的第 i 个字节，
// 一旦某个字符串长度不足 i、或第 i 个字节不同，答案就是 strs[0][:i]；
// 所有列都能比完，说明 strs[0] 本身就是公共前缀。
//
// 复杂度：时间 O(S)，S 为所有字符串长度之和（最坏情况下每个字符串的每一列都比较一次）；
// 空间 O(1)（Go 的切片是视图，返回 first[:i] 不拷贝底层数组）。
//
// 卡点：
//  1. 必须先判空切片，否则 strs[0] 会 panic（index out of range）。nil 切片同样要挡住。
//  2. 内层判断要先看 i >= len(s) 再看 s[i]；顺序写反会越界 panic。
//  3. 返回 first[:i] 而不是 first[:i+1]：第 i 列已经确定不同了，不能算进前缀。
//  4. 按字节扫描对多字节字符有隐患：["中", "一"] 在第 2 个字节就分叉，
//     会返回 0xE4 0xB8 这样的「半个汉字」（非法 UTF-8）。要保证返回合法字符串，
//     得先 []rune 再逐字符比较（代价是多一次 O(n) 拷贝）。LeetCode 本题只要求 ASCII，
//     所以字节扫描够用，但这个坑要知道。
//
// Go 注意点：字符串切片 s[i:j] 是共享底层数组的视图，不产生拷贝；
// 比较用 s[i] 拿到的 byte 与字节字面量比较即可，不需要 string(s[i])（那会额外分配）。
func LongestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	first := strs[0]
	for i := 0; i < len(first); i++ {
		c := first[i]
		for _, s := range strs[1:] {
			if i >= len(s) || s[i] != c {
				return first[:i]
			}
		}
	}
	return first
}
