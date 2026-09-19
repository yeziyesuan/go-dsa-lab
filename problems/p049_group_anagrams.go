package problems

import "sort"

// GroupAnagrams 把互为字母异位词的字符串分到同一组（LeetCode 49. 字母异位词分组）。
//
// 题目：给定字符串数组 strs，把互为字母异位词的字符串放进同一个子切片后返回。
//
// 思路：异位词的「计数指纹」相同，用指纹当 map 的 key。
//   - 只含 a-z 时，key = 26 个字节的计数数组（第 i 个字节是字母 'a'+i 的出现次数），
//     一次遍历 O(n) 就能算出来，不需要排序（排序是 O(n log n)）。
//   - 含其它字符（大写、数字、中文……）时 26 个格子不够，退回「rune 排序后的字符串」当 key。
//
// 再用 map[string]int 记录 key → 结果切片下标，遇到新 key 就 append 一个新分组。
//
// 复杂度：时间 O(N·L)（N 个串、平均长 L；走 rune 排序兜底分支时是 O(N·L log L)）；
// 空间 O(N·L)，主要花在 map 的 key 和返回的分组上。
//
// 卡点：
//  1. 返回值顺序不定。分组先后取决于输入顺序，Go 的 map 遍历顺序又是随机的，
//     所以测试不能直接比较 [][]string，必须「组内排序 + 整体按首元素排序」后再比
//     （见测试文件里的 normalizeGroups）。
//  2. 计数数组直接 string(cnt[:]) 当 key 时，每个计数只占 1 个字节，最多表示 255。
//     LeetCode 约束单个串长 <= 100，安全；若串长可能超过 255，同字母计数会回绕串味，
//     必须换成 strconv.Itoa 之类的变长编码。
//  3. 空串 "" 的 key 是 26 个 0，天然和别的空串一组；"" 与 "a" 的 key 不同，不会混组。
//  4. 不能用「字符求和」「字符乘积」当 key：ad/bc 和相同、ab/cc 积相同，都会把不同组混在一起。
//
// Go 注意点：
//   - map 的 key 必须可比较；string 可以，[]byte / []rune 不行 —— 这就是要把计数数组
//     「转换」成 string 的原因。Go 允许 []byte → string 的转换，转换会拷贝一份内容，
//     得到的 string 才是合法且不可变的 map key。
//   - [26]byte 是数组（值类型、可比较，本身也能当 key），但用 string 更省心、更通用。
//   - res[i] = append(res[i], s) 必须把结果写回 res[i]：append 可能换底层数组，
//     只写 append(res[i], s) 不接收返回值，新元素会丢。
func GroupAnagrams(strs []string) [][]string {
	idx := make(map[string]int, len(strs)) // key → res 中的分组下标
	res := make([][]string, 0, len(strs))
	for _, s := range strs {
		k := anagramKey(s)
		i, ok := idx[k]
		if !ok {
			i = len(res)
			idx[k] = i
			res = append(res, []string{s})
			continue
		}
		res[i] = append(res[i], s)
	}
	return res
}

// anagramKey 生成异位词分组的 key：只含 a-z 时用 26 个字节的计数数组，
// 否则退回「rune 排序后的字符串」。
func anagramKey(s string) string {
	if isLowerAlphaASCII(s) {
		var cnt [26]byte
		for i := 0; i < len(s); i++ {
			cnt[s[i]-'a']++
		}
		return string(cnt[:])
	}
	rs := []rune(s)
	sort.Slice(rs, func(i, j int) bool { return rs[i] < rs[j] })
	return string(rs)
}
