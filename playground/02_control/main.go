package main

import "fmt"

func main() {
	fmt.Println("===带初始化语句===")

	m := map[string]int{"apple": 3, "banana": 5}
	if v, ok := m["apple"]; ok {
		fmt.Println("apple存在,值=", v)
	} else {
		fmt.Println("apple不存在")
	}
	if v, ok := m["cherry"]; ok {
		fmt.Println("cherry存在,值=", v)
	} else {
		fmt.Println("cherry不存在")
	}
	a := 7
	if a%2 == 0 {
		fmt.Println("是偶数", a)
	} else if a%3 == 0 {
		fmt.Println("能被3整除", a)
	} else {
		fmt.Println("既不是偶数也不能被3整除", a)
	}
	fmt.Println()
	fmt.Println("===for的三种写法===")

	sum := 0
	for i := 0; i <= 5; i++ {
		sum += i
	}
	fmt.Println("三段式 1+2+3+4+5=", sum)

	k := 1
	for k < 100 {
		k *= 3
	}
	fmt.Println("条件式(while 效果)k增长到", k)

	count := 0
	for {
		count++
		if count%2 == 0 {
			continue
		} else if count >= 7 {
			break
		}
	}
	fmt.Println("无限循环里count停在", count)
	fmt.Println()
	fmt.Println("===for range===")
	arr := [3]int{10, 20, 30}
	for e, v := range arr {
		fmt.Printf("索引%d -> 值%d\n", e, v) //fmt.Printf("索引 %d", e, "-> 值 %d\n", v)
	}

	for _, v := range arr {
		fmt.Print(v, "") //fmt.Printf( "",v)
	}
	fmt.Println()

	for x, v := range "Go语言" {
		fmt.Printf("字节下标 %d -> 字符 %c\n", x, v) //fmt.Printf("字节下标%d", x, "->字符%c\n", v)
	}
	for v := range 3 {
		fmt.Print(v)
	}
	fmt.Println("<-整数（Go 1.22新特性）")

	fmt.Println()
	fmt.Println("===switch:默认不穿透===")
	gradn := 'B'
	switch gradn {
	case 'A':
		fmt.Println("优秀")
	case 'B':
		fmt.Println("良好")
	case 'C':
		fmt.Println("及格")
	}

	switch 1 {
	case 1:
		fmt.Println("case 1 命中，fallthrough 到下一个case")
		fallthrough
	case 2:
		fmt.Println("这一行是被fallthrough带进来的")
	case 3:

	}

	switch day := 6; day {
	case 6, 7:
		fmt.Println("day = 6,周末")
	default:
		fmt.Println("上班")
	}
	fmt.Println()
	fmt.Println("===switch{case条件:}: 无表达式的 switch===")
	z := -3
	switch {
	case z > 0:
		fmt.Println(z, "是正数")
	case z < 0:
		fmt.Println(z, "是负数")
	default:
		fmt.Println("为0")
	}

}
