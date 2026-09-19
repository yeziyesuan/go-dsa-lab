// 语法点：并发入门——goroutine + WaitGroup、channel 收发、select（本周只求看懂，W2 才深入）。
//
// 运行：go run ./syntax/11_concurrency
package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("=== goroutine + sync.WaitGroup ===")
	// go 关键字启动一个 goroutine（由运行时调度的轻量线程）。
	// 主 goroutine 结束 => 整个程序结束，所以必须等子 goroutine 干完活。
	// WaitGroup 就是干这个的：Add 登记、Done 减一、Wait 阻塞到归零。
	var wg sync.WaitGroup
	for i := 1; i <= 3; i++ {
		wg.Add(1) // 必须在启动 goroutine 之前 Add，避免 Wait 提前返回
		// Go 1.22 起循环变量每轮独立，所以这里能安全地闭包捕获 i。
		go func(n int) {
			defer wg.Done() // 用 defer 保证即使中途 panic/return 也会减一
			fmt.Printf("goroutine %d 正在跑\n", n)
		}(i)
	}
	wg.Wait() // 阻塞，直到 3 个 goroutine 都调用了 Done
	fmt.Println("所有 goroutine 结束（wg.Wait() 返回了）")

	fmt.Println()
	fmt.Println("=== 无缓冲 channel：发送和接收必须配对 ===")
	// 无缓冲 channel 的发送会阻塞，直到有另一个 goroutine 来接收；
	// 所以在同一个 goroutine 里先发后收会死锁（fatal error: all goroutines are asleep）。
	ch := make(chan string) // 无缓冲
	go func() {
		ch <- "来自 goroutine 的消息" // 阻塞在这里，直到 main 接收
	}()
	msg := <-ch // 接收，配对的发送同时解除阻塞
	fmt.Println("收到：", msg)
	// 下面这种写法会死锁，不能真的执行：
	//	ch <- "x" // 当前没有别的 goroutine 在接收，直接 deadlock

	// 用完就关：close 表示「不会再有值发进来了」，
	// 接收方可以用 v, ok := <-ch 判断通道是否已关闭并读空。
	close(ch)
	v, ok := <-ch
	fmt.Printf("通道已关闭：读到 %q, ok = %v（读空后得到零值，ok 为 false）\n", v, ok)

	fmt.Println()
	fmt.Println("=== 带缓冲 channel：缓冲区没满就不阻塞 ===")
	buf := make(chan int, 3) // 容量 3
	buf <- 1                 // 不阻塞：进缓冲区
	buf <- 2
	fmt.Printf("往容量 3 的缓冲通道发了 2 个值，len = %d, cap = %d\n", len(buf), cap(buf))
	buf <- 3 // 缓冲区刚好满
	fmt.Printf("发满 3 个后 len = %d\n", len(buf))
	// 第 4 个会阻塞（缓冲区满），所以下面的循环要先把值取出来再发。
	//	buf <- 4 // 阻塞：缓冲区已满，需要有人接收
	close(buf) // 关闭后仍可把缓冲区里剩余的值读完
	for x := range buf {
		fmt.Print(x, " ")
	}
	fmt.Println("<- range 一个已关闭的缓冲通道，把剩下的值读完就退出循环")
	fmt.Printf("读空后 len = %d\n", len(buf))

	fmt.Println()
	fmt.Println("=== select：多路复用 ===")
	c1 := make(chan string, 1)
	c2 := make(chan string, 1)
	c1 <- "来自 c1"
	// c2 故意不发：select 会走 default 分支，而不是一直卡住。
	select {
	case m1 := <-c1:
		fmt.Println("select 命中 c1：", m1)
	case m2 := <-c2:
		fmt.Println("select 命中 c2：", m2)
	default:
		// default 让 select 变成非阻塞：没有就绪的 case 就立刻执行它。
		fmt.Println("select 走 default：当前没有通道就绪")
	}

	// 再发一个到 c2，这次 select 就能命中 c2 了（多个就绪时随机选一个）。
	c2 <- "来自 c2"
	select {
	case m1 := <-c1:
		fmt.Println("这一轮命中 c1：", m1)
	case m2 := <-c2:
		fmt.Println("这一轮命中 c2：", m2)
	default:
		fmt.Println("这一轮没有通道就绪")
	}
	// 去掉 default 的 select 会一直阻塞，直到某个 case 就绪——
	// 这是「等待多个事件中任意一个」的标准写法。
	//	select {
	//	case m := <-c1: // 阻塞直到 c1 或 c2 有值
	//	case m := <-c2:
	//	}
	//
	// 常用于超时的写法（W2 会细讲，这里只注释出来）：
	//	select {
	//	case m := <-c1:
	//		fmt.Println(m)
	//	case <-time.After(time.Second):
	//		fmt.Println("超时")
	//	}

	fmt.Println()
	fmt.Println("=== 多生产者汇总到同一个 channel ===")
	sum := make(chan int, 3)
	var wg2 sync.WaitGroup
	for i := 1; i <= 3; i++ {
		wg2.Add(1)
		go func(n int) {
			defer wg2.Done()
			sum <- n * n // 缓冲够大，不会阻塞
		}(i)
	}
	wg2.Wait()
	close(sum) // 所有发送者都结束后才能关，否则可能 panic: send on closed channel
	total := 0
	for x := range sum {
		total += x
	}
	fmt.Println("1+4+9 的和 =", total)

	fmt.Println()
	fmt.Println("=== 本周只需看懂，别踩的坑 ===")
	// 1) 主 goroutine 退出会直接终止程序，子 goroutine 不会「优雅收尾」——用 WaitGroup 等。
	// 2) 无缓冲 channel 同 goroutine 内自发自收 = 死锁（deadlock）。
	// 3) 往已关闭的 channel 发送会 panic；从已关闭的 channel 接收会立刻返回零值。
	// 4) 多个 goroutine 同时读写同一个变量（比如 map）会 data race，
	//    用 go run -race ./syntax/11_concurrency 可以检测，修法用 channel 或 sync.Mutex。
	//    共享内存 + 锁 vs channel 通信的取舍，W2 专门讲。
	// 5) GMP 调度器、channel 底层实现这些属于 W43 之后的内容，本周不碰。
	fmt.Println("W2 才深入并发，本周只求看懂语法形状。")
}
