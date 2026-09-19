package problems

import (
	"reflect"
	"sort"
	"testing"
)

func TestGroupAnagrams(t *testing.T) {
	tests := []struct {
		name string
		in   []string
		want [][]string
	}{
		{
			"题目示例",
			[]string{"eat", "tea", "tan", "ate", "nat", "bat"},
			[][]string{{"bat"}, {"nat", "tan"}, {"ate", "eat", "tea"}},
		},
		{"空输入", []string{}, [][]string{}},
		{"nil输入", nil, [][]string{}},
		{"只有一个空串", []string{""}, [][]string{{""}}},
		{"两个空串同组", []string{"", ""}, [][]string{{"", ""}}},
		{"空串与非空串不同组", []string{"", "a"}, [][]string{{""}, {"a"}}},
		{"单元素", []string{"a"}, [][]string{{"a"}}},
		{"全相同字符串", []string{"aaa", "aaa", "aaa"}, [][]string{{"aaa", "aaa", "aaa"}}},
		{"多组混合", []string{"abc", "bca", "cab", "xyz"}, [][]string{{"abc", "bca", "cab"}, {"xyz"}}},
		{"长度不同就不同组", []string{"ab", "ba", "abc", "cba"}, [][]string{{"ab", "ba"}, {"abc", "cba"}}},
		{"和相同但不是异位词", []string{"ad", "bc", "da"}, [][]string{{"ad", "da"}, {"bc"}}},
		{"中文异位词", []string{"中文", "文中", "abc"}, [][]string{{"中文", "文中"}, {"abc"}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GroupAnagrams(tt.in)
			if !reflect.DeepEqual(normalizeGroups(got), normalizeGroups(tt.want)) {
				t.Fatalf("GroupAnagrams(%q) = %v, want %v（分组顺序无关）", tt.in, got, tt.want)
			}
		})
	}
}

// normalizeGroups 消除返回值的顺序不确定性：先给每个分组内部排序，
// 再把所有分组按首元素排序。题目允许多种输出顺序，只有「分组内容」才是答案。
// 前提：分组不会为空（GroupAnagrams 只在放入元素时才新建分组）。
func normalizeGroups(groups [][]string) [][]string {
	out := make([][]string, 0, len(groups))
	for _, g := range groups {
		c := make([]string, len(g))
		copy(c, g)
		sort.Strings(c)
		out = append(out, c)
	}
	sort.Slice(out, func(i, j int) bool { return out[i][0] < out[j][0] })
	return out
}
