# map 为什么并发不安全：这是 crash，不是 panic

- 类型：语法 · Go 基础
- 一句话定位：map 是纯单线程容器，**并发读写会让整个进程直接死掉**，
  这一点和 Java 里下意识信任 `ConcurrentHashMap` 的习惯正好相反。
- 配套示例：`syntax/04_map`（基础操作）、`syntax/11_concurrency`（并发形状）

## 1. 零值不可写：nil map 只能读不能写

```go
var m map[string]int
fmt.Println(m["a"]) // 0，读 nil map 是合法的
m["a"] = 1          // panic: assignment to entry in nil map
```

声明只给了个 nil 描述符，没有任何哈希桶。所以 map 要么 `make`，
要么用字面量 `map[string]int{"a": 1}` 初始化。读不 panic 是因为设计上
"取不存在的 key 返回零值"，写没有零值可返回，只能 panic。

## 2. 三个日常操作

```go
m := make(map[string]int)
m["a"] = 1
v, ok := m["a"]   // 二值取法：ok 区分"值是 0"和"key 不存在"
_ = v
delete(m, "a")    // 删不存在的 key 是合法且无操作，不 panic
```

`v, ok := m[k]` 是这个语言里最常用的惯用法之一，因为 Go 没有 `containsKey`。
另一个必知的点：**遍历顺序是随机的**，运行时故意在每轮 `range` 时打乱起始位置，
所以你没法依赖 map 的遍历顺序，需要有序就自己取 key 排序。`len(m)` 是 O(1)，
但 `cap(m)` 不合法（cap 只对数组、切片、通道有效），写了会直接编译不过。

## 3. 并发读写 = runtime 直接 crash

Go 运行时会在 map 的写路径上打标记，一旦发现有另一个 goroutine 正在操作同一个 map，
就调用 `throw` 打出 `fatal error: concurrent map writes`。

关键区别：**`throw` 不经过 panic 机制**。普通 panic 是可以用 `recover` 拦下来的，
而 `fatal error` 会直接终止进程，`defer` 里的 `recover` 也救不回来。
死掉的 goroutine 不会报错给调用方，它的错误就是"整个程序没了"。
更阴的是这种检查是"能发现就发现"：有时能检测到，有时只是数据竞争后 map 内部
结构被写坏，表现为读到错值或诡异的 panic。所以不能靠"没崩"证明代码安全，
要在测试期用 race 检测器。

## 4. 复现代码 + 用 -race 观察

新建一个空目录（或在 `syntax/` 下自建目录），放 `main.go`：

```go
package main

import (
	"fmt"
	"sync"
)

func main() {
	m := make(map[int]int)
	var wg sync.WaitGroup
	for g := 0; g < 4; g++ {
		wg.Add(1)
		go func(g int) { // 加锁前，这段是错的写法，故意保留
			defer wg.Done()
			for i := 0; i < 2000; i++ {
				m[g*10000+i] = i
			}
		}(g)
	}
	wg.Wait()
	fmt.Println("size:", len(m))
}
```

- `go run -race .`：稳定报 `WARNING: DATA RACE`，指出两个 goroutine 读写同一地址，
  这是**排查阶段的正确姿势**（不用等它崩）。
- `go run .`：很可能直接 `fatal error: concurrent map writes` 或 `concurrent map read and map write`。
  注意程序不一定每次都崩，取决于调度，别把一次跑通当成没问题。

## 5. 三条解法

```go
// 解法一：sync.Mutex，最通用
var mu sync.Mutex
mu.Lock()
m[k] = v
mu.Unlock()

// 解法二：sync.RWMutex，读多写少时读之间不互斥
var rw sync.RWMutex
rw.RLock()
v := m[k]
rw.RUnlock()
```

**解法三 `sync.Map`**：`Store` / `Load` / `Delete` / `Range`，内部有只读快照与
dirty map 的分层，适合**读多写少、key 集合基本固定**的场景（缓存、注册表、
一次写入多次读取）。它不是"并发版 map"的通用替代：`Load` 返回 `any` 要类型断言、
`Range` 也是无序遍历、写多时性能反而不如 `Mutex` 包一层。
**默认选 `Mutex`/`RWMutex`，确认真的是读多写少且 key 稳定再换 `sync.Map`。**

## 一句话结论

**map 不是并发安全的容器**——nil map 不能写，并发读写是 runtime 直接 `throw` 的
`fatal error`（`recover` 无效），要并发就用 `sync.Mutex` / `sync.RWMutex`，
读多写少且 key 稳定才考虑 `sync.Map`；这一点和 Java 里 `ConcurrentHashMap` 的直觉相反。

## 我踩过的坑

- 从 Java 过来，以为 map 像 `ConcurrentHashMap` 一样自带并发保护，
  在 http handler 里直接写全局 map，压测时整个服务进程消失，日志里只有一行 fatal error。
- 先写了 `defer recover()` 想兜住，结果毫无作用——`fatal error` 不是 panic，
  只好改成入口处统一加锁。
- `if m["n"] == 0` 判断 key 是否存在，遇到"值真的是 0"的用例全错，改用二值取法。
- 用 `range` 遍历 map 直接生成接口返回，以为顺序稳定，前端对比 diff 时才发现每轮都不一样。
