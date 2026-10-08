// 语法点：函数——多返回值、命名返回值、可变参数、闭包、defer（LIFO / 改返回值 / recover）。
//
// 运行：go run ./playground/05_func
// 约定：凭记忆手敲（卡住才回去瞄 syntax/05_func/main.go）；本文件只给骨架与 TODO，
// 答案自己写；随时保持能编译（写完一段就跑一次）。
//
// ⛳ 今晚（2026/10/8）只做这 5 处 —— 做完就算重启成功，其余的今晚不许碰：
//
//	TODO 1   divmod        多返回值 + error（要补 errors 的 import）
//	TODO 5   countUp       命名返回值 + defer 改返回值（先想清楚：return 之后 defer 才跑）
//	TODO 6   sum           可变参数；sum() / sum(1,2,3) / sum(nums...) 三种都要跑
//	TODO 7   counter       闭包；两个计数器实例要互不影响
//	TODO 11  recoverDemo   defer + recover 把 panic 变成普通返回值
//	判据：go run ./playground/05_func → 上面这 5 行括号里的期望值全中；文件顶部补 3 行反直觉笔记
//
// ⏳ 剩下 6 处（TODO 2 / 3 / 4 / 8 / 9 / 10）都是这 5 处的变形 → **留到 W4 弹性块再补**。
//
//	今晚顺手做了不算超额，叫抢跑——抢跑吃的是明天的额度。
package main

import "fmt"

func main() {
	fmt.Println("=== 1. 多返回值（值 + error） ===")
	// TODO 1：divmod(a,b) 返回 (商, 余数, error)；b == 0 时返回错误（要用 errors.New，记得补 import）
	// TODO 2：twoValues() 返回两个 int，用 _ 丢掉不想要的那个
	q, r, err := divmod(17, 5)
	fmt.Println("TODO 1：divmod(17,5) =", q, r, err, "（期望 3 2 <nil>）")

	fmt.Println()
	fmt.Println("=== 2. 命名返回值 ===")
	// TODO 3：namedSum(a,b) 用命名返回值 + 裸 return（只写 return，不写 return total）
	// TODO 4：namedQuotient(a,b) 命名返回值 (q int, err error)，b == 0 时走 err 分支
	// TODO 5：countUp(n) 命名返回值 + defer 里 ret++（先想清楚：return n 之后 defer 还能改到返回值吗？）
	fmt.Println("TODO 3：namedSum(3,4) =", namedSum(3, 4), "（期望 7）")
	fmt.Println("TODO 5：countUp(5) =", countUp(5), "（期望 6，不是 5）")

	fmt.Println()
	fmt.Println("=== 3. 可变参数 ===")
	// TODO 6：sum(nums ...int) 求和。三件事都要做：
	//         ① sum()、sum(1,2,3)、sum(slice...) 都要能跑；
	//         ② 故意把 sum(nums)（不加 ...）写一次，看编译器报什么错，抄进笔记；
	//         ③ 记住：nums 在函数内就是一个 []int
	nums := []int{10, 20, 30}
	fmt.Println("TODO 6：sum() =", sum(), "｜sum(1,2,3) =", sum(1, 2, 3), "｜sum(nums...) =", sum(nums...), "（期望 0 6 60）")

	fmt.Println()
	fmt.Println("=== 4. 闭包 ===")
	// TODO 7：counter() 返回一个闭包，每次调用返回递增的计数；
	//         再开一个独立实例（next2 := counter()），验证两个计数器互不影响
	// TODO 8：循环变量捕获：for i := range 3 { 把 func(){fmt.Print(i," ")} 存起来 }，再依次调用
	//         想清楚：Go 1.22 起为什么输出 0 1 2，而 1.21 及以前是 3 3 3
	next := counter()
	fmt.Println("TODO 7：next() 连调三次 =", next(), next(), next(), "（期望 1 2 3）")

	fmt.Println()
	fmt.Println("=== 5. defer ===")
	// TODO 9：deferOrder() —— 循环里注册 3 个 defer，验证「后注册的先执行」(LIFO)
	//         顺带想清楚：defer 的参数是注册时求值，还是函数返回时求值？
	// TODO 10：addThenDouble(3) 命名返回值 + defer 里 *=2 → 期望 8，不是 4
	// TODO 11：recoverDemo() defer + recover 把 panic("故意 panic 一下") 变成普通返回值
	deferOrder()
	fmt.Println("TODO 10：addThenDouble(3) =", addThenDouble(3), "（期望 8，不是 4）")
	fmt.Println("TODO 11：recoverDemo() =", recoverDemo(), "（期望不带 panic 正常返回）")
}

// TODO 1：返回 (商, 余数, error)；b == 0 时返回 errors.New("除数不能为 0")
func divmod(a, b int) (int, int, error) {
	return 0, 0, nil
}

// TODO 2：返回两个 int，用来练习用 _ 丢弃返回值
func twoValues() (int, int) {
	return 0, 0
}

// TODO 3：命名返回值 total + 裸 return
func namedSum(a, b int) (total int) {
	return 0
}

// TODO 4：命名返回值 (q int, err error)；b == 0 时只赋值给 err 再 return
func namedQuotient(a, b int) (q int, err error) {
	return 0, nil
}

// TODO 5：命名返回值 ret + defer 里 ret++（关键是 defer 在 return 赋值之后才执行）
func countUp(n int) (ret int) {
	return n
}

// TODO 6：可变参数求和；nums 在函数内就是 []int，直接 range 它
func sum(nums ...int) int {
	return 0
}

// TODO 7：返回闭包，每次调用 +1；每次调用 counter() 都得到一份独立的 count
func counter() func() int {
	return func() int { return 0 }
}

// TODO 9：循环注册 3 个 defer，验证 LIFO；函数体最后打印一行提示
func deferOrder() {
}

// TODO 10：命名返回值 + defer 把结果 *=2（先 return n+1，再被 defer 翻倍）
func addThenDouble(n int) (result int) {
	return 0
}

// TODO 11：defer + recover，把 panic("故意 panic 一下") 转成返回值
func recoverDemo() (msg string) {
	return ""
}
