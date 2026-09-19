package problems

// AddStrings 用字符串模拟十进制大数相加（LeetCode 415. 字符串相加）。
//
// 题目：给定两个非负整数的字符串形式 num1、num2，返回它们的和（同样用字符串表示）。
// 不能使用任何大数库，也不能把输入直接转成内置整数类型。
//
// 思路：模拟竖式加法。i、j 分别从两个串的末尾往前走，carry 记录进位；
// 每轮 sum = carry + (i 位数字) + (j 位数字)，结果位是 sum%10，新进位是 sum/10。
// 因为是从低位往高位算，先追加进缓冲区的是低位，最后整体反转才是正确顺序。
//
// 复杂度：时间 O(max(m, n))，空间 O(max(m, n))（结果缓冲区）。
//
// 卡点：
//  1. 循环条件是 i >= 0 || j >= 0 || carry > 0，三个条件缺一不可：
//     少了 carry 会丢掉最高位的进位（"99" + "1" 会算成 "00"，正确是 "100"）。
//  2. 两个串长度不同：短的一侧在 i/j < 0 时要补 0，不能越界取值。
//  3. 结果是从低位往高位生成的，最后必须反转；忘了反转会得到倒过来的答案。
//  4. 结果字符要写 byte('0'+sum%10)，直接写 sum%10 会得到不可见的控制字符。
//  5. "0" + "0" 要返回 "0" 而不是 ""：循环至少要执行一次（靠 i >= 0 条件保证）。
//
// Go 注意点（为什么不用 strconv.Atoi / ParseInt）：
//   - Atoi/ParseInt 的结果受 int64 限制（Atoi 得到 int，64 位平台也是 64 位，
//     上限约 9.2e18）。题目给的是任意长度数字串，几百位很常见，ParseInt 会返回
//     ErrRange，改用 ParseFloat 又会丢精度 —— 从根上就不该转数字。
//   - 竖式加法逐位处理，耗时只和位数有关，天然支持任意长度，这也是「大数运算」的基本功。
//   - 字符转数字用 c - '0'；byte 相减仍是 byte，参与算术前显式转 int，
//     否则 byte + byte 会按 uint8 回绕（'9' + '9' = 0x12 而不是 18）。
//   - 缓冲区用 make([]byte, 0, n) 预留容量，避免 append 过程中反复扩容。
func AddStrings(num1, num2 string) string {
	i, j := len(num1)-1, len(num2)-1
	carry := 0
	n := len(num1)
	if len(num2) > n {
		n = len(num2)
	}
	buf := make([]byte, 0, n+1)
	for i >= 0 || j >= 0 || carry > 0 {
		sum := carry
		if i >= 0 {
			sum += int(num1[i] - '0')
			i--
		}
		if j >= 0 {
			sum += int(num2[j] - '0')
			j--
		}
		buf = append(buf, byte('0'+sum%10))
		carry = sum / 10
	}
	// 目前 buf 是低位在前，反转成高位在前才是答案。
	for l, r := 0, len(buf)-1; l < r; l, r = l+1, r-1 {
		buf[l], buf[r] = buf[r], buf[l]
	}
	return string(buf)
}
