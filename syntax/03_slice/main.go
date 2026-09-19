// 语法点：数组 vs 切片——数组是值类型（赋值即整份拷贝），切片是共享底层数组的描述符。
//
// 运行：go run ./syntax/03_slice
package main

import "fmt"

func main() {
	fmt.Println("=== 数组是值类型：赋值就是拷贝 ===")
	// 数组的类型包含长度：[3]int 和 [4]int 是两个不同的类型。
	arr1 := [3]int{1, 2, 3}
	arr2 := arr1 // 整份拷贝，arr2 和 arr1 没有任何关系
	arr2[0] = 999
	fmt.Println("arr1 =", arr1, "（没被改动）")
	fmt.Println("arr2 =", arr2, "（只改了 arr2）")

	// 传参同理：数组作为参数也是拷贝一份，函数内改不动原数组。
	modifyArray(arr1)
	fmt.Println("modifyArray(arr1) 之后 arr1 =", arr1)

	// 长度是类型的一部分，可以用 len 拿到。
	fmt.Printf("arr1 长度 = %d，类型 = %T\n", len(arr1), arr1)

	fmt.Println()
	fmt.Println("=== 切片：ptr / len / cap ===")
	// 切片本身是个小结构体：一个指向底层数组的指针 + 长度 + 容量。
	s := []int{1, 2, 3}
	fmt.Printf("s = %v, len = %d, cap = %d, 类型 = %T\n", s, len(s), cap(s), s)

	// make 可以显式指定 len 和 cap：前 len 个元素是零值。
	s2 := make([]int, 3, 10)
	fmt.Printf("make([]int, 3, 10) = %v, len = %d, cap = %d\n", s2, len(s2), cap(s2))

	// 切片表达式：左闭右开，取出来的切片和原切片共用底层数组。
	sub := s[1:3]
	fmt.Printf("s[1:3] = %v, len = %d, cap = %d（cap 从起点算到原数组末尾）\n",
		sub, len(sub), cap(sub))

	// nil 切片：可以直接 len/cap/append/range，不需要初始化。
	var nilSlice []int
	fmt.Printf("nil 切片 = %v, len = %d, cap = %d, 是否等于 nil：%v\n",
		nilSlice, len(nilSlice), cap(nilSlice), nilSlice == nil)

	fmt.Println()
	fmt.Println("=== append 与扩容 ===")
	// append 返回新切片，必须用返回值覆盖原变量（append 可能换一块新数组）。
	grow := make([]int, 0, 2)
	for i := 1; i <= 6; i++ {
		beforeLen, beforeCap := len(grow), cap(grow)
		grow = append(grow, i)
		fmt.Printf("append 第 %d 个元素：len %d->%d, cap %d->%d\n",
			i, beforeLen, len(grow), beforeCap, cap(grow))
	}
	// 观察点：cap 不够时 append 会分配更大的数组，cap 突然变大（约 2 倍增长，
	// 元素变多后增长因子会变小），并且新数组的地址和旧数组不同。
	fmt.Println("扩容后的 grow =", grow)

	fmt.Println()
	fmt.Println("=== 坑：切片共享底层数组 ===")
	s1 := []int{1, 2, 3}
	// sub2 是 s1 的子切片，和 s1 指向同一块底层数组
	// （注意别和上面 make([]int, 3, 10) 的 s2 重名，同一个函数里 := 不能重复声明同名变量）
	sub2 := s1[:2]
	// len=2 < cap=3，有余量，于是 append 在原地写入
	sub2 = append(sub2, 9)
	fmt.Println("s1 =", s1, "（第 3 个元素被覆盖了！）")
	fmt.Println("sub2 =", sub2)
	// 为什么：s1 的底层数组是 [1 2 3]，cap 为 3；s1[:2] 得到的 sub2 与它共享这块
	// 数组，且 cap 仍有 3。append(sub2, 9) 发现 cap 够用，就不再分配新数组，
	// 直接把 9 写到下标 2，也就是 s1 的第 3 个元素的位置。
	// 结论：子切片不是副本；只要共享底层数组，任何一方的原地写入另一方都看得见。

	fmt.Println()
	fmt.Println("=== 修复：make + copy 做真拷贝 ===")
	fix1 := []int{1, 2, 3}
	// 新数组，len=2、cap=2
	fix2 := make([]int, 2)
	// 把前两个元素拷过去
	copy(fix2, fix1[:2])
	// cap 已满，这次 append 必然分配新数组
	fix2 = append(fix2, 9)
	fmt.Println("fix1 =", fix1, "（没被改动）")
	fmt.Println("fix2 =", fix2)
	fmt.Printf("fix1 len/cap = %d/%d, fix2 len/cap = %d/%d\n",
		len(fix1), cap(fix1), len(fix2), cap(fix2))

	// 另一个常用修复方式：三索引切片 s[low:high:max]，把 cap 限死，
	// 这样 append 一定会分配新数组，不会再写回原数组。
	// 下面这个三索引切片的 len=2、cap=2
	guard := fix1[0:2:2]
	guard = append(guard, 7)
	fmt.Println("三索引切片 append 后 fix1 =", fix1, "（安全），guard =", guard)

	fmt.Println()
	fmt.Println("=== copy 的返回值 ===")
	dst := make([]int, 3)
	// copy 只拷 min(len(dst), len(src)) 个元素，返回值是实际拷贝的个数
	copied := copy(dst, []int{4, 5})
	fmt.Printf("copy(dst, []int{4, 5}) 返回 %d，dst = %v\n", copied, dst)
}

// modifyArray 演示数组传参是拷贝：函数内改了，调用方看不到。
func modifyArray(a [3]int) {
	a[0] = -1
}
