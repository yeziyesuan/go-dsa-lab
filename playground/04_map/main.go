// 语法点：map——增删查改；零值 map 只能读不能写；遍历顺序永远不保证。
//
// 运行：go run ./playground/04_map
// 约定：本文件的代码由我凭记忆手敲（卡住才回去瞄 syntax/04_map/main.go）；
// 下面只给骨架和 TODO，答案自己写，随时保持能编译。
package main

import (
	"fmt"
	"sort"    // ⟵ 助手补：TODO 8 固定遍历顺序要用
	"strings" // ⟵ 助手补：TODO 11 countWords 按空格切词要用
)

func main() {
	fmt.Println("=== 增删查改 ===")
	// TODO 1：make 一个 map[string]int，写入 3 个人（alice=90 / bob=85 / carol=78），打印 map 和 len
	// TODO 2：把已有的 key 改一次（覆盖），打印新值
	// TODO 3：查一个不存在的 key —— 先用单值写法（看看返回什么），再用 `v, ok :=` 二值写法判断存在
	// TODO 4：delete 一个存在的 key + delete 一个不存在的 key（结论：删不存在的不会报错）
	// TODO 5：用 Go 1.21+ 的 clear 清空，打印 map 和 len
	scores := make(map[string]int)
	scores["alice"] = 90
	scores["bob"] = 85
	scores["carol"] = 78
	fmt.Printf("scores = %v, len = %v\n", scores, len(scores))
	scores["bob"] = 88
	fmt.Println("bob 改成 88 后 = ", scores["bob"])
	fmt.Println("scores[\"dave\"] = ", scores["dave"], "(key 不存在, 返回 int 零值 0)")
	if v, ok := scores["dave"]; ok {
		fmt.Println("dave 存在 值 = ", v)
	} else {
		fmt.Println("dave 不存在 (用 ok 才能和[恰好是0]区分开来)")
	}
	delete(scores, "carol")
	fmt.Printf("删除 carol 后 = %v len = %v\n", scores, len(scores))
	clear(scores)
	fmt.Printf("clear 之后 = %v len = %v\n", scores, len(scores))

	fmt.Println()
	fmt.Println("=== 用 map 计数 ===")
	// TODO 6：给 []string{"go", "slice", "go", "map", "go", "slice"} 数字典，
	//         打印词频，并单独打印 "go" 出现了几次
	words := []string{"go", "map", "go", "slice", "go", "slice"}
	scores1 := make(map[string]int)
	for _, v := range words {
		scores1[v]++
	}
	fmt.Printf("词频 = %v len = %v\n", scores1, len(scores1))
	fmt.Println("go 出现次数 = ", scores1["go"])

	fmt.Println()
	fmt.Println("=== 遍历顺序是随机的 ===")
	// TODO 7：range 一个有 4 个 key 的 map，打印 3 轮，观察顺序每次都不一样
	// TODO 8：需要固定顺序时——把 key 取进切片 → sort.Strings → 按序取值打印
	num := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}
	for i := 1; i <= 3; i++ {
		fmt.Printf("第 %d 轮：", i)
		for k, v := range num {
			fmt.Printf(" %s=%d", k, v) // ⟵ 助手补了前导空格：不然四组会连成一串看不出顺序
		}
		fmt.Println()
	}
	fmt.Println("需要固定顺序时 : 先把 key 取出来放进切片 , 排序后再按照顺序取。")

	// ⟵ 助手补：TODO 8 —— 固定顺序三步：取 key → 排序 → 按 key 取值
	keys := make([]string, 0, len(num))
	for k := range num {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	fmt.Print("排序后：")
	for _, k := range keys {
		fmt.Printf(" %s=%d", k, num[k])
	}
	fmt.Println()

	fmt.Println()
	fmt.Println("=== 坑 1：对零值（nil）map 写入会 panic ===")
	// TODO 9：var nilMap map[string]int
	//         ① 先证明读、len、range、delete 都是安全的；
	//         ② 再把 `nilMap["x"] = 1` 这行打开跑一次，把 panic 原文（assignment to entry in nil map）
	//            抄进 notes/ 的 3–5 行笔记里；
	//         ③ 最后用 make(map[string]int) 写出正确版本
	// ⟵ 助手修正：原来写的是 `var nilMap = make(map[string]int)`，
	//    make 出来的 map 已经初始化，`nilMap == nil` 会打印 false，等于没演示到坑。
	//    要复现必须「只声明、不 make」：
	var nilMap map[string]int
	fmt.Printf("nil map: len = %d , 读 key = %d , 是否等于 nil: %t\n", len(nilMap), nilMap["x"], nilMap == nil)
	for range nilMap {
		fmt.Println("这行会不会输出呢？")
	}
	delete(nilMap, "x")
	fmt.Println("对 nil map 做读/len/range/delete 都安全。")
	// nilMap["x"] = 1 // ⟵ 取消注释单独跑一次，就会看到：panic: assignment to entry in nil map
	goodmap := make(map[string]int)
	goodmap["x"] = 1
	fmt.Println("make(map[string]int) 之后可以写入 = ", goodmap)
	literal := map[string]int{"x": 1}
	fmt.Println("字面量写法 = ", literal)

	fmt.Println()
	fmt.Println("=== 坑 2：map 元素不可寻址 ===")
	// TODO 10：value 是切片时，`m["a"] = append(m["a"], x)` 必须整体写回；
	//          value 是 struct 时不能直接 `m["k"].X = 1`——把编译器报的错原文抄下来

	// ⟵ 助手补：① value 是切片 —— append 的结果必须整体写回
	graph := map[string][]int{"a": {1, 2}}
	// graph["a"] = append(graph["a"], 3) 拆开看：append 返回的是新切片头，
	// 不写回 map 的话，扩容后的新底层数组就丢了（原地的改动也可能只改到旧数组）。
	graph["a"] = append(graph["a"], 3)
	graph["b"] = []int{9}
	fmt.Println("graph =", graph)

	// ⟵ 助手补：② value 是 struct —— 不能直接改字段（map 元素不可寻址）
	type point struct{ X, Y int }
	centers := map[string]point{"origin": {0, 0}}
	// centers["origin"].X = 9 // 取消注释就是编译错误：
	//   cannot assign to struct field centers["origin"].X in map
	// 原因：map 扩容时元素会搬家，允许取地址会出现悬空指针，语言层面直接禁止。
	p := centers["origin"] // 正确姿势：取出来 → 改副本 → 整体写回
	p.X = 9
	centers["origin"] = p
	fmt.Println("centers =", centers)

	fmt.Println()
	fmt.Println("=== 收口 ===")
	// TODO 11（有余力再做）：写 countWords(s string) map[string]int，用它统计一句话的词频
	// ⟵ 助手补：调用
	fmt.Println(`countWords("go is fast and go is fun") =`, countWords("go is fast and go is fun"))
}

// TODO 11 的函数写在这里：
// ⟵ 助手补：countWords 按空白切词并计数；返回的 map 是 make 出来的，所以调用方可以直接继续写
func countWords(s string) map[string]int {
	counts := make(map[string]int)
	for _, w := range strings.Fields(s) { // Fields 会按任意连续空白切分，且自动丢掉空串
		counts[w]++
	}
	return counts
}
