package problems

import (
	"strings"
	"testing"
)

func TestAddStrings(t *testing.T) {
	tests := []struct {
		name string
		a    string
		b    string
		want string
	}{
		{"题目示例_1", "11", "123", "134"},
		{"题目示例_2", "456", "77", "533"},
		{"题目示例_3", "0", "0", "0"},
		{"单字符进位", "1", "9", "10"},
		{"连续进位", "99", "1", "100"},
		{"全9相加", "999", "999", "1998"},
		{"全相同字符", "111", "111", "222"},
		{"一边为空", "123", "", "123"},
		{"另一边为空", "", "8", "8"},
		{"两边都为空", "", "", ""},
		{"加零不变", "1234", "0", "1234"},
		{"位数差很多", "1", "1000000000", "1000000001"},
		// 下面这些远超 int64 上限（约 9.2e18，19 位），Atoi/ParseInt 会 ErrRange，
		// 竖式加法必须照样算对。两个 30 位串相加时每列都是 9+1+进位=11，
		// 所以结果是 30 个 1 后面再跟 1 个 0。
		{"超长_20位9加1", strings.Repeat("9", 20), "1", "1" + strings.Repeat("0", 20)},
		{"超长_100位9加1", strings.Repeat("9", 100), "1", "1" + strings.Repeat("0", 100)},
		{"超长_30位9加30位1", strings.Repeat("9", 30), strings.Repeat("1", 30), strings.Repeat("1", 30) + "0"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AddStrings(tt.a, tt.b); got != tt.want {
				t.Fatalf("AddStrings(%q, %q) = %q, want %q", tt.a, tt.b, got, tt.want)
			}
			// 加法满足交换律，顺便验证两个参数位置都走通（补 0 分支要对称）。
			if got := AddStrings(tt.b, tt.a); got != tt.want {
				t.Fatalf("AddStrings(%q, %q) = %q, want %q（交换参数后）", tt.b, tt.a, got, tt.want)
			}
		})
	}
}
