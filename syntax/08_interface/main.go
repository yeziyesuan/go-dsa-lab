// 语法点：接口——隐式实现、方法集决定谁能满足接口、类型断言与 type switch、any。
//
// 运行：go run ./syntax/08_interface
package main

import (
	"fmt"
	"math"
)

// Shape 是接口：只要一个类型有 Area() float64，就自动满足它（隐式实现）。
type Shape interface {
	Area() float64
}

// Stringer 是标准库 fmt.Stringer 的签名，这里再定义一次用来对照。
// 注意：Go 的接口是结构化的，两个接口只要方法集相同就完全等价，可以互相赋值。
type Stringer interface {
	String() string
}

// Rect 用「值接收者」实现 Area：Rect 和 *Rect 都满足 Shape。
type Rect struct {
	W, H float64
}

// Area 实现 Shape 接口（值接收者）。
func (r Rect) Area() float64 {
	return r.W * r.H
}

// Scale 是指针接收者方法，不影响值接收者那边的接口实现。
func (r *Rect) Scale(f float64) {
	r.W *= f
	r.H *= f
}

// Circle 用「指针接收者」实现 Area：只有 *Circle 满足 Shape，Circle 不满足。
type Circle struct {
	R float64
}

// Area 实现 Shape 接口（指针接收者）。
func (c *Circle) Area() float64 {
	return math.Pi * c.R * c.R
}

// Named 用来演示 Stringer 的隐式实现。
type Named struct {
	Name string
}

// String 实现 Stringer（也实现了标准库的 fmt.Stringer）。
func (n Named) String() string {
	return "Named(" + n.Name + ")"
}

