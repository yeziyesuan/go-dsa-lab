package problems

import "testing"

func TestIsAnagram(t *testing.T) {
	tests := []struct {
		name string
		s    string
		t    string
		want bool
	}{
		{"题目示例_是异位词", "anagram", "nagaram", true},
		{"题目示例_不是异位词", "rat", "car", false},
		{"空串与空串", "", "", true},
		{"空串与非空串", "", "a", false},
		{"单字符相同", "a", "a", true},
		{"单字符不同", "a", "b", false},
		{"长度不等", "ab", "a", false},
		{"全相同字符_相等", "aaaa", "aaaa", true},
		{"全相同字符_长度不等", "aaaa", "aaa", false},
		{"同字母不同个数", "aab", "abb", false},
		{"顺序打乱_是异位词", "listen", "silent", true},
		{"和相同但不是异位词", "ad", "bc", false},
		{"含数字", "a1b", "b1a", true},
		{"含数字_不同", "a1b", "b2a", false},
		{"大小写敏感", "Abc", "abc", false},
		{"大小写混合_不能走快路径", "abc", "abC", false},
		{"中文异位词", "中文", "文中", true},
		{"中文非异位词", "中文", "中a", false},
		{"中文异位词_你好", "你好", "好你", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsAnagram(tt.s, tt.t); got != tt.want {
				t.Fatalf("IsAnagram(%q, %q) = %v, want %v", tt.s, tt.t, got, tt.want)
			}
			// 异位词关系是对称的，顺便验证 s/t 互换后结论一致。
			if got := IsAnagram(tt.t, tt.s); got != tt.want {
				t.Fatalf("IsAnagram(%q, %q) = %v, want %v（交换参数后）", tt.t, tt.s, got, tt.want)
			}
		})
	}
}
