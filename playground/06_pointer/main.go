// 语法点：指针与值传递——Go 只有值传递；指针是为了「能改到原值」，不是为了省拷贝。
//
// 运行：go run ./playground/06_pointer
// 约定：凭记忆手敲（卡住才回去瞄 syntax/06_pointer/main.go）；本文件只给骨架与 TODO，
// 答案自己写；随时能编译（写完一段就跑一次）。
//
// ⛳ 10/9（周五）做这 4 组 —— 每组都要看得见一个"差别"：
//
//	G1 = TODO 1            取地址 / 解引用：*p = 100 之后 x 真的变了
//	G2 = TODO 3 + 4        值传递 vs 指针传递（**本点核心**）；一句话想清楚"Go 只有值传递"
//	G3 = TODO 5 + 6 + 7    切片：改元素看得见 / 重新赋值看不见 / 传 *[]int 才换得掉
//	G4 = TODO 9 + 10       结构体：指针能改字段，值拷贝改不动
//	排法：08:15–10:10（上机锚 2h）做 G1+G2；15:30–19:30（弹性块）做 G3+G4
//	判据：go run ./playground/06_pointer → 这 8 行期望值全中
//
// ⏳ 剩下 3 处（TODO 2 nil 指针 / TODO 8 map / §6 三条编译错误实验）→ 10/9 弹性块的尾巴或 W4。
package main

import "fmt"

// Point 用来演示「结构体值 vs 结构体指针」的差别。
type Point struct {
	X, Y int
}

func main() {
	fmt.Println("=== 1. 取地址与解引用 ===")
	// TODO 1：自己写三行 —— x := 42 → p := &x
	//         打印 p（地址）、*p（42）、%T（*int）；再 *p = 100，打印 x（期望 100）
	// TODO 2：nil 指针 —— var np *int；打印 np 与 np == nil（期望 <nil> true）
	//         想看 panic：把 fmt.Println(*np) 取消注释跑一次，报错原话抄进笔记
	fmt.Println("TODO 1：（自己写）*p = 100 之后 x = ...（期望 100：改的是 x 本身）")
	fmt.Println("TODO 2：（自己写）np = ... ｜ np == nil = ...（期望 <nil> true）")

	fmt.Println()
	fmt.Println("=== 2. 值传递 vs 指针传递（本点核心） ===")
	n := 10
	addByValue(n)
	fmt.Println("TODO 3：addByValue(n) 之后 n =", n, "（期望 10：函数拿到的是副本）")
	addByPointer(&n)
	fmt.Println("TODO 4：addByPointer(&n) 之后 n =", n, "（期望 11：通过指针改到了原变量）")

	fmt.Println()
	fmt.Println("=== 3. 为什么说 Go 只有值传递 ===")
	// 想清楚一句话再往下写：addByPointer 也是值传递，它拷贝的是「地址」这个值。
	// 面试常问：Go 有引用传递吗？——答案 + 理由（写进笔记，别只记结论）

	fmt.Println()
	fmt.Println("=== 4. 切片 / map 传的是什么 ===")
	s := []int{1, 2, 3}
	modifySliceElem(s)
	fmt.Println("TODO 5：modifySliceElem(s) 之后 s =", s, "（期望 [100 2 3]：改元素看得见）")
	reassignSlice(s)
	fmt.Println("TODO 6：reassignSlice(s) 之后 len(s) =", len(s), "（期望 3：只改副本，看不见）")
	replaceSlice(&s)
	fmt.Println("TODO 7：replaceSlice(&s) 之后 s =", s, "（期望 [100 2 3 4 5]：传 *[]int 才换得掉）")
	m := map[string]int{"a": 1}
	modifyMap(m)
	fmt.Println("TODO 8：modifyMap(m) 之后 m =", m, "（期望 map[a:100 b:2]：map 内部就是一个指针）")

	fmt.Println()
	fmt.Println("=== 5. 结构体：值 vs 指针 ===")
	p2 := &Point{X: 1, Y: 2}
	moveRight(p2)
	fmt.Println("TODO 9：moveRight(p2) 之后 p2 =", *p2, "（期望 {2 2}：指针能改字段）")
	p3 := Point{}
	moveRightByValue(p3)
	fmt.Println("TODO 10：moveRightByValue(p3) 之后 p3 =", p3, "（期望 {0 0}：拷贝，改不动）")

	fmt.Println()
	fmt.Println("=== 6. 故意写错一次（把编译器原话抄进笔记） ===")
	// ① 不可寻址：_ = &m["a"]                → cannot take the address of m["a"]
	// ② 类型不匹配：var bad *float64 = p     → cannot use p (variable of type *int) as *float64
	// ③ 裸 append：append(s, 999)            → append(s, 999) (value of type []int) is not used
	fmt.Println("（三条都在注释里：亲手取消注释跑一次，再把报错原话抄下来）")
}

// TODO 1：取地址 + 解引用（在 main 里写，不用函数）
// TODO 2：nil 指针的零值与「解引用 nil 会 panic」

// TODO 3：值传递——参数是 int 的副本，n++ 改不到调用方
func addByValue(n int) {
}

// TODO 4：指针传递——参数是 *int，*n++ 能改到调用方
func addByPointer(n *int) {
}

// TODO 5：改切片元素（底层数组共享，调用方看得见）
func modifySliceElem(s []int) {
}

// TODO 6：只改函数内的切片变量（append 的结果不写回调用方，或有容量时写进了共享数组的哪个位置）
func reassignSlice(s []int) {
}

// TODO 7：通过 *[]int 真正替换调用方的切片：*s = append(*s, 4, 5)
func replaceSlice(s *[]int) {
}

// TODO 8：改 map 元素：m["a"] = 100; m["b"] = 2
func modifyMap(m map[string]int) {
}

// TODO 9：结构体指针改字段（p.X++）
func moveRight(p *Point) {
}

// TODO 10：结构体值——参数是副本，p.X++ 改不到调用方
func moveRightByValue(p Point) {
}
