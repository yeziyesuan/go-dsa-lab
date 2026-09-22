# playground/ —— 你自己的练习代码写在这里

> **为什么有这个目录**：`syntax/` 和 `problems/` 里的代码是**参考资料**，不要在上面直接改。
> 你需要一个地方把「看来的」变成「自己写出来的」——就是这个目录。

## 约定

| 规则 | 说明 |
|---|---|
| 一个练习一个目录 | 目录名用 `01_vars`、`02_loop` 这种，和 `syntax/` 的编号尽量对应 |
| 每个目录是 `package main` | 有 `func main()`，能独立跑 |
| 文件叫 `main.go` | 一个目录一个文件就够了，练复杂了再拆 |
| 不引入第三方依赖 | 只用标准库 |
| 写完运行 | `go run ./playground/01_vars` |

> 🖊️ **写题时**：交卷前按 [纸面自检四问](../docs/纸面自检四问.md) 把「变量 → 边界 → 返回值 → 复杂度」各查一遍——这是纸面练习固定的最后 4 分钟。

## 建议的练习方式（对应 W1 第 2 节）

```powershell
cd D:\workday\go-dsa-lab
mkdir playground\01_vars
code playground\01_vars\main.go     # 在 VS Code 里新建并打开
go run ./playground/01_vars         # 跑起来看输出
```

1. 打开 `syntax/01_types/main.go` **看懂**；
2. 关掉它，在 `playground/01_vars/main.go` 里**凭记忆重敲**（允许卡住时回去瞄一眼，但别复制粘贴）；
3. 改几个值、故意写错一次，看看编译器骂你什么——**报错信息是最好的老师**；
4. 跑通后 `git add -A; git commit -m "playground: 变量与类型"`。

## 关于提交

这个目录**要提交**：它是你「每天都在写代码」的证据，面试时比 20 道题的题解更有说服力。
写得丑没关系，第一周就该丑。

## 注意

- 这里的 `main` 包**不会**被 `go test ./...` 当成测试，但它会被 `go build ./...` 编译——所以**别留编译不过的文件**（写一半就先注释掉或删掉）。
- 想快速检查有没有写坏：`go build ./...`，或者跑 `scripts/check.ps1`。
