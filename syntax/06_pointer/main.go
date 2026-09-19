// 语法点：指针与值传递——Go 只有值传递；指针是为了「能改到原值」，不是为了省拷贝。
//
// 运行：go run ./syntax/06_pointer
package main

import "fmt"

// Point 用来演示 &T{} 取地址。
type Point struct {
	X, Y int
}

func main() {
	fmt.Println("=== 取地址与解引用 ===")
	x := 42
	// & 取地址，p 的类型是 *int
	p := &x
	fmt.Printf("x = %d, p = %v, *p = %d, p 的类型 = %T\n", x, p, *p, p)

	// * 解引用，改的是 x 本身
	*p = 100
	fmt.Println("执行 *p = 100 之后，x =", x)

	// 指针的零值是 nil，解引用 nil 指针会 panic。
	var nilPtr *int
	fmt.Println("nil 指针 =", nilPtr, "，是否等于 nil：", nilPtr == nil)
	//	fmt.Println(*nilPtr) // panic: runtime error: invalid memory address or nil pointer dereference

	fmt.Println()
	fmt.Println("=== 值传递 vs 指针传递 ===")
	n := 10
	addByValue(n)
	fmt.Println("addByValue(n) 之后 n =", n, "（没变：函数拿到的是副本）")

	addByPointer(&n)
	fmt.Println("addByPointer(&n) 之后 n =", n, "（变了：函数通过指针改到了原变量）")

	fmt.Println()
	fmt.Println("=== 为什么说 Go 只有值传递 ===")
	// addByPointer 也是值传递：它拷贝的是「指针这个值」（一个地址），
	// 只不过通过这个地址能找到原来的变量，所以能改到原值。
	// 换句话说，Go 里没有引用传递，只有「值传递 + 指针」。
	fmt.Println("addByValue 收到的是 int 的副本，addByPointer 收到的是 *int 的副本（地址）。")

	fmt.Println()
	fmt.Println("=== 切片和 map 传的是描述符 ===")
	s := []int{1, 2, 3}
	modifySliceElem(s) // 改元素：调用方看得见
	fmt.Println("modifySliceElem(s) 之后 s =", s, "（元素被改了）")

	reassignSlice(s) // 只改函数内的切片头：调用方看不见
	fmt.Println("reassignSlice(s) 之后 s =", s, "（长度没变，仍是 3）")
	// 注意：s 的 len 和 cap 都是 3（字面量切片的 cap 就是元素个数），
	// 所以 append 一定扩容到新数组，原切片连元素都不会被改。

	m := map[string]int{"a": 1}
	modifyMap(m)
	fmt.Println("modifyMap(m) 之后 m =", m, "（map 元素被改了）")
	// 为什么：切片变量本身是个小结构体 {ptr, len, cap}，map 变量本身是个指针。
	// 传参时拷贝的是这个「描述符」，但描述符里的 ptr 仍指向同一块数据，
	// 所以改元素调用方看得见；而给函数内的切片变量整体重新赋值（或 append
	// 触发扩容），只改了副本里的描述符，调用方看不到。
	// 想让函数「换掉整个切片」，得传 *[]int。

	replaceSlice(&s)
	fmt.Println("replaceSlice(&s) 之后 s =", s, "（传 *[]int 才能真正换掉）")

	fmt.Println()
	fmt.Println("=== new(T) 与 &T{} ===")
	// new(int) 分配一块 int 内存，置为零值，返回 *int。
	// 下面这行 *pi 是 0
	pi := new(int)
	fmt.Printf("new(int)：*pi = %d, pi = %v\n", *pi, pi)
	// 赋值后 *pi 是 5
	*pi = 5
	fmt.Println("赋值后 *pi =", *pi)

	// &T{} 是更常用的写法：一步完成「分配 + 初始化」，得到 *T。
	p2 := &Point{X: 1, Y: 2}
	fmt.Printf("&Point{1,2} = %v, (*p2).X = %d, 语法糖 p2.X = %d\n", *p2, (*p2).X, p2.X)
	// Go 允许对结构体指针直接用 . 访问字段，编译器会自动解引用，不用写 (*p2).X。

	// 结构体指针传参：函数内改字段，调用方看得见。
	moveRight(p2)
	fmt.Println("moveRight(p2) 之后 p2 =", *p2)

	// 对比：值传参会拷贝整个结构体，函数内改了没用。
	p3 := Point{X: 0, Y: 0}
	moveRightByValue(p3)
	fmt.Println("moveRightByValue(p3) 之后 p3 =", p3, "（没变）")

	// 指针有明确的类型，*int 和 *float64 不能互相赋值。
	//	var bad *float64 = p // 编译失败：cannot use p (variable of type *int) as *float64 value

	// 不能对不可寻址的值取地址（map 元素、字面量的字段都属于这类）。
	//	_ = &m["a"]              // 编译失败：cannot take the address of m["a"]
	//	_ = &Point{X: 1, Y: 2}.X // 编译失败：cannot take the address of Point{...}.X
}

// addByValue 收到 int 的副本，改不动调用方的变量。
func addByValue(n int) {
	n++
}

// addByPointer 收到 *int（地址的副本），可以改到调用方的变量。
func addByPointer(n *int) {
	*n++
}

// modifySliceElem 改切片元素：底层数组是共享的，调用方看得见。
func modifySliceElem(s []int) {
	s[0] = 100
}

// reassignSlice 只给函数内的切片变量重新赋值，调用方看不到。
func reassignSlice(s []int) {
	// 传进来的 s 是描述符的副本，append 的返回值赋给副本后函数就返回了，
	// 调用方的切片变量完全没被碰到（它的 len 仍是 3）。
	// 注意：append 的返回值必须接住，裸写 append(s, 999) 会编译失败
	// （append(s, 999) (value of type []int) is not used）。
	// 这里用 _ 显式丢弃，正好演示「不把 append 的结果写回去，这次追加就白做了」。
	// 如果 s 还有剩余容量，append 会写进共享的底层数组；本例 s 的 cap 已经用满，
	// 所以 append 必然分配新数组，原切片一个字节都没动。
	_ = append(s, 999)
}

// replaceSlice 通过 *[]int 真正替换调用方的切片。
func replaceSlice(s *[]int) {
	*s = append(*s, 4, 5)
}

// modifyMap 改 map 元素：map 变量内部就是指针，调用方看得见。
func modifyMap(m map[string]int) {
	m["a"] = 100
	m["b"] = 2
}

// moveRight 通过结构体指针修改字段。
func moveRight(p *Point) {
	p.X++
}

// moveRightByValue 收到结构体副本，改不动调用方。
func moveRightByValue(p Point) {
	p.X++
}
