# 从 Java 转到 Go 的三个思维切换

- 一句话定位：Go 用**组合、返回值、零值**三件事，替换掉 Java 里的继承、异常和 null。
- 配套示例：`syntax/07_struct_method`、`syntax/08_interface`、`syntax/02_control`

## 切换一：没有继承，用组合 + 接口

Java 里想复用字段和方法，第一反应是 `extends`、`abstract class`，再围着类型树写多态。
Go 没有继承，只有两件事：**嵌入结构体拿字段和方法提升**，**接口描述行为**。

```java
// Java：先有类型树，多态靠重写
class Animal { void speak() {} }
class Dog extends Animal { @Override void speak() { System.out.println("woof"); } }
```

```go
// Go：嵌入 = has-a 的语法糖，方法被"提升"到外层
type Animal struct{ Name string }

func (a Animal) Speak() string { return a.Name + " 发声" }

type Dog struct {
	Animal // 匿名字段；Dog 直接就有 Name 和 Speak
}

// 行为用接口表达：谁有 Speak() 谁就是 Speaker，不用声明 implements
type Speaker interface{ Speak() string }

func say(s Speaker) { fmt.Println(s.Speak()) }

d := Dog{Animal{Name: "旺财"}}
say(d) // 旺财 发声
```

**接口是隐式实现的**，不需要 `implements`，所以接口通常定义在使用方（消费方）而不是
实现方；**接口描述行为而不是类型树**，`io.Reader`/`io.Writer` 就是"一个方法一个接口"的例子。
想替换嵌入来的方法不是"重写"，而是在外层定义同名方法把它遮蔽掉，要调原来那个就写
`d.Animal.Speak()`。

## 切换二：没有 try/catch，错误是返回值

```java
// Java：错误从调用点飞走，方法签名上看不出会失败
Config c = load(path);
int n = parse(c);
```

```go
// Go：每个可能失败的点都摆在面前，if err != nil 写满整屏是正常的
c, err := load(path)
if err != nil {
	return fmt.Errorf("加载配置失败: %w", err)
}
n, err := parse(c)
if err != nil {
	return fmt.Errorf("解析配置失败: %w", err)
}
_ = n
```

习惯转变：Java 里错误处理是"异常路径"，可以集中 `catch`；Go 里**错误处理是主路径的一部分**，
和业务代码交替出现。看到满屏 `if err != nil` 不要想去抽掉它，那是这门语言在逼你写清每一步
失败后做什么（详见 `notes/03-Go的错误处理为什么这么写.md`）。

## 切换三：没有 null，只有零值（还有 typed nil 陷坑）

Go 的每个类型都有零值：`string` 是 `""`、数值是 `0`、`bool` 是 `false`、
指针/切片/map/接口/函数/通道是 `nil`。所以 `var n int` 可以直接用，不像 Java 的
`Integer` 默认 null 那样一用就 NPE。但 `nil` 的 map 不可写、`nil` 的切片可 append，
脾气各不相同（见 `notes/01-切片与数组.md`、`notes/02-map为什么并发不安全.md`）。

**typed nil 陷阱**：接口变量的值由"类型 + 值"两部分组成，只有两部分都为 nil 它才等于 nil。

```go
type MyErr struct{}

func (e *MyErr) Error() string { return "boom" }

func bad() error {
	var p *MyErr = nil // 指针是 nil
	return p           // 装进 error 接口后，类型信息是 *MyErr，不是 nil
}

func main() {
	var i any = (*MyErr)(nil)
	fmt.Println(i == nil)     // false：类型部分不为空
	fmt.Println(bad() == nil) // false：经典的"返回了 nil 却不等于 nil"
}
```

修法：返回错误时别声明具体指针变量再返回它，**要返回就直接写 `return nil`**。

## 两个小切换

**没有方法重载**，同名函数只有一种签名，"可选参数"用选项模式或结构体参数代替。
**没有 `while`、没有三元运算符**：只有 `for`（`for cond {}`、`for {}`、三段式），
三元就写成普通的 `if`。

```java
// Java
int max = a > b ? a : b;
int i = 0;
while (i < 10) { i++; }
```

```go
// Go
max := b
if a > b { max = a } // 没有 ?:
i := 0
for i < 10 { i++ }   // while 就是 for
```

最后一条最容易咬人：**`for range` 拿到的永远是副本**，改它不影响原切片元素
（Java 的增强 for 也一样是副本，这点两门语言相同）。

```go
for _, x := range nums { x = 0 }    // 无效：x 是副本
for i := range nums { nums[i] = 0 } // 有效：必须按下标写回
```

## 一句话结论

把 Java 的**继承换成组合+隐式接口**、**异常换成 `if err != nil` 返回值**、
**null 换成各类型零值**（并小心 typed nil 让接口不等于 nil），再加上"没有重载/while/三元、
`for range` 拿的是副本"这几个小切换，就能开始用 Go 的方式思考了。

## 我踩过的坑

- 函数里 `var e *MyErr; return e`，调用方 `if err != nil` 永远成立，排查很久才想起接口里的类型信息不为 nil。
- 以为嵌入结构体等于继承，在外层"重写"方法后想用 `super`，实际要写 `d.Animal.Speak()`。
- 在 `range` 里给 `v` 赋值当修改元素，切片毫无变化。
- 到处想找 `?:`，最后接受多写两行 `if`——Go 优先可读性。
