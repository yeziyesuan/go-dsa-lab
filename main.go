// Command go-dsa-lab 是 W1 的第一个可运行程序。
//
// 目标：验证「环境通」——go run 和 go test 都能跑。
// 用法：
//
//	go run .            输出 hello 与本周进度
//	go test ./...       跑全部测试
package main

import "fmt"

// Week1Goal 返回 W1 的唯一目标，被 main_test.go 用来验证测试链路。
func Week1Goal() string {
	return "环境通、语法入门、20 题落地、笔记体系立起来"
}

// Hello 返回给指定名字的问候语。
func Hello(name string) string {
	if name == "" {
		name = "Gopher"
	}
	return "hello, " + name
}

func main() {
	fmt.Println(Hello("W1"))
	fmt.Println("W1 目标：", Week1Goal())
	fmt.Println("下一步：go test ./...   然后 go run ./syntax/01_types")
}