func main() {
	fmt.Println("=== 隐式实现：不用写 implements ===")
	// Rect / Circle / Named 都没有写任何 "implements Shape"，
	// 只要方法集包含 Area() float64，就自动是 Shape。
	var s Shape = Rect{W: 3, H: 4}
	fmt.Printf("s = %v, s.Area() = %.2f\n", s, s.Area())
	// 编译期断言：把 *Rect 赋给 Shape 变量，不行的话直接编译失败。
	var _ Shape = (*Rect)(nil)
	var _ Shape = (*Circle)(nil)
	fmt.Println("Rect、*Rect、*Circle 都满足 Shape（这行上面的编译期断言就是证据）。")

	fmt.Println()
	fmt.Println("=== 值实现 vs 指针实现：方法集差异 ===")
	// 方法集规则：
	//   T  的方法集 = 值接收者方法
	//   *T 的方法集 = 值接收者方法 + 指针接收者方法
	// Circle 的 Area 是指针接收者 => Circle 的方法集里没有 Area => Circle 不满足 Shape。
	// 把下面两行取消注释会编译失败：
	//
	//	var bad Shape = Circle{R: 1}   // 编译失败：Circle does not implement Shape
	//	                                // (method Area has pointer receiver)
	//
	// 反过来，Rect 的 Area 是值接收者 => Rect 和 *Rect 都满足 Shape。
	var sh1 Shape = Rect{W: 1, H: 1}  // 存值：动态类型是 Rect
	var sh2 Shape = &Rect{W: 1, H: 1} // 存指针：动态类型是 *Rect
	fmt.Printf("值存接口：%v -> Area=%.2f\n", sh1, sh1.Area())
	fmt.Printf("指针存接口：%v -> Area=%.2f\n", sh2, sh2.Area())

	// 为什么接口里存值就调用不到指针接收者方法：
	// 接口变量内部是 (类型, 数据) 两个字段。存 Circle 的值时，接口里放的是
	// 一份不可寻址的 Circle 拷贝，Go 没法对它自动取地址，所以指针接收者方法
	// 不在方法集里。存 *Circle 时，指针本身就是数据，方法集完整。
	// 结论：想让「值和指针都能满足接口」，就把接收者写成值接收者；
	//       想避免拷贝、或需要修改接收者，就用指针接收者，然后统一用 *T 存接口。

	var sh3 Shape = &Circle{R: 2} // 只能用指针
	fmt.Printf("Circle 必须用指针：%v -> Area=%.2f\n", sh3, sh3.Area())

	fmt.Println()
	fmt.Println("=== 接口变量存值 vs 存指针：类型断言结果不同 ===")
	fmt.Printf("sh1 的动态类型：%T\n", sh1)
	fmt.Printf("sh2 的动态类型：%T\n", sh2)
	fmt.Printf("sh3 的动态类型：%T\n", sh3)
	// Rect 和 *Rect 是两个不同的动态类型，断言时要写对。

	fmt.Println()
	fmt.Println("=== 类型断言 v, ok := s.(Rect) ===")
	// 单返回值写法在类型不对时会 panic；双返回值写法返回 ok，永远不会 panic。
	if v, ok := sh1.(Rect); ok {
		fmt.Println("sh1 确实是 Rect，W =", v.W, "H =", v.H)
	} else {
		fmt.Println("sh1 不是 Rect")
	}
	if v, ok := sh1.(*Rect); ok {
		fmt.Println("这一行不会执行：", v)
	} else {
		fmt.Println("sh1 不是 *Rect（它存的是值，不是指针）")
	}
	// 单值断言失败会 panic，下面这行不能真的执行：
	//	_ = sh1.(*Rect) // panic: interface conversion: main.Shape is main.Rect, not *main.Rect

	// 断言成另一个接口也是合法的：前提是动态类型满足那个接口。
	if st, ok := sh1.(Stringer); ok {
		fmt.Println("sh1 也满足 Stringer：", st.String())
	}

	fmt.Println()
	fmt.Println("=== type switch ===")
	shapes := []Shape{Rect{W: 2, H: 5}, &Rect{W: 1, H: 1}, &Circle{R: 1}}
	for i, sh := range shapes {
		describe(i, sh)
	}

	fmt.Println()
	fmt.Println("=== any 与空接口 ===")
	// any 是 interface{} 的别名（Go 1.18 起），可以装任何类型的值。
	var anything any = 42
	fmt.Printf("anything = %v, 动态类型 = %T\n", anything, anything)
	anything = "hello"
	fmt.Printf("anything = %v, 动态类型 = %T\n", anything, anything)
	anything = []int{1, 2}
	fmt.Printf("anything = %v, 动态类型 = %T\n", anything, anything)

	// 从 any 取回具体值必须用类型断言（编译期不知道动态类型）。
	if n, ok := anything.([]int); ok {
		fmt.Println("断言成 []int 成功，和 =", sumInts(n))
	}

	// 注意接口的 nil 陷阱：装着 nil 指针的接口变量不等于 nil。
	var np *Rect
	var shape Shape = np
	fmt.Println("np == nil：", np == nil)
	fmt.Println("shape == nil：", shape == nil, "（接口里存了 (*Rect, nil)，所以不等于 nil）")
	// 为什么：接口值 = (动态类型, 动态值)。只有当类型和值都是 nil 时，接口才 == nil。
	// 判空要先断言出具体类型再判断，别直接和 nil 比。

	fmt.Println()
	fmt.Println("=== 接口赋值与比较 ===")
	// 接口变量可以互相赋值，只要方法集兼容。
	var st fmt.Stringer = Named{Name: "go"}
	fmt.Println("fmt.Stringer 存 Named：", st) // fmt 自动调用 String()
	//	Named.Name = "x" // 编译失败：cannot assign to Named.Name，字面量不可寻址
	n := Named{Name: "go"}
	n.Name = "golang" // 变量可以改
	st = n
	fmt.Println("改名后：", st)

	// 接口值可比较，但底层类型不可比较时会在运行期 panic。
	fmt.Println("sh1 == sh2：", sh1 == sh2, "（动态类型不同，直接 false）")
	fmt.Println("sh1 == Rect{1, 1}：", sh1 == Rect{W: 1, H: 1})

	// 接口也可以嵌入接口，把方法集合并起来（这里注释出来对照）。
	//	type ReadWriter interface {
	//		Reader
	//		Writer
	//	}
}

// describe 用 type switch 打印接口变量的动态类型和具体值。
func describe(i int, s Shape) {
	switch v := s.(type) {
	case Rect:
		fmt.Printf("[%d] 是 Rect 值：%v，Area = %.2f\n", i, v, v.Area())
	case *Rect:
		fmt.Printf("[%d] 是 *Rect：%v，W = %.1f H = %.1f，Area = %.2f\n",
			i, v, v.W, v.H, v.Area())
	case *Circle:
		fmt.Printf("[%d] 是 *Circle：R = %.1f，Area = %.2f\n", i, v.R, v.Area())
	case nil:
		// nil 接口会命中 nil 这个 case。
		fmt.Printf("[%d] 是 nil 接口\n", i)
	default:
		// 兜底：动态类型是上面都没列出的类型。
		fmt.Printf("[%d] 未知形状：%T\n", i, v)
	}
}

// sumInts 把 []int 求和，供 any 断言后使用。
func sumInts(nums []int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}
