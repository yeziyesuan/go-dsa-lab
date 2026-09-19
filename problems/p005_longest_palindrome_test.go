package problems

import "testing"

func TestLongestPalindrome(t *testing.T) {
	// want 放「所有可接受的答案」：相同长度时题目允许多解
	// （"babad" 返回 "bab" 或 "aba" 都对），所以用集合比较而不是字符串比较。
	tests := []struct {
		name string
		in   string
		want []string
	}{
		{"题目示例_多解", "babad", []string{"bab", "aba"}},
		{"题目示例_偶数回文", "cbbd", []string{"bb"}},
		{"空串", "", []string{""}},
		{"单字符", "a", []string{"a"}},
		{"两个不同字符_多解", "ac", []string{"a", "c"}},
		{"全相同字符", "aaaa", []string{"aaaa"}},
		{"整个串是回文_奇数", "racecar", []string{"racecar"}},
		{"整个串是回文_偶数", "abba", []string{"abba"}},
		{"前缀是最长回文", "abac", []string{"aba"}},
		{"后缀是最长回文", "caba", []string{"aba"}},
		{"中文偶数回文", "中中", []string{"中中"}},
		{"中文整串回文", "上海自来水来自海上", []string{"上海自来水来自海上"}},
		{"中文非回文_多解", "中文", []string{"中", "文"}},
		{"中英混合", "a中a", []string{"a中a"}},
	}

	// 用闭包而不是包级函数，避免和同 package 其它题解文件重名。
	acceptable := func(want []string, got string) bool {
		for _, w := range want {
			if w == got {
				return true
			}
		}
		return false
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := LongestPalindrome(tt.in)
			if !acceptable(tt.want, got) {
				t.Fatalf("LongestPalindrome(%q) = %q, 不在可接受集合 %v 中", tt.in, got, tt.want)
			}

			// 交叉验证暴力解：长度必须与中心扩展法一致，且同样落在可接受集合里。
			brute := longestPalindromeBrute(tt.in)
			if len([]rune(brute)) != len([]rune(got)) {
				t.Fatalf("暴力解与中心扩展法长度不一致: brute=%q(%d 字符) center=%q(%d 字符)",
					brute, len([]rune(brute)), got, len([]rune(got)))
			}
			if !acceptable(tt.want, brute) {
				t.Fatalf("longestPalindromeBrute(%q) = %q, 不在可接受集合 %v 中", tt.in, brute, tt.want)
			}
		})
	}
}
