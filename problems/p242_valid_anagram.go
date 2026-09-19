package problems

// IsAnagram 判断两个字符串是否互为字母异位词（LeetCode 242. 有效的字母异位词）。
//
// 题目：给定字符串 s 和 t，判断 t 是否为 s 的字母异位词
// （字母种类和每个字母的出现次数完全相同，只是排列顺序不同）。
//
// 思路：长度不同直接 false（异位词长度必然相同，一步挡掉大部分用例）。
// 长度相同就计数：s 的字符 +1、t 的字符 -1，最后计数全为 0 即是异位词。
// 只含 a-z 时用定长数组 [26]int 计数（快路径）；字符集超出 a-z 时退回 map[rune]int。
//
// 复杂度：快路径时间 O(n)、空间 O(1)（[26]int 是定长数组）；
// 慢路径时间 O(n)、空间 O(k)，k 为不同字符的个数。
//
// 卡点：
//  1. 快路径的前置判断必须同时检查 s 和 t。只查 s 时，若 t 里出现 'A' 这类非小写字符，
//     cnt[t[i]-'a'] 会算出负下标 —— 运行期 panic（下标越界）。这是最容易漏的一处。
//  2. 不能用「字符求和相等」判断异位词：ad 与 bc 和相同，但并不是异位词。
//     同理「字符乘积相等」也不行（ab 与 cc 乘积相同）。
//  3. 长度比较直接 len(s) != len(t)（按字节）就够：异位词的字符多重集相同，
//     编码后的字节数必然相同。注意它只是「必要条件」，只用来快速否决。
//  4. 「一边 +1 一边 -1」比「分别统计两个数组再逐个比较」少一次循环，也不容易漏字符。
//
// Go 注意点：
//   - 为什么 [26]int 比 map[byte]int 快：定长数组在栈上分配（是否逃逸留到 W43 之后再看），
//     寻址就是「基址 + 下标」一条指令；map 每次读写都要算哈希、定位桶、比较 key，
//     还可能触发扩容，并且多一层指针间接寻址和内存分配。字符集已知且很小时，
//     数组计数是常数级最优解 —— 这就是「key 空间小用数组，key 空间大才用 map」。
//   - 若字符集超出 a-z（大写字母、数字、中文等多字节字符），26 个格子装不下，
//     应当改用 map[rune]int（本实现自带这个兜底分支）。key 必须用 rune：
//     汉字在 UTF-8 里是多字节，用 byte 计数会把它拆成 3 个「字符」，语义就错了。
//   - range string 拿到的是 (字节下标, rune)；想要「第几个字符」要用 []rune 或自增计数器。
func IsAnagram(s, t string) bool {
	if len(s) != len(t) {
		return false
	}

	// 快路径：两侧都只含 a-z，用定长数组计数（必须同时检查 s 和 t，否则会负下标 panic）。
	if isLowerAlphaASCII(s) && isLowerAlphaASCII(t) {
		var cnt [26]int
		for i := 0; i < len(s); i++ {
			cnt[s[i]-'a']++
			cnt[t[i]-'a']--
		}
		for _, c := range cnt {
			if c != 0 {
				return false
			}
		}
		return true
	}

	// 慢路径：字符集超出 a-z，退回 map[rune]int。
	cnt := make(map[rune]int, len(s))
	for _, r := range s {
		cnt[r]++
	}
	for _, r := range t {
		cnt[r]--
	}
	for _, c := range cnt {
		if c != 0 {
			return false
		}
	}
	return true
}

// isLowerAlphaASCII 判断 s 是否只由小写 ASCII 字母 a-z 组成（空串返回 true）。
// 多字节 UTF-8 的首字节 >= 0x80，会直接返回 false，从而交给 map 兜底分支处理。
func isLowerAlphaASCII(s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] < 'a' || s[i] > 'z' {
			return false
		}
	}
	return true
}
