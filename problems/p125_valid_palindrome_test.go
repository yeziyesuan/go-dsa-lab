package problems

import "testing"

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want bool
	}{
		{"题目示例_带标点回文", "A man, a plan, a canal: Panama", true},
		{"题目示例_非回文", "race a car", false},
		{"题目示例_只有空格", " ", true},
		{"空串", "", true},
		{"单字符_字母", "a", true},
		{"单字符_数字", "7", true},
		{"单字符_标点", ",", true},
		{"全相同字符", "aaaa", true},
		{"纯字母回文", "level", true},
		{"短串非回文", "ab", false},
		{"纯数字回文", "12321", true},
		{"数字夹字母", "1a2", false},
		{"大小写混合_是回文", "Aa", true},
		{"经典陷阱_0P", "0P", false},
		{"连续多个标点", "a,,,a", true},
		{"全部是非字母数字", ".,!", true},
		{"中文_多字节整体被跳过", "中文", true},
		{"中文回文_多字节整体被跳过", "上海自来水来自海上", true},
		{"中文夹字母_是回文", "中a中", true},
		{"中文夹字母_大小写也算", "中A中", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsPalindrome(tt.in); got != tt.want {
				t.Fatalf("IsPalindrome(%q) = %v, want %v", tt.in, got, tt.want)
			}
		})
	}
}
