// 语法点：map——哈希表的增删查改；零值 map 只能读不能写，遍历顺序永远不保证。
//
// 运行：go run ./syntax/04_map
package main

import "fmt"

func main() {
	fmt.Println("=== 增删查改 ===")
	// make 出来的 map 已经初始化，可以直接写入。
	scores := make(map[string]int)
	scores["alice"] = 90 // 增
	scores["bob"] = 85
	scores["carol"] = 78
	fmt.Println("scores =", scores, "len =", len(scores))

	scores["bob"] = 88 // 改（key 已存在就是覆盖）
	fmt.Println("bob 改成 88 后 =", scores["bob"])

	// 查：单值写法在 key 不存在时返回零值，没法区分「值是 0」和「key 不存在」。
	fmt.Println("scores[\"dave\"] =", scores["dave"], "（key 不存在，返回 int 零值 0）")

	// 二值写法 ok 才是判断存在性的正确方式。
	if v, ok := scores["dave"]; ok {
		fmt.Println("dave 存在，值 =", v)
	} else {
		fmt.Println("dave 不存在（用 ok 才能和「值恰好是 0」区分开）")
	}

	// 删：delete 删不存在的 key 是安全的空操作，不会报错也不会 panic。
	delete(scores, "carol")
	delete(scores, "not-exist") // 安全
	fmt.Println("删除 carol 后 =", scores, "len =", len(scores))

	// 清空整个 map 的 Go 1.21+ 标准做法（等价于重新 make，但会保留桶的容量）。
	clear(scores)
	fmt.Println("clear 之后 =", scores, "len =", len(scores))

	fmt.Println()
	fmt.Println("=== 用 map 计数 ===")
	words := []string{"go", "slice", "go", "map", "go", "slice"}
	counter := make(map[string]int)
	for _, w := range words {
		counter[w]++ // key 不存在时读到的零值 0，自增后写回 1，天然适合计数
	}
	fmt.Println("词频 =", counter, "len =", len(counter))
	fmt.Println("go 出现次数 =", counter["go"])

	fmt.Println()
	fmt.Println("=== 遍历顺序是随机的 ===")
	// Go 故意在每次 range 时随机化起始位置，防止程序依赖遍历顺序。
	// 把这一段多跑几次（go run ./syntax/04_map）会看到输出顺序变化。
	ordered := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}
	for round := 1; round <= 3; round++ {
		fmt.Printf("第 %d 轮：", round)
		for k, v := range ordered {
			fmt.Printf("%s=%d ", k, v)
		}
		fmt.Println()
	}
	fmt.Println("需要固定顺序时：先把 key 取出来放进切片，排序后再按顺序取。")

	fmt.Println()
	fmt.Println("=== 坑：对零值 map 写入会 panic ===")
	// 只声明不 make 的 map 是 nil map：读、len、range、delete 都安全，
	// 但写入会直接 panic。
	var nilMap map[string]int
	fmt.Println("nil map：len =", len(nilMap), "，读 key =", nilMap["x"],
		"，是否等于 nil：", nilMap == nil)
	for range nilMap { // 空 map 的 range 一次都不执行，是安全的
		fmt.Println("这行永远不会执行")
	}
	delete(nilMap, "x") // delete 对 nil map 也是安全的空操作
	fmt.Println("对 nil map 做读/len/range/delete 都安全。")
	// 为什么写入会 panic：map 内部指向一个 runtime 的 hmap 结构，
	// nil map 的指针是 nil，赋值时需要写哈希桶，于是触发
	// "assignment to entry in nil map" 运行时 panic。下面这行不能真的执行：
	//
	//	nilMap["x"] = 1 // panic: assignment to entry in nil map

	// 正确做法：用 make 初始化，或者用字面量。
	goodMap := make(map[string]int)
	goodMap["x"] = 1
	fmt.Println("make(map[string]int) 之后可以写入 =", goodMap)

	// 字面量写法等价于 make + 逐个赋值。
	literal := map[string]int{"x": 1}
	fmt.Println("字面量写法 =", literal)

	fmt.Println()
	fmt.Println("=== map 的值类型可以是任意类型 ===")
	// value 是切片时，注意不能直接对 map 的元素做 append 后原地生效。
	type point struct{ X, Y int }
	graph := map[string][]int{"a": {1, 2}} // 省略 value 类型字面量
	graph["a"] = append(graph["a"], 3)     // 必须整体写回，map 元素不可寻址
	graph["b"] = []int{9}
	centers := map[string]point{"origin": {0, 0}, "one": {1, 1}}
	fmt.Println("graph =", graph)
	fmt.Println("centers =", centers)
	fmt.Println("centers[\"one\"] =", centers["one"], "，X =", centers["one"].X)
	// 为什么 map 元素不能取地址：map 扩容时元素会在内存中搬家，
	// 允许取地址会导致悬空指针，所以语言层面直接禁止（编译错误：
	// cannot take address of ... / cannot assign to struct field in map）。
}
