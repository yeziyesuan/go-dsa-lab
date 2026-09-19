# syntax/ —— W1 语法示例

一个语法点一个目录，每个目录都是 `package main`，直接 `go run` 就能看到打印结果。
不引入任何第三方依赖，只用标准库（唯一的例外是 `10_package` 会 import 本仓库自己的 `pkgexample` 包，见文末说明）。

## 12 个示例

| 目录 | 语法点 | 运行命令 | 一句话要点 |
|---|---|---|---|
| `01_types` | 变量、常量、基本类型、类型转换 | `go run ./syntax/01_types` | 零值是 Go 的默认初值；**没有隐式转换**，`float64(3)/2` 这类转换必须手写 |
| `02_control` | if / for / switch | `go run ./syntax/02_control` | 没有 `while`，用 `for 条件` 代替；switch 默认不穿透，`fallthrough` 才穿透 |
| `03_slice` | 数组 vs 切片 | `go run ./syntax/03_slice` | 数组赋值即拷贝；切片是 `{ptr,len,cap}` 描述符，**子切片共享底层数组**，`append` 可能覆盖原数据 |
| `04_map` | map | `go run ./syntax/04_map` | 取存在的 key 要用 `v, ok :=`；**nil map 能读不能写**（写会 panic）；遍历顺序随机 |
| `05_func` | 函数特性 | `go run ./syntax/05_func` | 多返回值 + `error` 是 Go 的惯例；`defer` 是 LIFO，且能改命名返回值 |
| `06_pointer` | 指针与值传递 | `go run ./syntax/06_pointer` | Go **只有值传递**；要改到原值就传指针；切片/map 传的是描述符（指针） |
| `07_struct_method` | 结构体与方法 | `go run ./syntax/07_struct_method` | 值接收者改不动原对象，指针接收者才行；组合用**嵌入**，Go 没有继承 |
| `08_interface` | 接口 | `go run ./syntax/08_interface` | 接口是**隐式实现**的；`T` 和 `*T` 方法集不同，指针接收者让值类型不满足接口 |
| `09_error` | 错误处理 | `go run ./syntax/09_error` | 错误是返回值；**包装过就用 `errors.Is`/`errors.As`，不要用 `==`** |
| `10_package` | 包与模块 | `go run ./syntax/10_package` | 首字母大写才导出；`go.mod` 定导入路径前缀，`go run` = 编译 + 运行 |
| `11_concurrency` | 并发入门 | `go run ./syntax/11_concurrency` | `go` + `WaitGroup` 等结果，channel 传数据，`select` 多路等待（W2 才深入） |
| `12_stdlib` | 标准库速览 | `go run ./syntax/12_stdlib` | `fmt` / `strings` / `strconv` / `sort` / `os` / `bufio` 六个包的高频入口 |

## 建议学习顺序

**先建立「类型 → 控制流 → 数据结构 → 抽象」的主线，再补工程与工具：**

1. **`01_types`** —— 一切的基础：类型、零值、显式转换。看不懂零值，后面全是坑。
2. **`02_control`** —— 只有 if/for/switch 三种控制结构，半小时就能过完，但 `for range` 会贯穿整个仓库。
3. **`03_slice`** —— **本周最重要的一节**。切片共享底层数组是 LeetCode 里最隐蔽的 bug 来源，务必亲手改一改 `s2 := s1[:2]` 那段。
4. **`04_map`** —— 哈希表，配合 W1 的「两数之和」「字母异位词」直接就用上了。
5. **`05_func`** —— 多返回值 + `error`、闭包、`defer`。`defer` 的 LIFO 和「改命名返回值」在写题解时天天见。
6. **`06_pointer`** —— 理解「Go 只有值传递」这一句话，就能解释为什么改不动原值、为什么 map 不用指针也能改。
7. **`07_struct_method`** —— 结构体、方法、值/指针接收者。这是 08 的前置知识。
8. **`08_interface`** —— 接口与方法集。**必须放在 07 之后**，否则「方法集」讲不清。
9. **`09_error`** —— 哨兵错误、`%w` 包装、`errors.Is` / `errors.As`，以及 panic/recover 的边界。
10. **`10_package`** —— 包与模块，理解导入路径从哪来、首字母大小写怎么控制可见性。
11. **`11_concurrency`** —— 只求看懂形状：goroutine、channel、`select`。W2 再深入，别在这周纠结底层。
12. **`12_stdlib`** —— 当手册用：写题解时哪个函数想不起来就跑一遍。**`sort.Slice` 和 `strconv.Atoi` 建议背下来。**

一句话版：**1 → 2 → 3 → 4 → 5 → 6 → 7 → 8 → 9 → 10 → 11 → 12**，
其中 3、4、8 三节值得反复看，11 和 12 可以先跑一遍不求甚解。

## 自查命令

```bash
go vet ./syntax/...          # 静态检查，必须零报错
go run ./syntax/03_slice     # 跑单个示例
gofmt -l ./syntax            # 列出格式不正确的文件（应为空）
```

## 依赖说明

- `10_package` 依赖仓库根目录下的 `pkgexample` 包（`pkgexample/pkgexample.go`），
  调用的是 `pkgexample.Exported()` 和 `pkgexample.Greet("W1")`。
  **如果那个包不存在或被移动**，`10_package` 会报
  `no required module provides package github.com/yeziyesuan/go-dsa-lab/pkgexample`；
  其余 11 个示例不受影响，都是纯标准库、零外部依赖。

## 输出不唯一的示例（正常现象）

下面几处每次运行的输出会变，不是 bug：

| 目录 | 现象 | 原因 |
|---|---|---|
| `04_map` | 三轮遍历顺序不同 | Go 故意随机化 map 遍历起点 |
| `11_concurrency` | goroutine 打印顺序不同 | goroutine 调度顺序不保证 |
| `12_stdlib` | `os.Getenv("HOME")` 在 Windows 上为空 | Windows 用 `USERPROFILE`，不是 `HOME`；`os.Args` 长度也随参数变化 |

## 第一次跑之前请先验证

本目录的代码是按约定手写并逐文件复查过的，但**请在目标机器上先跑一遍再开始写题解**：

```bash
go vet ./syntax/...        # 先过静态检查（含上面提到的 10_package 依赖）
for d in 01_types 02_control 03_slice 04_map 05_func 06_pointer \
         07_struct_method 08_interface 09_error 10_package 11_concurrency 12_stdlib; do
  go run ./syntax/$d || echo "FAILED: $d"
done
gofmt -l ./syntax          # 期望输出为空
```

