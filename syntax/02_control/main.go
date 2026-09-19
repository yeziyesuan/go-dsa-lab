// 语法点：if / for / switch——Go 只有 if、for、switch 三种控制结构，没有 while。
//
// 运行：go run ./syntax/02_control
package main

import "fmt"

func main() {
	fmt.Println("=== if：带初始化语句 ===")
	m := map[string]int{"apple": 3, "banana": 5}

	// 初始化语句里声明的 v、ok 只在 if/else 块内可见，出了大括号就没了。
	// 这是 Go 里最常用的「查 map 顺手判存在」写法。
	if v, ok := m["apple"]; ok {
		fmt.Println("apple 存在，值 =", v)
	} else {
		fmt.Println("apple 不存在")
	}
	if v, ok := m["cherry"]; ok {
		fmt.Println("cherry 存在，值 =", v)
	} else {
		fmt.Println("cherry 不存在，v 的零值 =", v, "（ok 为 false 时 v 是零值）")
	}
	// 注意：这里的 ok 已经不在作用域里了，下面这样写会编译失败。
	//	fmt.Println(ok) // 编译失败：undefined: ok

	// if 的条件必须是 bool，不能像 C 那样写 if n。
	n := 7
	if n%2 == 0 {
		fmt.Println(n, "是偶数")
	} else if n%3 == 0 {
		fmt.Println(n, "能被 3 整除")
	} else {
		fmt.Println(n, "既不是偶数也不能被 3 整除")
	}

	fmt.Println()
	fmt.Println("=== for 的三种写法 ===")

	// 1) 三段式：初始化; 条件; 后置语句（三者都可以省略）
	sum := 0
	for i := 1; i <= 5; i++ {
		sum += i
	}
	fmt.Println("三段式 1+2+3+4+5 =", sum)

	// 2) 条件式：只有条件，相当于别的语言的 while
	// Go 没有 while 关键字，需要 while 就用这种写法。
	k := 1
	for k < 100 {
		k *= 3
	}
	fmt.Println("条件式（while 效果）k 增长到", k)

	// 3) 无限循环 + break / continue
	count := 0
	for {
		count++
		if count%2 == 0 {
			continue // 跳过本轮剩下的语句
		}
		if count >= 7 {
			break // 跳出整个循环
		}
	}
	fmt.Println("无限循环里 count 停在", count)

	fmt.Println()
	fmt.Println("=== for range ===")
	nums := []int{10, 20, 30}
	for i, val := range nums {
		fmt.Printf("索引 %d -> 值 %d\n", i, val)
	}
	// 只要值不要索引，用 _ 占位；Go 不允许声明了不用的变量。
	for _, val := range nums {
		fmt.Print(val, " ")
	}
	fmt.Println()

	// range 字符串时，索引是字节下标，值是该位置的 rune。
	for i, r := range "Go语言" {
		fmt.Printf("字节下标 %d -> 字符 %c\n", i, r)
	}
	// range map：顺序是随机的，见 04_map。
	// range 整数（Go 1.22 起）：直接迭代 0..n-1
	for i := range 3 {
		fmt.Print(i, " ")
	}
	fmt.Println("<- range 整数（Go 1.22 新特性）")

	fmt.Println()
	fmt.Println("=== switch：默认不穿透 ===")
	grade := 'B'
	switch grade {
	case 'A':
		fmt.Println("优秀")
	case 'B':
		fmt.Println("良好")
		// Go 每个 case 结束自动 break，不需要写 break，
		// 也不会像 C 那样自动往下穿透；想穿透必须显式写 fallthrough。
	case 'C':
		fmt.Println("及格")
	default:
		fmt.Println("未知等级")
	}

	// 显式 fallthrough：无条件执行下一个 case 的语句体（很少用，知道即可）。
	switch 1 {
	case 1:
		fmt.Println("case 1 命中，fallthrough 到下一个 case")
		fallthrough
	case 2:
		fmt.Println("这行是被 fallthrough 带进来的")
	case 3:
		fmt.Println("这行不会执行")
	}

	// 一个 case 可以列多个值。
	switch day := 6; day {
	case 6, 7:
		fmt.Println("day =", day, "，周末")
	default:
		fmt.Println("day =", day, "，工作日")
	}

	fmt.Println()
	fmt.Println("=== switch { case 条件: }：无表达式的 switch ===")
	// switch 后面不写表达式时，等价于 switch true，每个 case 就是一个布尔条件，
	// 比一串 if / else if 更整齐。
	x := -3
	switch {
	case x > 0:
		fmt.Println(x, "是正数")
	case x < 0:
		fmt.Println(x, "是负数")
	default:
		fmt.Println(x, "是零")
	}
}
