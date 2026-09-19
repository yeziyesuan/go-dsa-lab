// 语法点：变量、常量、基本类型与类型转换——Go 的变量有明确类型，且不会替你做隐式转换。
//
// 运行：go run ./syntax/01_types
package main

import "fmt"

// Global 演示包级变量：零值是 0，声明在函数外面。
var Global int

// 常量块 + iota：iota 在每个 ConstSpec 行递增，从 0 开始。
// 下面这组注释写在声明行上方，是为了避免依赖 gofmt 的注释列对齐细节。
const (
	// StatusTodo 等于 0
	StatusTodo = iota
	// StatusDoing 等于 1
	StatusDoing
	// StatusDone 等于 2
	StatusDone
	// statusSkipped 等于 3；首字母小写 => 包外不可见
	statusSkipped
)

// iota 参与表达式：每行左移一位，得到一个 KB/MB/GB 序列。
const (
	// 第一个 iota 是 0，用 _ 丢掉，让 KB 从 1 << 10 开始
	_ = iota
	// KB = 1 << 10 = 1024
	KB = 1 << (10 * iota)
	// MB = 1 << 20
	MB
	// GB = 1 << 30
	GB
)

// Week 是一个无类型字符串常量，可以直接参与字符串拼接。
const Week = "W1"

func main() {
	fmt.Println("=== 变量声明的三种写法 ===")
	// 包级变量的零值就是 0，这里打印出来确认。
	fmt.Println("包级变量 Global =", Global)

	var a int = 10 // 完整写法：var + 类型 + 初值
	var b = 20     // 省略类型，由编译器推断为 int
	c := 30        // 短声明（只能在函数体内用，:= 左边至少有一个新变量）
	fmt.Println("a =", a, "b =", b, "c =", c)

	// 一次声明多个变量
	var x, y int = 1, 2
	p, q := "left", "right"
	fmt.Println("x, y =", x, y, "| p, q =", p, q)

	// 短声明的「至少一个新变量」规则：err 已存在，v 是新的，所以 := 合法。
	v, err := 100, error(nil)
	fmt.Println("v, err =", v, err)

	fmt.Println()
	fmt.Println("=== 常量与 iota ===")
	fmt.Println("StatusTodo / StatusDoing / StatusDone =",
		StatusTodo, StatusDoing, StatusDone)
	fmt.Println("statusSkipped =", statusSkipped)
	fmt.Println("KB / MB / GB =", KB, MB, GB)
	fmt.Println("Week =", Week)

	fmt.Println()
	fmt.Println("=== 基本类型的零值 ===")
	// 零值规则：数值类型是 0，bool 是 false，string 是 ""，指针/切片/map/接口是 nil。
	var (
		zi  int
		zi8 int8
		zi6 int64
		zf  float64
		zs  string
		zb  bool
		zr  rune // rune 就是 int32，零值是 0，即 NUL 字符
	)
	fmt.Printf("int=%v int8=%v int64=%v float64=%v\n", zi, zi8, zi6, zf)
	// string 的零值用 %q 打印，才能看出它是空串而不是空格。
	fmt.Printf("string=%q bool=%v rune=%v (%T)\n", zs, zb, zr, zr)

	fmt.Println()
	fmt.Println("=== 类型转换：必须显式写 ===")
	var i32 int32 = 7
	i64 := int64(i32) // 窄转宽也要写出来
	fmt.Println("int32(7) -> int64 =", i64)

	// 整数除法会截断：3/2 == 1。想拿到小数结果，得先把操作数转成 float64。
	fmt.Println("3 / 2            =", 3/2, "（整数除法，直接截断）")
	fmt.Println("float64(3) / 2   =", float64(3)/2)

	// 注意：float64(3)/2 是「常量表达式」，它的值是常量 1.5。
	// 常量转换必须能精确表示，所以下面这样直接转 int 会编译失败：
	//	int(float64(3) / 2) // 编译失败：cannot convert float64(3) / 2 (constant 1.5 of type float64) to type int
	// 为什么：常量 1.5 转 int 会丢信息，Go 在编译期就直接拒绝，而不是悄悄截断。
	// 想让截断真的发生，先把它落到变量里，再转换（这时是运行期转换，规则是「向零取整」）：
	half := float64(3) / 2
	fmt.Println("int(half)        =", int(half), "（half 是变量，值是 1.5，转成 int 得 1）")

	// 变量转换不会四舍五入，而是丢掉小数部分（向零截断）。
	f2 := 2.9
	f3 := -2.9
	fmt.Printf("int(f2) = %d（f2 = 2.9，向零截断得到 2）\n", int(f2))
	fmt.Printf("int(f3) = %d（f3 = -2.9，负数也向零截断，不是 -3）\n", int(f3))
	// 同样是常量转换的坑：下面两行也会编译失败，必须先赋给变量再转。
	//	int(2.9)  // 编译失败：cannot convert 2.9 (untyped float constant) to type int
	//	int(-2.9) // 编译失败：cannot convert -2.9 (untyped float constant) to type int

	// 数值与字符串之间：string(n) 得到的是「码点是 n 的那个字符」，不是十进制数字串。
	// 用变量 n65 而不是常量 65，是为了演示运行期的类型转换；
	// 注意 go vet 对「int -> string 转换」一律会告警（不管常量还是变量）：
	//   conversion from int to string yields a string of one rune,
	//   not a string of digits (did you mean fmt.Sprint(x)?)
	// 因为这么写极易被误读成得到 "65"。这里用 %c 达到同样效果又不触发告警，
	// 真正「数字转十进制字符串」要用 strconv.Itoa（见 ./syntax/12_stdlib）。
	n65 := 65
	fmt.Printf("string(n65) 会得到字母 A，等价于 %%c：%c；想要 \"65\" 得用 strconv.Itoa\n", n65)

	// string 与 []byte / []rune 的互转：string 底层是只读的字节序列。
	word := "Go语言"
	bs := []byte(word)
	rs := []rune(word)
	fmt.Printf("len(%q) = %d 字节，[]rune 长度 = %d 个字符\n", word, len(bs), len(rs))
	fmt.Printf("[]byte 前两个元素 = %v，转回 string = %q\n", bs[:2], string(bs[:2]))

	fmt.Println()
	fmt.Println("=== Go 没有隐式转换（下面这些写法会编译失败）===")
	// 为什么：Go 的类型系统把 int 和 float64 当作完全不同的类型，
	// 不存在别的语言那种「自动提升」。下面每一行删掉注释都会编译报错。
	//
	//	var f float64 = 1          // 编译失败：cannot use 1 (untyped int constant) as float64 value
	//	f = f + a                  // 编译失败：invalid operation: mismatched types float64 and int
	//	i32 = i64                  // 编译失败：cannot use i64 (variable of type int64) as int32 value
	//	var s string = 'A'         // 编译失败：cannot use 'A' (untyped rune constant) as string value
	//	if a { }                   // 编译失败：non-boolean condition，Go 里 1 不等于 true
	fmt.Println("上面 5 行注释掉的代码，任意一行取消注释都会编译失败。")

	// 正确写法：老老实实转换。
	var f float64 = float64(a)
	fmt.Println("float64(a) =", f, "| 相加要同类型：", f+float64(b))
}
