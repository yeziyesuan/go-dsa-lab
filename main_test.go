package main

import "testing"

// W1 的第一个测试：先把 go test 这条链路跑通，再谈别的。
func TestHello(t *testing.T) {
	cases := []struct {
		name string
		in   string
		want string
	}{
		{"普通名字", "ziye", "hello, ziye"},
		{"空名字用默认值", "", "hello, Gopher"},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if got := Hello(c.in); got != c.want {
				t.Errorf("Hello(%q) = %q, want %q", c.in, got, c.want)
			}
		})
	}
}

func TestWeek1Goal(t *testing.T) {
	if Week1Goal() == "" {
		t.Fatal("Week1Goal() 不应为空")
	}
}
