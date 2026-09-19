// Package pkgexample 是给 syntax/10_package 用的示例包。
//
// 它存在的唯一目的，是让你亲眼看到 Go 的「导出规则」：
//   - 首字母大写的标识符（Greeting、Greet）包外可见；
//   - 首字母小写的标识符（hidden、helper）只有包内可见，
//     包外写 pkgexample.hidden 会直接**编译失败**（不是运行时报错）。
//
// 另外注意：包名可以和目录名不同，但惯例是保持一致；
// import 的是**目录路径**，调用时用的是**包名**。
package pkgexample

import "strings"

// Greeting 是导出常量，包外可读：pkgexample.Greeting。
const Greeting = "你好"

// hidden 是小写常量，包外不可见 —— 这就是 Go 的「私有」。
const hidden = "（这句来自包内部的私有常量）"

// Exported 返回一句固定的话，用来验证包外能调到导出函数。
func Exported() string {
	return Greeting + "，这是 pkgexample 包的导出函数" + hidden
}

// Greet 返回带名字的问候；空名字时回退到默认称呼。
func Greet(name string) string {
	name = strings.TrimSpace(name)
	if name == "" {
		name = "Gopher"
	}
	return Greeting + "，" + name
}
