package problems

import "testing"

func TestLongestCommonPrefix(t *testing.T) {
	tests := []struct {
		name string
		in   []string
		want string
	}{
		{"题目示例_有公共前缀", []string{"flower", "flow", "flight"}, "fl"},
		{"题目示例_无公共前缀", []string{"dog", "racecar", "car"}, ""},
		{"空切片", []string{}, ""},
		{"nil切片", nil, ""},
		{"单个空串", []string{""}, ""},
		{"只有一个元素", []string{"abc"}, "abc"},
		{"全相同字符串", []string{"abc", "abc", "abc"}, "abc"},
		{"某个串为空", []string{"abc", ""}, ""},
		{"最短串就是答案", []string{"ab", "abc", "abcd"}, "ab"},
		{"大小写敏感", []string{"Abc", "abc"}, ""},
		{"第一列就不同", []string{"a", "b"}, ""},
		{"中文公共前缀_一个字", []string{"中文学校", "中文", "中华"}, "中"},
		{"中文完全相同", []string{"你好", "你好"}, "你好"},
		{"中文与英文混合", []string{"go语言", "go入门"}, "go"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LongestCommonPrefix(tt.in); got != tt.want {
				t.Fatalf("LongestCommonPrefix(%q) = %q, want %q", tt.in, got, tt.want)
			}
		})
	}
}
