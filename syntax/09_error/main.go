// 语法点：错误处理——error 是普通返回值；哨兵错误用 errors.Is，自定义错误用 errors.As。
//
// 运行：go run ./syntax/09_error
package main

import (
	"errors"
	"fmt"
)

// ErrNotFound 是哨兵错误（sentinel error）：包级变量，调用方用 errors.Is 判断。
// 命名惯例是 Err 前缀 + 大驼峰。
var ErrNotFound = errors.New("not found")

// ValidationError 是自定义错误类型：实现 error 接口只需要一个 Error() string 方法。
type ValidationError struct {
	Field string
	Msg   string
}

// Error 实现 error 接口。指针接收者 => *ValidationError 才是 error。
func (e *ValidationError) Error() string {
	return fmt.Sprintf("字段 %q 校验失败：%s", e.Field, e.Msg)
}

func main() {
	fmt.Println("=== error 就是一个普通返回值 ===")
	// error 是内置接口：type error interface { Error() string }
	// 惯例：error 永远是最后一个返回值，成功时返回 nil。
	if _, err := findUser("bob"); err != nil {
		fmt.Println("findUser(\"bob\") 出错：", err)
	}
	if u, err := findUser("alice"); err == nil {
		fmt.Println("findUser(\"alice\") 成功：", u)
	}

	fmt.Println()
	fmt.Println("=== 哨兵错误 + errors.Is ===")
	_, err := findUser("bob")
	// 判断哨兵错误要用 errors.Is，不要用 ==。
	fmt.Println("errors.Is(err, ErrNotFound) =", errors.Is(err, ErrNotFound))
	fmt.Println("err == ErrNotFound          =", err == ErrNotFound)

	_, err2 := loadConfig("prod")
	fmt.Println("loadConfig(\"prod\") 的错误：", err2)
	// 为什么这里 err == ErrNotFound 会是 false：
	// loadConfig 用 fmt.Errorf("...: %w", ErrNotFound) 包了一层，
	// err2 的最外层是 *fmt.wrapError，和 ErrNotFound 不是一个对象。
	fmt.Println("err2 == ErrNotFound          =", err2 == ErrNotFound, "（被 %w 包过，== 失效）")
	fmt.Println("errors.Is(err2, ErrNotFound) =", errors.Is(err2, ErrNotFound), "（errors.Is 会沿链解包）")
	// 为什么：errors.Is 会一层层 Unwrap 下去，拿每一层和 ErrNotFound 比，
	// 而 == 只比较最外层那个错误值本身。所以包过一层之后必须用 errors.Is。

	fmt.Println()
	fmt.Println("=== 自定义错误类型 + errors.As ===")
	_, err3 := validate("age", -1)
	if err3 != nil {
		fmt.Println("validate 返回：", err3)
	}
	// errors.As 把错误链里第一个能转成目标类型的错误提取出来。
	var ve *ValidationError
	if errors.As(err3, &ve) {
		fmt.Println("errors.As 提取成功：Field =", ve.Field, "，Msg =", ve.Msg)
	}
	// 注意 errors.As 的第二个参数必须是「指向实现了 error 的类型的指针」。
	// 因为 Error() 定义在 *ValidationError 上，所以这里用 *ValidationError。

	// errors.As 会穿透包装层，所以包过的错误也能提出来。
	// 传 -1 让它真的校验失败，这样才有错误可提取。
	_, wrapped := validateAndWrap("score", -1)
	fmt.Println("包装后的错误：", wrapped)
	var ve2 *ValidationError
	if errors.As(wrapped, &ve2) {
		fmt.Println("errors.As 穿透包装层提取到：Field =", ve2.Field, "，Msg =", ve2.Msg)
	}
	fmt.Println("errors.As 用的是「类型」匹配；上面 wrapped 的最外层是 *fmt.wrapError，内含 *ValidationError。")
	// 对比：errors.As 找「类型」，errors.Is 找「特定的那个值」。

	fmt.Println()
	fmt.Printf("=== fmt.Errorf 的 %%w 与 %%v ===\n")
	base := errors.New("底层错误")
	w1 := fmt.Errorf("用 %%w 包装：%w", base)
	w2 := fmt.Errorf("用 %%v 拼接：%v", base)
	fmt.Println("w1 =", w1, "-> errors.Is(w1, base) =", errors.Is(w1, base))
	fmt.Println("w2 =", w2, "-> errors.Is(w2, base) =", errors.Is(w2, base), "（%v 只是拼字符串，链断了）")
	// 为什么：%w 会返回一个带 Unwrap() 方法的新错误，保留错误链；
	// %v 只是把错误文本拼进新错误，新错误和 base 没有关系。
	// 需要让别人用 errors.Is/As 判断时，必须用 %w；只想加备注、不打算判断时 %v 也行。
	// 一个 fmt.Errorf 里只能有一个 %w（Go 1.20 起 errors.Join 可以合并多个错误）。

	fmt.Println()
	fmt.Println("=== 错误链的手工 Unwrap ===")
	fmt.Println("errors.Unwrap(w1) =", errors.Unwrap(w1))
	fmt.Println("errors.Unwrap(base) =", errors.Unwrap(base), "（没有包装层，返回 nil）")

	fmt.Println()
	fmt.Println("=== panic 与 recover（一）手动 panic ===")
	// 多返回值不能直接当 Println 的单个参数，先接住再打印。
	ok1, err1 := safeDiv(10, 2)
	if err1 != nil {
		fmt.Println("safeDiv(10, 2) 出错：", err1)
	} else {
		fmt.Println("safeDiv(10, 2) =", ok1)
	}
	// 除数为 0：函数内部 panic，被 defer 里的 recover 抓住，转成了普通的 error 返回值。
	bad1, err1 := safeDiv(1, 0)
	if err1 != nil {
		fmt.Println("safeDiv(1, 0) 返回错误：", err1, "（result 是零值", bad1, "）")
	}

	fmt.Println()
	fmt.Println("=== panic 与 recover（二）捕获运行时恐慌 ===")
	fmt.Println("catchOutOfRange() =", catchOutOfRange())

	fmt.Println()
	fmt.Println("=== 不该用 panic 的场景 ===")
	// 结论：可预期的失败（找不到、参数非法、IO 失败）用 error 返回值；
	// panic 只留给「程序有 bug、继续跑没意义」的情况（空指针、下标越界等），
	// 而且要让它尽量在测试里就炸出来。库代码不应该把 panic 抛给调用方。
	fmt.Println("本示例里只有演示 panic/recover 的地方才真的 panic。")
}

