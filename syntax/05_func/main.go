// 语法点：函数——多返回值、命名返回值、可变参数、闭包，以及 defer 的 LIFO 与改返回值。
//
// 运行：go run ./syntax/05_func
package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("=== 多返回值 ===")
	// Go 的惯用法：结果 + error 一起返回，调用方必须显式处理 err。
	q, r, err := divmod(17, 5)
	if err != nil {
		fmt.Println("出错：", err)
	} else {
		fmt.Println("17 / 5 =", q, "余", r)
	}

	_, _, err = divmod(1, 0)
	if err != nil {
		fmt.Println("除数为 0 时返回错误：", err)
	}

	// 常见写法：只用得到值时，用 _ 丢掉不想要的返回值。
	v, _ := twoValues()
	fmt.Println("只要第一个返回值：", v)

	fmt.Println()
	fmt.Println("=== 命名返回值 ===")
	// 命名返回值会在函数开始时被声明并置为零值，可以裸 return。
	fmt.Println("namedSum(3, 4) =", namedSum(3, 4))
	// 多返回值不能直接塞进 Println 当单个参数用（会编译失败：
	// multiple-value namedQuotient(7, 2) in single-value context），
	// 必须先接住再打印。
	// 注意变量名别和前面已经声明的 q、err 撞车，否则 := 会报 no new variables。
	quot, err := namedQuotient(7, 2)
	if err != nil {
		fmt.Println("namedQuotient(7, 2) 出错：", err)
	} else {
		fmt.Println("namedQuotient(7, 2) =", quot, "（商按命名返回值 return）")
	}
	fmt.Println()
	// 命名返回值 + defer 是「延迟修改返回值」的经典用法，见 defer 一节。
	fmt.Println("countUp(5) 返回 =", countUp(5), "（defer 在 return 之后又加了一次）")

	fmt.Println()
	fmt.Println("=== 可变参数 ===")
	fmt.Println("sum() =", sum())
	fmt.Println("sum(1) =", sum(1))
	fmt.Println("sum(1, 2, 3, 4, 5) =", sum(1, 2, 3, 4, 5))

	nums := []int{10, 20, 30}
	// 已经有了切片，用 nums... 展开传入。
	fmt.Println("sum(nums...) =", sum(nums...))
	// 注意：传切片时必须加 ...，直接写 sum(nums) 会编译失败（类型不匹配）。

	fmt.Println()
	fmt.Println("=== 闭包 ===")
	// 闭包 = 函数 + 它捕获的外部变量。next 捕获了 count，count 在多次调用间存活。
	next := counter()
	fmt.Println("第一次调用 next() =", next())
	fmt.Println("第二次调用 next() =", next())
	fmt.Println("第三次调用 next() =", next())

	// 再调用一次 counter() 会得到一份独立的 count，互不影响。
	other := counter()
	fmt.Println("另一个计数器 other() =", other(), "，next() 仍然是", next())

	// 循环变量捕获：Go 1.22 起每轮迭代都是新变量，所以这里会打印 0 1 2。
	var closures []func()
	for i := range 3 {
		closures = append(closures, func() { fmt.Print(i, " ") })
	}
	for _, fn := range closures {
		fn()
	}
	fmt.Println("<- Go 1.22 起循环变量每轮独立；1.21 及以前这里会全部打印 3")

	fmt.Println()
	fmt.Println("=== defer 的执行顺序是 LIFO ===")
	// defer 注册在函数返回时执行，后注册的先执行（栈结构）。
	deferOrder()

	fmt.Println()
	fmt.Println("=== defer 修改命名返回值 ===")
	// 执行顺序：return 把 4 赋给命名返回值 result -> defer 把 result 乘 2 -> 函数真正返回。
	// 所以最终返回 8，而不是 4。
	fmt.Println("addThenDouble(3) =", addThenDouble(3), "（不是 4）")
	fmt.Println("recoverDemo 子函数返回：", recoverDemo())
}

// divmod 返回商、余数和错误：Go 用多返回值代替异常。
func divmod(a, b int) (int, int, error) {
	if b == 0 {
		return 0, 0, errors.New("除数不能为 0")
	}
	return a / b, a % b, nil
}

// twoValues 返回两个 int，用来演示用 _ 丢弃返回值。
func twoValues() (int, int) {
	return 7, 8
}

// namedSum 使用命名返回值，函数体内可以裸 return。
func namedSum(a, b int) (total int) {
	total = a + b
	return // 等价于 return total
}

// namedQuotient 演示命名返回值同样可以显式 return 覆盖。
func namedQuotient(a, b int) (q int, err error) {
	if b == 0 {
		err = errors.New("除数不能为 0")
		return 0, err
	}
	q = a / b
	return q, nil
}

// countUp 演示 defer 可以修改命名返回值：返回前再 +1。
func countUp(n int) (ret int) {
	defer func() {
		ret++ // ret 是命名返回值，这个修改会生效
	}()
	return n
}

// sum 是可变参数函数：nums 在函数内就是一个 []int。
func sum(nums ...int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

// counter 返回一个闭包，每次调用返回递增的计数。
func counter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

// deferOrder 演示 defer 的 LIFO 顺序，以及 defer 参数在注册时求值。
func deferOrder() {
	for i := 1; i <= 3; i++ {
		// 注意：i 在 defer 注册的那一刻就被求值复制了，所以打印的是 1 2 3；
		// 如果这里写 func(){ fmt.Println(i) }()，闭包在函数返回时才读 i。
		defer fmt.Println("defer 注册的 i =", i)
	}
	fmt.Println("deferOrder 函数体结束，下面开始按 LIFO 执行 defer")
}

// addThenDouble 演示 defer 修改命名返回值：返回值会是 4+4=8，而不是 4。
func addThenDouble(n int) (result int) {
	defer func() {
		result *= 2 // 在 return 赋值之后执行，所以能改到最终返回值
	}()
	return n + 1 // 先把 4 赋给 result，再执行 defer
}

// recoverDemo 演示 defer + recover 把 panic 转成普通返回值。
func recoverDemo() (msg string) {
	defer func() {
		if r := recover(); r != nil {
			msg = fmt.Sprintf("从 panic 恢复：%v", r)
		}
	}()
	panic("故意 panic 一下")
}
