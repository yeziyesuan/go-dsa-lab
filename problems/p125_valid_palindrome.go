package problems

// IsPalindrome 判断字符串是否为回文串（LeetCode 125. 验证回文串）。
//
// 题目：给定字符串 s，只考虑其中的字母和数字字符、忽略字母大小写，
// 判断它是否为回文串。空串、以及去掉非字母数字后为空的串都算回文。
//
// 思路：左右双指针相向而行。左指针先跳过所有非字母数字字符，右指针也跳过，
// 然后比较两者的小写形式；不同直接 false，相同则两指针一起向中间收缩。
//
// 复杂度：时间 O(n)，空间 O(1)（不构造过滤后的新字符串，只扫原串一遍）。
//
// 卡点：
//  1. 跳过无效字符必须「跳到有效字符为止」，用 continue 回到循环重新判断；
//     只 if 一次就往下走，遇到连续标点（如 "a,,,a"）会拿标点去比较，直接判错。
//  2. 忽略大小写要显式限定在字母范围内（'A'-'Z' 才转小写）。
//     用「两字符相差 32」这类技巧忽略大小写会把数字误判成字母：
//     '0'(48) 与 'P'(80) 恰好差 32，于是 "0P" 被错判成回文 —— 这是本题经典陷阱。
//  3. 循环条件是 i < j。多字节字符整体会被跳过，所以 "中文" 这种全非字母数字的串
//     按题目定义返回 true（过滤后是空串，空串是回文）。
//
// Go 注意点（byte 与 rune 的区别，本题重点）：
//   - byte 是 uint8，表示「一个字节」；rune 是 int32，表示「一个 Unicode 码点」。
//     Go 的 string 底层是只读的字节序列，s[i] 取到的是第 i 个字节，不是第 i 个字符；
//     一个汉字在 UTF-8 里占 3 个字节，"中"[0] 是 0xE4，只是编码的头一个字节。
//   - 本题用 byte 遍历是安全的：题目只要求比较 ASCII 字母和数字，而 ASCII 字符在
//     UTF-8 里恒为单字节（0x00-0x7F），永远不可能是多字节字符的续字节
//     （续字节一律 >= 0x80）。续字节会被 isAlnumASCII 判为「非字母数字」直接跳过，
//     所以既不会漏掉有效字符，也不会把半个汉字当成字符去比较。
//   - 什么时候不安全：一旦需要「按字符」处理非 ASCII —— 统计字符个数、按字符反转、
//     判断中文回文、按字符下标切片 —— byte 遍历就会把多字节字符拆开。
//     此时要么 for range s（rune 迭代，下标是字节偏移），要么先 []rune(s)。
func IsPalindrome(s string) bool {
	i, j := 0, len(s)-1
	for i < j {
		if !isAlnumASCII(s[i]) {
			i++
			continue
		}
		if !isAlnumASCII(s[j]) {
			j--
			continue
		}
		if lowerASCII(s[i]) != lowerASCII(s[j]) {
			return false
		}
		i++
		j--
	}
	return true
}

// isAlnumASCII 判断字节 b 是否为 ASCII 字母或数字（0-9 / a-z / A-Z）。
// 多字节 UTF-8 的每个字节都 >= 0x80，因此一定返回 false，会被调用方跳过。
func isAlnumASCII(b byte) bool {
	switch {
	case b >= '0' && b <= '9':
		return true
	case b >= 'a' && b <= 'z':
		return true
	case b >= 'A' && b <= 'Z':
		return true
	default:
		return false
	}
}

// lowerASCII 只把 ASCII 大写字母转成小写，其它字节原样返回，
// 这样数字和标点不会被误伤（对比无脑 b|0x20 的写法）。
func lowerASCII(b byte) byte {
	if b >= 'A' && b <= 'Z' {
		return b + ('a' - 'A')
	}
	return b
}
