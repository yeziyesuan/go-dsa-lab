package problems

import "testing"

func TestLengthOfLongestSubstring(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want int
	}{
		{"题目示例_abcabcbb", "abcabcbb", 3},
		{"题目示例_bbbbb", "bbbbb", 1},
		{"题目示例_pwwkew", "pwwkew", 3},
		{"空串", "", 0},
		{"单字符", "a", 1},
		{"全相同字符", "aaaaaa", 1},
		{"卡点用例_abba", "abba", 2},
		{"卡点用例_abba带尾巴", "abbac", 3},
		{"卡点用例_tmmzuxt", "tmmzuxt", 5},
		{"卡点用例_dvdf", "dvdf", 3},
		{"无重复字符", "abcdef", 6},
		{"含空格与符号", "ab c!d", 6},
		{"中文_按字节去重", "中中", 3},
		{"中文单字_按字节计数", "中", 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LengthOfLongestSubstring(tt.in); got != tt.want {
				t.Fatalf("LengthOfLongestSubstring(%q) = %d, want %d", tt.in, got, tt.want)
			}
		})
	}
}
