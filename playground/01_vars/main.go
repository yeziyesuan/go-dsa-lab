package main

import "fmt"

var Global int

const (
	StatusTodo = iota
	StatusDoing
	StatusDone
	statusSkipped
)

const (
	_  = iota
	KB = 1 << (10 * iota)
	MB
	GB
)
const week = "w1"

func main() {
	fmt.Println("===变量声明的三种方法===")
	fmt.Println("包级变量 Global =", Global)
	var a int = 10
	var b = 20
	c := 30
	fmt.Println("a = ", a, "b =", b, "c = ", c)
	var x, y int = 1, 2
	p, q := "left", "right"
	fmt.Println("x , y = ", x, y, "|p , q = ", p, q)
	var err error // 零值就是 nil；比写 error(nil) 清楚
	v := 100
	fmt.Println("v,err =", v, err)
	fmt.Println()
	fmt.Println("===常量与iota===")
	fmt.Println("StatusTodo / StatusDoing / StatusDone =", StatusTodo, StatusDoing, StatusDone)
	fmt.Println("StatusSkipped = ", statusSkipped)
	fmt.Println("KB / MB / GB =", KB, MB, GB)
	fmt.Println("week = ", week)
	fmt.Println()
	fmt.Println("===基本类型的零值===")
	var (
		zi  int
		zi8 int8
		zi6 int64
		zf  float64
		zs  string
		zb  bool
		zr  rune
	)
	fmt.Printf("int=%v int8=%v int64=%v float64=%v\n", zi, zi8, zi6, zf)
	fmt.Printf("string=%q bool=%v rune=%v (%T)\n", zs, zb, zr, zr)
	fmt.Println()
	fmt.Println("===类型转换：必须显式写===")
	var (
		zi32 int32 = 7
		zi64 int64 = int64(zi32)
	)
	fmt.Println("int32(7) -> int64 = ", zi64)
	fmt.Println("3/2 = ", 3/2)
	fmt.Println("float64(3)/2 =", float64(3)/2)
	half := float64(3) / 2
	fmt.Println("int(half) =", int(half))
	f2 := 2.9
	f3 := -2.9
	fmt.Println("int(f2) =", int(f2))
	fmt.Println("int(f3) =", int(f3))
	n65 := 65
	fmt.Printf("string(n65)会得到字母%c;想要\"%v\"得用strconv.Itoa\n", n65, n65)
	var (
		word = "Go语言"
		rw   = []byte(word)
		rb   = []rune(word)
	)
	fmt.Printf("len(\"%v\") = %d字节，[]rune 长度 = %d字符\n", word, len(rw), len(rb))
	fmt.Println("[]byte 前两个元素 =", rw[:2], "转回 string = ", string(rw[:2]))
	fmt.Println()
	fmt.Println("===Go没有隐式转换===")
	var f float64 = float64(a)
	fmt.Println("float(a) = ", f, "|相加要同类型:", f+float64(b))

}
