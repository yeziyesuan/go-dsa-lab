// 语法点：结构体与方法——值接收者改不动原对象，指针接收者才能改；组合用嵌入，不用继承。
//
// 运行：go run ./syntax/07_struct_method
package main

import "fmt"

// Point 是二维坐标点。字段首字母大写 => 包外可见。
type Point struct {
	X, Y int
}

// Base 用来演示结构体嵌入（组合）：它被嵌入到 Derived 里。
type Base struct {
	ID   int
	Kind string
}

// String 实现 fmt.Stringer 接口，fmt 打印 Base 时就会用这个格式。
func (b Base) String() string {
	return fmt.Sprintf("Base{ID:%d, Kind:%s}", b.ID, b.Kind)
}

// Derived 嵌入 Base：Base 是匿名字段，它的字段和方法会被「提升」到 Derived。
type Derived struct {
	// Base 是匿名字段，字段名就是类型名 Base
	Base
	// Name 和 Weight 是 Derived 自己的字段
	Name   string
	Weight float64
}

func main() {
	fmt.Println("=== 定义结构体与字面量 ===")
	p := Point{X: 1, Y: 2} // 带字段名的字面量（推荐，加字段也不会错位）
	q := Point{3, 4}       // 按顺序的字面量（字段多了容易写错）
	var z Point            // 零值：每个字段都是各自类型的零值
	fmt.Printf("p = %+v, q = %v, 零值 z = %+v\n", p, q, z)
	fmt.Printf("p 的类型 = %T，字段：X=%d Y=%d\n", p, p.X, p.Y)

	fmt.Println()
	fmt.Println("=== 值接收者 vs 指针接收者 ===")
	// 值接收者：方法拿到的是结构体副本，改不动原对象。
	p.MoveByValue(10, 10)
	fmt.Println("MoveByValue(10,10) 之后 p =", p, "（X、Y 没变）")
	// 为什么改不动：func (p Point) MoveByValue 里的 p 是调用者传进来的拷贝，
	// 改 p.X 只改了这份拷贝，函数返回就丢了。

	// 指针接收者：方法拿到的是地址，能改到原对象。
	p.MoveByPointer(10, 10)
	fmt.Println("MoveByPointer(10,10) 之后 p =", p, "（X、Y 变了）")
	// 语法糖：(p).MoveByPointer(...) 会自动改成 (&p).MoveByPointer(...)，
	// 前提是 p 可寻址。不可寻址的值没法调用指针接收者方法。

	// 因为 p 可寻址，Go 自动取地址；下面这样对临时值调用就会编译失败。
	//	Point{1, 1}.MoveByPointer(1, 1) // 编译失败：cannot call pointer method on Point{1, 1}

	fmt.Println()
	fmt.Println("=== 值接收者接收结构体 vs 指针接收者返回结构体 ===")
	scale := &Point{X: 2, Y: 3}
	fmt.Println("调用前 scale =", *scale)
	scale.Scale(3) // 指针接收者，原地放大
	fmt.Println("Scale(3) 之后 scale =", *scale)

	// 方法集规则（非常重要，接口那一节会用到）：
	//   T  的方法集 = 值接收者方法
	//   *T 的方法集 = 值接收者方法 + 指针接收者方法
	// 所以 *Point 能满足的接口，Point 不一定能满足。
	fmt.Println("Point 的方法集只含值接收者方法；*Point 的方法集还要加上指针接收者方法。")

	fmt.Println()
	fmt.Println("=== 结构体嵌入（组合，不是继承）===")
	d := Derived{
		Base:   Base{ID: 1, Kind: "node"},
		Name:   "leaf",
		Weight: 2.5,
	}
	// 提升字段：不用写 d.Base.ID，直接 d.ID 就行。
	fmt.Println("d.ID =", d.ID, "，d.Kind =", d.Kind, "，d.Name =", d.Name)
	// 也可以显式访问。
	fmt.Println("显式访问 d.Base.ID =", d.Base.ID)

	// 提升方法：Base 的 String() 被提升到 Derived。
	fmt.Println("d.String() =", d.String())

	// 嵌入的是「has-a」关系：Derived 有一个 Base，而不是「是一个」Base。
	// Go 没有继承，也没有方法重写；想改行为就在外层重新定义同名方法。
	// 外层同名方法会遮蔽（shadow）内层方法，但不会多态分派。

	// 一般规则：%v 只给值，%+v 带字段名。
	// 但这里有个例外：因为 Base 有 String() 且被提升到了 Derived，Derived 也满足
	// fmt.Stringer，而 fmt 对实现了 String() 的类型会优先调用它 —— 只要动词是针对
	// 字符串的（%s %q %x %X）或者是「不带 # 的 %v」。%+v 也不带 #，所以它同样会调用
	// String()。结果：下面两行打印出来都是 "Base{ID:1, Kind:node}"，都看不到 Name/Weight。
	// 想看字段得换个不含 String() 的类型，或者用 %#v（%#v 会走 GoStringer，不走 Stringer）。
	fmt.Printf("d = %v（触发了提升来的 String()）\n", d)
	fmt.Printf("d = %+v（同样触发了 String()，不是字段列表）\n", d)

	fmt.Println()
	fmt.Println("=== 实现 fmt.Stringer ===")
	// String() string 这个方法签名就是 fmt.Stringer 接口。
	// 实现了它，fmt 的 %v / %s / Println 都会自动调用，不再打印裸字段。
	b := Base{ID: 7, Kind: "root"}
	fmt.Println("Println(b) =", b)
	fmt.Printf("v=%v s=%s q=%q\n", b, b, b)
	fmt.Println("fmt 会优先用 String()，所以这里看不到 struct 的默认格式。")

	// 变量存指针时同样生效（Base 的 String 是值接收者，*Base 的方法集包含它）。
	bp := &Base{ID: 8, Kind: "child"}
	fmt.Println("Println(bp) =", bp)

	// fmt.Stringer 长这样，标准库里已经定义好了，这里只是注释出来对照：
	//	type Stringer interface { String() string }

	fmt.Println()
	fmt.Println("=== 结构体嵌套与匿名结构体 ===")
	// 具名字段嵌套：has-a，但字段不会被提升。
	// 注意 Line 定义在包级（见文件末尾）：方法只能挂在包级类型上，
	// 如果把 type Line struct{...} 写在 main 函数里，就没法给它写 distSq 方法。
	l := Line{Start: Point{0, 0}, End: Point{1, 1}}
	fmt.Printf("Line = %+v，长度平方 = %d\n", l, l.distSq())

	// 匿名结构体：一次性使用、不需要命名类型时很方便。
	tag := struct {
		Key   string
		Value int
	}{"count", 3}
	fmt.Println("匿名结构体 tag =", tag, "，tag.Key =", tag.Key)

	fmt.Println()
	fmt.Println("=== 结构体传参：值传参是整份拷贝 ===")
	big := Point{100, 200}
	modifyPointCopy(big)
	fmt.Println("modifyPointCopy(big) 之后 big =", big, "（副本被改，原值不变）")
	modifyPointPtr(&big)
	fmt.Println("modifyPointPtr(&big) 之后 big =", big, "（原值被改）")
	// 结构体是值类型，传参的拷贝成本 = 结构体自身大小。
	// 字段多的结构体通常用 *T 传参/做接收者，既省拷贝又能改原值。
	fmt.Println("Point 只有两个 int，拷贝成本很低；字段一多就该考虑用指针。")
}

// MoveByValue 是值接收者方法：改的是副本，调用方看不到。
func (p Point) MoveByValue(dx, dy int) {
	p.X += dx
	p.Y += dy
}

// MoveByPointer 是指针接收者方法：改的是原对象。
func (p *Point) MoveByPointer(dx, dy int) {
	p.X += dx
	p.Y += dy
}

// Scale 是指针接收者方法，把坐标按倍数放大。
func (p *Point) Scale(f int) {
	p.X *= f
	p.Y *= f
}

// Line 是由两个 Point 组合成的线段：具名字段嵌套（has-a）。
// 定义在包级，才能给它挂 distSq 方法（方法不能挂在函数内定义的类型上）。
type Line struct {
	Start Point
	End   Point
}

// distSq 计算 Line 两端点距离的平方（避免引入 math 包）。
func (l Line) distSq() int {
	dx := l.End.X - l.Start.X
	dy := l.End.Y - l.Start.Y
	return dx*dx + dy*dy
}

// modifyPointCopy 收到结构体副本，改不动调用方。
func modifyPointCopy(p Point) {
	p.X = -1
}

// modifyPointPtr 收到结构体指针，可以改到调用方。
func modifyPointPtr(p *Point) {
	p.X = -1
}
