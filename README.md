# go-dsa-lab

> **W1（2026/9/21–9/27）的 Go 学习仓库** —— 语言方案 A·纯 Go。
> 对应计划：[W1-启动清单-Go主线.md](../W1-启动清单-Go主线.md) ｜ [计算机四大件-求职学习计划.md](../计算机四大件-求职学习计划.md) ｜ [总计划-统合版.md](../总计划-统合版.md)
> **本周唯一目标：环境通、语法入门、20 题落地、笔记体系立起来。**
> ✅ 本仓库已用 **go1.22.12 与 go1.27.1 两个版本**分别验证过：`go vet ./...`、`go build ./...`、`go test ./...`（27 个测试函数 / 298 个子测试 / 覆盖率 100%）全部通过，12 个语法示例全部 `go run` 成功 —— 见 [docs/验证记录.md](docs/验证记录.md)。

---

## 一、三条命令先跑通

```powershell
go version                  # 能看到版本号 = 环境通了
go run .                    # 输出 hello + W1 目标
go test ./...               # 所有测试通过 = 工具链通了
```

如果 `go` 不是内部或外部命令，先装 Go：

```powershell
winget install --id GoLang.Go -e              # 装完重开一个终端；MSI 路径保持默认
go env -w GOPROXY=https://goproxy.cn,direct   # 国内必设，否则 go get 会一直卡
```

> **本机现状（2026/9/19 实测复核）**：Go **go1.27.1 已装**在 `C:\Program Files\Go`，机器 PATH 已配好，**新开一个终端**就能用 `go`；`GOPROXY` 已设成 `https://goproxy.cn,direct`；`gopls v0.23.0` 与 `dlv 1.27.2` 已装（补全 / 跳转 / 断点可用）；git 2.55.0 已装并已推送首次提交。
> 完整的环境自查清单见 [docs/W1-使用说明.md](docs/W1-使用说明.md) 第 0 节。

---

## 二、仓库结构

| 路径 | 里面是什么 | 怎么用 |
|---|---|---|
| `main.go` / `main_test.go` | 第一个程序和第一个测试（证明链路通） | `go run .` |
| `syntax/` | **12 个语法示例**，一个语法点一个目录 | `go run ./syntax/03_slice` |
| `problems/` | **20 道 LeetCode 的 Go 实现 + 表驱动测试** | `go test ./problems/ -v` |
| `problems/notes/` | 每题一篇题解笔记（思路 / 卡点 / 复做日期） | 阅读，**新题照着补** |
| `notes/` | 知识点笔记：切片、map 并发、错误处理、Java→Go | 周日复盘时用 |
| `playground/` | **你自己的练习代码**（看来的变成写出来的，要提交） | `go run ./playground/01_vars` |
| `docs/` | [使用说明](docs/W1-使用说明.md)、[资料索引](docs/W1-资料索引.md)、[编写约定](docs/编写约定.md)、[验证记录](docs/验证记录.md)、[纸面自检四问](docs/纸面自检四问.md) | 先看使用说明 |

> 📦 **仓库地址**：<https://github.com/yeziyesuan/go-dsa-lab>（public，面试可展示）
> 提交节奏：**每天至少一次** `git add -A; git commit -m "..." ; git push`。
| `pkgexample/` | 演示「包与导出规则」的小包（大写导出、小写私有） | 被 `syntax/10_package` 引用 |
| `scripts/check.ps1` | **一键自检**：go 环境 → gofmt → go vet → go test → 逐个 `go run syntax/*` | `powershell -NoProfile -ExecutionPolicy Bypass -File .\scripts\check.ps1` |

> ⚠️ 这些代码是**起点不是答案**：先自己写，写不出来再看 `problems/`，看完隔天必须重做一遍。
> `problems/notes/` 里的「卡点」才是这个仓库最值钱的部分。

---

## 三、每天做什么（照 W1 启动清单的时间表）

> 逐日窗口以 [课表落位](../课表-2026秋-与计划映射.md) 为准；W1 因中秋调休，**上机日挪到周五 9/25**。

| 时间 | 内容 | 产出 |
|---|---|---|
| 周一 10:10–12:10 | 装环境 + `go mod init` + 跑通 `go run .` / `go test ./...` | 环境清单全打勾 |
| 周一–四 **课时段** | 语法精读 3 块 × 50 分钟，每块跟着 `syntax/` 敲一遍 | 12 个语法点过完 |
| 周一–四 **课时段**（纸上） | 刷题 2 道 + 写卡点与复杂度 | 8 题/周 |
| 周一–五 **白天空档 2h**（逐日不同） | 上机：语法点写进 `playground/`、把纸上的题敲进去验证 | 每天都有能跑的代码 |
| 周五白天 | 把本周语法点整理成可运行示例 + 写上机日清单（≤6 条） | 一次整洁的提交 |
| **上机日 08:00–15:30（6.5h）** | 清账 + 攻坚：标准库速览 `fmt strings strconv sort os bufio`、12 个示例重写、提交 | `syntax/12_stdlib` |
| 周日 08:00–12:00 | 复盘 + 补欠账 + 预习 W2（30 分钟复盘照统合版第五节） | `notes/04` 收尾 |

**每天至少一次提交**，哪怕只加一道题的题解——这是秋招时「持续学习」的硬证据。

---

## 四、本周验收（周日逐条打勾）

- [ ] `go version` / `go run .` / `go test ./...` 三条命令都正常
- [ ] GitHub 仓库 `go-dsa-lab` 已有 ≥ 5 次提交
- [ ] 20 题全部有题解，其中 **≥ 15 题独立 AC**
- [ ] 能用自己的话解释：切片和数组的区别、map 为什么并发不安全、Go 的错误处理为什么这么写
- [ ] `notes/` 下已有 ≥ 3 篇知识点笔记
- [ ] 周日复盘写完，并写下 W2 的三个重点

## 五、提交记录（自己填）

| 日期 | 提交内容 | commit |
|---|---|---|
| 9/21 | 环境 + hello + 第一个测试 | |
| 9/22 | syntax/01–04 | |
| 9/23 | 题 1、26 | |
| 9/24 | 题 27、121 | |
| 9/25 | 题 283、88 | |
| 9/26 | 题 53、15 | |
| 9/27 | 复盘 + notes 收尾 | |

---

## 六、W2 预告

Go 并发入门（goroutine / channel / select / sync）+ Linux 命令行 + Makefile / gdb。
`for range` 变量捕获、`select` 的默认分支、`sync.WaitGroup` 的 `Add` 位置——这三个坑 W2 会踩到，先记在 `notes/` 里。
