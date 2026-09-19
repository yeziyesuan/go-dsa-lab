// 语法点：包与模块——首字母大小写决定导出与否；go mod init / build / run 的关系。
//
// 运行：go run ./syntax/10_package
//
// 前置条件（重要）：本文件 import 的 github.com/yeziyesuan/go-dsa-lab/pkgexample
// 必须真实存在（由仓库作者在另一个文件里创建，本目录不负责创建它）。
// 缺少那个包时 go vet ./... 会报 "no required module provides package"——
// 这是预期现象，不是本文件写错了。
package main

import (
	// 导入路径 = module 路径 + 子目录名。
	// module 是 github.com/yeziyesuan/go-dsa-lab（见 go.mod），
	// 所以包目录 D:\workday\go-dsa-lab\pkgexample 的导入路径就是下面这个。
	// 注意：导入路径里带不带 github.com 跟「本地能不能跑」无关，
	// 同模块内部一律用完整路径，go 会直接映射到本地磁盘目录，不走网络。
	"github.com/yeziyesuan/go-dsa-lab/pkgexample"

	"fmt"
)

func main() {
	fmt.Println("=== 导入自己的包并使用 ===")
	// 包名默认取目录名（这里是 pkgexample），所以调用要写 pkgexample.Xxx。
	// 下面两个都是首字母大写的导出标识符，包外可见。
	fmt.Println("pkgexample.Exported() =", pkgexample.Exported())
	fmt.Println("pkgexample.Greet(\"W1\") =", pkgexample.Greet("W1"))

	fmt.Println()
	fmt.Println("=== 首字母大小写决定导出与否 ===")
	fmt.Println("大写开头（Exported、Greet）= 导出，包外可以访问。")
	fmt.Println("小写开头（exported、greet）= 未导出，只有包内部能访问。")
	// 想访问 pkgexample 里未导出的标识符会编译失败，例如：
	//
	//	pkgexample.unexported()  // 编译失败：undefined: pkgexample.unexported
	//	var x pkgexample.hidden  // 编译失败：undefined (cannot refer to unexported name)
	//
	// 注意：导出与否只看「首字母是不是大写」这一个规则，
	// 没有 C++/Java 那种 public/private 关键字，也没有包级 friend 机制。

	fmt.Println()
	fmt.Println("=== 包名 / 目录名 / 导入路径 是三件事 ===")
	// 1) 导入路径：github.com/yeziyesuan/go-dsa-lab/pkgexample（go.mod 里的 module + 目录）
	// 2) 包名：源文件第一行的 package pkgexample（默认等于目录名，也可以不同）
	// 3) 引用名：代码里用 pkgexample.Xxx 或者用别名 mypkg.Xxx
	// 三者可以不一致。目录名带下划线时包名可以不带，比如目录 10_package
	// 里的包名就是 main 而不是 10_package（Go 不要求包名和目录名一致）。
	fmt.Println("本目录名是 10_package，但包名是 main，这就是「目录名 != 包名」的例子。")

	fmt.Println()
	fmt.Println("=== 导入别名、点导入、空白导入（这里只注释说明）===")
	// 别名：解决同名冲突或图省事
	//	import pkg "github.com/yeziyesuan/go-dsa-lab/pkgexample"
	//	然后写 pkg.Exported()
	//
	// 点导入：把包内导出标识符直接摊到当前文件（可读性差，不推荐）
	//	import . "github.com/yeziyesuan/go-dsa-lab/pkgexample"
	//	然后直接写 Exported()
	//
	// 空白导入：只为了触发包的副作用（比如注册驱动），不直接使用
	//	import _ "github.com/yeziyesuan/go-dsa-lab/pkgexample"
	//	本仓库约定不写 init()，所以基本用不到空白导入。
	fmt.Println("导入进去但没用到会编译失败（\"imported and not used\"），所以上面都是注释。")

	fmt.Println()
	fmt.Println("=== go mod init / go build / go run 的关系 ===")
	// go mod init github.com/yeziyesuan/go-dsa-lab
	//   在当前目录生成 go.mod，把这个目录标记成一个「模块」，
	//   模块路径就是 github.com/yeziyesuan/go-dsa-lab，它决定了所有内部导入的前缀。
	//   （已经改过 module 名：go mod edit -module github.com/<用户名>/go-dsa-lab）
	//
	// go mod tidy
	//   扫描源码，补齐/删除 go.mod 里需要的依赖。本仓库只用标准库，
	//   所以 go.mod 里不应该出现 require 行。
	//
	// go build ./...
	//   编译所有包，检查能不能过；默认把可执行文件生成在当前目录（main 包）。
	//   go build 不运行任何代码。
	//
	// go run ./syntax/10_package
	//   = 编译 + 立刻运行，产物放在临时目录里，不污染工作区。
	//   想固定产物就用 go build -o bin/demo.exe ./syntax/10_package。
	//
	// go install ./...
	//   编译并把可执行文件装到 $(go env GOPATH)/bin。
	//
	// 一句话：go.mod 定义模块边界与路径前缀；build 负责编译；run = build + 执行。
	fmt.Println("总结：go.mod 定路径前缀，go build 只编译，go run = 编译 + 运行。")
}
