// 语法点：数组 vs 切片——数组是值类型（赋值即整份拷贝），切片是共享底层数组的描述符。
//
// 运行：go run ./playground/03_slice
package main

import "fmt"

func main() {
	fmt.Println("===数组是值类型：赋值就是拷贝===")
	arr1 := [3]int{1, 2, 3}
	fmt.Println("arr1 =", arr1, "（没被改动）")
	arr2 := arr1
	arr2[0] = 999
	fmt.Println("arr2=", arr2, "(只改动了arr2)")
	modifyArray(arr1) // 数组传参也是拷贝：函数里把 a[0] 改成 999，arr1 不受影响
	fmt.Println("modifyArray(arr1)之后arr1 = ", arr1)
	fmt.Printf("arr1 长度 = %v,类型 = %T\n", len(arr1), arr1)
	fmt.Println()
	fmt.Println("====切片 ptr/len/cap===")
	s := []int{1, 2, 3}
	fmt.Printf("s = %v,len = %v,cap = %v,类型 = %T\n", s, len(s), cap(s), s)
	sr := make([]int, 3, 10)
	fmt.Printf("make([]int,3,10) = %v, len = %v,cap = %v\n", sr, len(sr), cap(sr))
	sub := s[1:]
	fmt.Printf("s[1:3] = %v,len = %v,cap = %v\n", sub, len(sub), cap(sub))
	var nilSlice []int
	fmt.Printf("nil切片 = %v,len = %v,cap = %v,是否等于nil: %v\n", nilSlice, len(nilSlice), cap(nilSlice), nilSlice == nil)
	fmt.Println()
	fmt.Println("===append与扩容===")
	grow := make([]int, 0, 2)
	for i := 1; i <= 6; i++ {
		calen, cacap := len(grow), cap(grow)
		grow = append(grow, i)
		fmt.Printf("append 第 %d 个元素 :len %v -> %v,cap %v ->%v\n", i, calen, len(grow), cacap, cap(grow))

	}

	fmt.Println()
	fmt.Println("===坑：切片共享底层数组===")
	s1 := []int{1, 2, 3}
	sub2 := s1[:2]
	sub2 = append(sub2, 9) // len=2 < cap=3，有余量，于是 append 在原地写入
	fmt.Println("s1 =", s1, "(第三个元素被覆盖了！)")
	fmt.Println("sub2 : ", sub2)
	// 结论：子切片不是副本；只要共享底层数组，任何一方的原地写入另一方都看得见。

	fmt.Println()
	fmt.Println("===修复:make + copy 做真拷贝===")
	fix1 := []int{1, 2, 3}
	fix2 := make([]int, 2)
	copy(fix2, fix1[:2])
	fix2 = append(fix2, 9)
	fmt.Println("fix1 = ", fix1, "(没被改动)")
	fmt.Println("fix2 =", fix2)
	fmt.Printf("fix1 len/cap = %d/%d,fix2 len/cap = %d/%d\n", len(fix1), cap(fix1), len(fix2), cap(fix2))
	guard := fix1[0:2:2]
	guard = append(guard, 7)
	fmt.Println("三索引切片 append 后 fix1 =", fix1, "(安全),guard =", guard)

	fmt.Println()
	fmt.Println("===copy的返回值===")
	dst := make([]int, 3)
	copied := copy(dst, []int{4, 5})
	fmt.Printf("copy(dst, []int{4, 5}) 返回 %d,dst = %v\n", copied, dst)
}

// modifyArray 把数组的第一个元素改成 999。
//
// 参数是 [3]int（数组）⇒ 调用时整份拷贝进函数，函数里改的是副本，
// 所以调用方的 arr1 不会变——这就是"数组是值类型"的证据。
func modifyArray(a [3]int) {
	a[0] = 999
}
