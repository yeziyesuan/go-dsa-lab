package problems

import (
	"testing"
	"unicode/utf8"
)

func TestReverseString(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want string
	}{
		{"题目示例_hello", "hello", "olleh"},
		{"题目示例_Hannah", "Hannah", "hannaH"},
		{"空串", "", ""},
		{"单字符", "a", "a"},
		{"全相同字符", "aaaa", "aaaa"},
		{"偶数长度", "ab", "ba"},
		{"回文串_反转后不变", "aba", "aba"},
		{"含空格标点", "a b!", "!b a"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := []byte(tt.in)
			ReverseString(got)
			if string(got) != tt.want {
				t.Fatalf("ReverseString(%q) = %q, want %q", tt.in, string(got), tt.want)
			}
		})
	}

	// 多字节字符：ReverseString 按 byte 反转，会把 UTF-8 编码打乱（结果不再是合法 UTF-8），
	// 但「反转两次等于原样」这条性质依然成立。这里把这两点都固定成测试。
	t.Run("中文多字节_按字节反转可逆但会破坏UTF8", func(t *testing.T) {
		src := []byte("中文ab")
		orig := string(src)

		ReverseString(src)
		if string(src) == orig {
			t.Fatalf("按字节反转后不应与原串相同, got %q", string(src))
		}
		if utf8.Valid(src) {
			t.Fatalf("按字节反转多字节字符串后不应仍是合法 UTF-8, got %q", string(src))
		}
		t.Logf("按字节反转后的原始字节: %q（非法 UTF-8，属于预期行为）", string(src))

		ReverseString(src)
		if string(src) != orig {
			t.Fatalf("反转两次应还原, got %q, want %q", string(src), orig)
		}
	})
}
