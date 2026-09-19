// 语法点：标准库速览——fmt / strings / strconv / sort / os / bufio 的常用入口。
//
// 运行：go run ./syntax/12_stdlib
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// Person 用于演示 sort.Slice / sort.SliceStable。
type Person struct {
	Name string
	Age  int
}

func main() {
	fmt.Println("=== fmt：Sprintf 与 Printf 动词 ===")
	name := "Go"
	pi := 3.14159
	// Printf 直接输出，Sprintf 返回字符串（不会打印）。
	line := fmt.Sprintf("%s 的圆周率约等于 %.2f", name, pi)
	fmt.Println("Sprintf 得到：", line)

	// 常用动词速查。
	fmt.Printf("%%v  = %v   （默认格式，万能）\n", 42)
	fmt.Printf("%%d  = %d   （十进制整数）\n", 42)
	fmt.Printf("%%5d = %5d （右对齐，宽度 5）\n", 42)
	fmt.Printf("%%-5d= %-5d（左对齐，宽度 5）|\n", 42)
	fmt.Printf("%%05d= %05d （补零到 5 位）\n", 42)
	fmt.Printf("%%f  = %f   （浮点，默认 6 位小数）\n", pi)
	fmt.Printf("%%.3f= %.3f （保留 3 位小数）\n", pi)
	fmt.Printf("%%8.3f=%8.3f（宽度 8，保留 3 位）\n", pi)
	fmt.Printf("%%q  = %q   （带引号的字符串）\n", name)
	fmt.Printf("%%s  = %s   （字符串）\n", name)
	fmt.Printf("%%t  = %t   （布尔）\n", true)
	fmt.Printf("%%c  = %c   （字符）\n", 65)
	fmt.Printf("%%b  = %b   （二进制）\n", 5)
	fmt.Printf("%%x  = %x   （十六进制）\n", 255)
	fmt.Printf("%%T  = %T   （类型名）\n", pi)
	fmt.Printf("%%p  = %p   （指针地址）\n", &pi)
	fmt.Printf("%%+v = %+v （结构体带字段名）\n", Person{Name: "ziye", Age: 18})
	fmt.Printf("%%#v = %#v （Go 语法形式）\n", Person{Name: "ziye", Age: 18})
	fmt.Printf("%%v   = %v   （%% 本身要写成 %%%%）\n", 100)

	// Sprintf 拼格式化字符串，Fprintln 输出到指定 Writer（这里用标准输出）。
	fmt.Fprintln(os.Stdout, "Fprintln 写给 os.Stdout：", line)
	// Sprint 系列不格式化，直接在操作数之间加空格。
	fmt.Println("Sprintln =", fmt.Sprintln("a", 1, true))

	fmt.Println()
	fmt.Println("=== strings：切分、连接、判断、清理 ===")
	s := "  Go is a statically typed language  "
	fmt.Printf("原始字符串 = %q\n", s)

	trimmed := strings.TrimSpace(s) // 去掉首尾空白（含 \t \n）
	fmt.Printf("TrimSpace = %q\n", trimmed)

	parts := strings.Split(trimmed, " ") // 按分隔符切成切片
	fmt.Printf("Split 得到 %d 段：%v\n", len(parts), parts)

	// Fields 按「连续空白」切分，自动忽略多余空格，比 Split(s, " ") 更稳。
	fmt.Printf("Fields = %v\n", strings.Fields(trimmed))

	joined := strings.Join([]string{"a", "b", "c"}, "-") // 切片拼成字符串
	fmt.Println("Join =", joined)

	fmt.Println("Contains(language) =", strings.Contains(trimmed, "language"))
	fmt.Println("HasPrefix(Go)      =", strings.HasPrefix(trimmed, "Go"))
	fmt.Println("HasSuffix(age)     =", strings.HasSuffix(trimmed, "age"))
	fmt.Println("Index(typed)       =", strings.Index(trimmed, "typed"))
	fmt.Println("ToUpper            =", strings.ToUpper("go"))
	fmt.Println("ToLower            =", strings.ToLower("GO"))
	fmt.Println("ReplaceAll         =", strings.ReplaceAll("a-b-c", "-", "+"))
	fmt.Println("Count(a)           =", strings.Count("banana", "a"))
	// 大小写不敏感的比较，用 EqualFold。
	fmt.Println("EqualFold(GO, go)  =", strings.EqualFold("GO", "go"))

	// 用 Builder 高效拼字符串：避免 += 产生一堆中间字符串。
	// strings.Builder 是 W1 唯一建议记的「性能习惯」，原理（底层 []byte）W43 之后再看。
	var sb strings.Builder
	words := []string{"W1", "语法", "示例"}
	for i, w := range words {
		if i > 0 {
			sb.WriteString(" / ")
		}
		sb.WriteString(w)
	}
	fmt.Println("Builder 拼接结果 =", sb.String(), "，长度 =", sb.Len())

	fmt.Println()
	fmt.Println("=== strconv：字符串与数字互转 ===")
	// Atoi = ASCII to integer，Itoa 是反过来；失败时 Atoi 返回 error。
	n, err := strconv.Atoi("123")
	if err != nil {
		fmt.Println("Atoi 出错：", err)
	} else {
		fmt.Println("Atoi(\"123\") =", n, "，n+1 =", n+1)
	}

	// 关键：一定要检查 error，否则会拿到零值还以为转换成功了。
	bad, err := strconv.Atoi("12x")
	if err != nil {
		// Atoi 的错误是 *strconv.NumError，里面带着原始输入。
		fmt.Printf("Atoi(\"12x\") 失败：%v，返回值是零值 %d\n", err, bad)
	}

	fmt.Println("Itoa(456) =", strconv.Itoa(456))

	// ParseInt 可以指定进制和位宽，比 Atoi 更灵活。
	v64, err := strconv.ParseInt("-1000000", 10, 64)
	if err != nil {
		fmt.Println("ParseInt 出错：", err)
	} else {
		fmt.Println("ParseInt(\"-1000000\", 10, 64) =", v64)
	}
	// 超出位宽会返回 error，而不是悄悄截断。
	overflow, err := strconv.ParseInt("99999999999999999999", 10, 64)
	if err != nil {
		fmt.Printf("ParseInt 溢出：%v（返回值 %d 不可用）\n", err, overflow)
	}
	// 其他常用：ParseFloat、ParseBool、FormatInt、Quote。
	f, err := strconv.ParseFloat("2.5", 64)
	if err != nil {
		fmt.Println("ParseFloat 出错：", err)
	} else {
		fmt.Println("ParseFloat(\"2.5\", 64) =", f)
	}
	b, err := strconv.ParseBool("true")
	if err != nil {
		fmt.Println("ParseBool 出错：", err)
	} else {
		fmt.Println("ParseBool(\"true\") =", b)
	}

	fmt.Println()
	fmt.Println("=== sort：排序三位一体 ===")
	nums := []int{5, 2, 9, 1, 5}
	sort.Ints(nums) // 原地升序
	fmt.Println("sort.Ints 之后 =", nums)

	strs := []string{"banana", "apple", "cherry"}
	sort.Strings(strs)
	fmt.Println("sort.Strings 之后 =", strs)

	// 经典错误写法：sort.Ints 是原地排序，没有返回值。
	//	sorted := sort.Ints(nums) // 编译失败：sort.Ints (value of type func...) used as value

	people := []Person{
		{"carol", 30},
		{"alice", 25},
		{"bob", 25},
		{"dave", 20},
	}
	// sort.Slice：用自定义 Less 排序（不稳定）。
	sort.Slice(people, func(i, j int) bool {
		return people[i].Age < people[j].Age
	})
	fmt.Println("按年龄升序（Slice，不稳定）：", people)

	// sort.SliceStable：相等元素保持原有相对顺序。
	// 下面按年龄升序排，同龄的 carol(原下标 0) 会排在 bob(原下标 2) 前面，
	// 这正是原始顺序，稳定排序保证这一点。
	sort.SliceStable(people, func(i, j int) bool {
		return people[i].Age < people[j].Age
	})
	fmt.Println("按年龄升序（SliceStable）：", people)

	// 多关键字排序：先按年龄，年龄相同再按姓名。
	sort.Slice(people, func(i, j int) bool {
		if people[i].Age != people[j].Age {
			return people[i].Age < people[j].Age
		}
		return people[i].Name < people[j].Name
	})
	fmt.Println("先年龄后姓名：", people)

	// 降序：把 Less 的比较反过来即可。
	sort.Slice(people, func(i, j int) bool {
		return people[i].Age > people[j].Age
	})
	fmt.Println("按年龄降序：", people)

	fmt.Println()
	fmt.Println("=== os：命令行参数与环境变量 ===")
	// os.Args[0] 是程序路径，真正的参数从下标 1 开始。
	fmt.Printf("os.Args = %v，参数个数（不含程序名）= %d\n", os.Args, len(os.Args)-1)
	if len(os.Args) > 1 {
		fmt.Println("第一个参数 =", os.Args[1])
	} else {
		fmt.Println("没有额外参数。试一下：go run ./syntax/12_stdlib hello world")
	}

	// Getenv 没设置时返回空串，不会报错；想区分「未设置」用 LookupEnv。
	home := os.Getenv("HOME") // Windows 上通常为空，Linux/macOS 有值
	fmt.Printf("os.Getenv(\"HOME\") = %q\n", home)
	if v, ok := os.LookupEnv("PATH"); ok {
		fmt.Println("PATH 已设置，长度 =", len(v), "（LookupEnv 的 ok 表示环境变量是否存在）")
	}
	// 设置环境变量（只影响当前进程及子进程，不改系统设置）。
	if err := os.Setenv("DSA_LAB_DEMO", "W1"); err != nil {
		fmt.Println("Setenv 出错：", err)
	} else {
		fmt.Println("Setenv 之后 Getenv =", os.Getenv("DSA_LAB_DEMO"))
	}
	// defer os.Unsetenv("DSA_LAB_DEMO") 这样用完就清理，避免污染后续逻辑。
	defer os.Unsetenv("DSA_LAB_DEMO")

	// 其他常用：os.ReadFile / os.WriteFile / os.Exit(1) / os.Stdin / os.Stdout。
	// os.Exit 会跳过所有 defer，所以要放在最后或显式处理清理。

	fmt.Println()
	fmt.Println("=== bufio：逐行读标准输入 ===")
	// 逐行读 os.Stdin 的固定写法（下面这段是注释，真的读会一直等你敲键盘）：
	//
	//	scanner := bufio.NewScanner(os.Stdin)
	//	for scanner.Scan() {              // 每读一行返回 true，读到 EOF 或出错返回 false
	//		line := scanner.Text()        // 当前行内容（不含换行符）
	//		fmt.Println("读到：", line)
	//	}
	//	if err := scanner.Err(); err != nil { // 一定要查 Err()，区分 EOF 和真错误
	//		fmt.Fprintln(os.Stderr, "读输入出错：", err)
	//	}
	//
	// 要点：
	//   1) Scanner 默认按行切，单行上限 64KB（bufio.MaxScanTokenSize），超长行要用 Buffer 调大。
	//   2) 想同时读文件和标准输入，代码完全一样，只要把 os.Stdin 换成打开的文件。
	//   3) 非交互输入可以用管道喂：echo hello | go run ./syntax/12_stdlib
	fmt.Println("上面给出了 bufio.NewScanner 逐行读的写法（注释形式，不阻塞）。")

	// 用 strings.Reader 造一个假的输入源，这样既能真跑 Scanner，又不会卡住。
	br := bufio.NewReader(strings.NewReader("第一行\n第二行\n"))
	for {
		// ReadString 读到分隔符为止，返回内容包含分隔符。
		lineText, err := br.ReadString('\n')
		if lineText != "" {
			fmt.Printf("模拟读入：%q\n", strings.TrimRight(lineText, "\n"))
		}
		if err != nil {
			// io.EOF 表示读完了；这里不 import io，只判断非 nil 就退出。
			fmt.Println("读到结尾（err 非 nil 说明后面没有数据了）")
			break
		}
	}
}