// findUser 模拟查用户：没有就返回哨兵错误 ErrNotFound。
func findUser(name string) (string, error) {
	if name == "alice" {
		return "Alice", nil
	}
	return "", ErrNotFound
}

// loadConfig 演示用 %w 包装哨兵错误，保留错误链。
func loadConfig(env string) (string, error) {
	if env == "prod" {
		// 用 %w 包装：既有上下文，又能被 errors.Is 认出来。
		return "", fmt.Errorf("加载 %s 配置失败: %w", env, ErrNotFound)
	}
	return "config-of-" + env, nil
}

// validate 返回自定义错误类型 *ValidationError。
func validate(field string, value int) (int, error) {
	if value < 0 {
		return 0, &ValidationError{Field: field, Msg: "不能为负数"}
	}
	return value, nil
}

// validateAndWrap 演示自定义错误也能被 %w 包装，再被 errors.As 提出来。
func validateAndWrap(field string, value int) (int, error) {
	v, err := validate(field, value)
	if err != nil {
		return 0, fmt.Errorf("处理 %s 时: %w", field, err)
	}
	return v, nil
}

// safeDiv 用 defer + recover 把除零 panic 转成普通返回值。
func safeDiv(a, b int) (result int, err error) {
	// defer 注册的匿名函数在函数返回前执行；recover 只有在 defer 里调用才有效。
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("safeDiv(%d, %d) 触发 panic: %v", a, b, r)
		}
	}()
	return a / b, nil
}

// catchOutOfRange 演示捕获运行时越界恐慌。
func catchOutOfRange() (msg string) {
	defer func() {
		if r := recover(); r != nil {
			// 越界 panic 的值是一个 runtime.Error，文本形如
			// "runtime error: index out of range [5] with length 3"。
			msg = fmt.Sprintf("已捕获：%v", r)
		}
	}()
	nums := []int{1, 2, 3}
	idx := 5
	return fmt.Sprint(nums[idx]) // 越界，触发 panic
}
