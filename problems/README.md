# problems/ 题解总索引（W1）

> W1 目标：**20 题全部有题解，其中 ≥ 15 题独立 AC**。
> 题号与题目以 [W1-启动清单-Go主线.md](../W1-启动清单-Go主线.md) 第四节为准，写作规范见 [docs/编写约定.md](../docs/编写约定.md)。

## 一、这个目录怎么用

每题一个实现文件 + 一个表驱动测试 + 一篇笔记，统一 `package problems`，只用标准库：

```powershell
go test ./problems/ -v                    # 跑全部题解测试（每题一个 TestXxx，逐个子测试列出）
go test ./problems/ -run TestTwoSum -v    # 只跑一道题：-run 接的是测试函数名
go vet ./problems/                        # 静态检查；本仓库不允许出现第三方依赖
```

`-run` 接的是**测试函数名**（如 `TestTwoSum`、`TestLongestConsecutive`），不是题号也不是文件名；
写 `-run TestLongestConsecutive/示例` 还能只跑某个子测试，想看覆盖率再补 `-cover`。

## 二、20 题总表

> 题目分组照抄启动清单第四节。表中「文件 / 笔记」两列都是**实际存在的文件**（20 / 20 已落地，已编译验证）。
> `_test.go` 与实现文件同名，不再单列。

### 数组 · 双指针（10 题）

| 题号 | 题目 | 类型 | 文件 | 笔记 | 完成 |
|---|---|---|---|---|---|
| 1 | 两数之和 | 数组 · 哈希 | `p001_two_sum.go` | `notes/p001-two-sum.md` | ⬜ |
| 26 | 删除有序数组中的重复项 | 数组 · 双指针 | `p026_remove_duplicates.go` | `notes/p026-remove-duplicates.md` | ⬜ |
| 27 | 移除元素 | 数组 · 双指针 | `p027_remove_element.go` | `notes/p027-remove-element.md` | ⬜ |
| 121 | 买卖股票的最佳时机 | 数组 · 贪心 | `p121_max_profit.go` | `notes/p121-max-profit.md` | ⬜ |
| 283 | 移动零 | 数组 · 双指针 | `p283_move_zeroes.go` | `notes/p283-move-zeroes.md` | ⬜ |
| 88 | 合并两个有序数组 | 数组 · 双指针（从后往前） | `p088_merge_sorted_array.go` | `notes/p088-merge-sorted-array.md` | ⬜ |
| 53 | 最大子数组和 | 数组 · 动态规划 / 贪心 | `p053_max_subarray.go` | `notes/p053-max-subarray.md` | ⬜ |
| 15 | 三数之和 | 数组 · 排序 + 双指针 | `p015_three_sum.go` | `notes/p015-three-sum.md` | ⬜ |
| 11 | 盛最多水的容器 | 数组 · 双指针 | `p011_max_area.go` | `notes/p011-max-area.md` | ⬜ |
| 42 | 接雨水 | 数组 · 双指针 / 单调栈 | `p042_trap.go` | `notes/p042-trap.md` | ⬜ |

### 字符串（8 题）

| 题号 | 题目 | 类型 | 文件 | 笔记 | 完成 |
|---|---|---|---|---|---|
| 344 | 反转字符串 | 字符串 · 双指针 | `p344_reverse_string.go` | `notes/p344-reverse-string.md` | ⬜ |
| 125 | 验证回文串 | 字符串 · 双指针 | `p125_valid_palindrome.go` | `notes/p125-valid-palindrome.md` | ⬜ |
| 14 | 最长公共前缀 | 字符串 · 纵向扫描 | `p014_longest_common_prefix.go` | `notes/p014-longest-common-prefix.md` | ⬜ |
| 242 | 有效的字母异位词 | 字符串 · 计数哈希 | `p242_valid_anagram.go` | `notes/p242-valid-anagram.md` | ⬜ |
| 49 | 字母异位词分组 | 字符串 · 哈希分组 | `p049_group_anagrams.go` | `notes/p049-group-anagrams.md` | ⬜ |
| 3 | 无重复字符的最长子串 | 字符串 · 滑动窗口 | `p003_length_of_longest_substring.go` | `notes/p003-length-of-longest-substring.md` | ⬜ |
| 5 | 最长回文子串 | 字符串 · 中心扩展 / 动态规划 | `p005_longest_palindrome.go` | `notes/p005-longest-palindrome.md` | ⬜ |
| 415 | 字符串相加 | 字符串 · 模拟竖式加法 | `p415_add_strings.go` | `notes/p415-add-strings.md` | ⬜ |

### 哈希（2 题）

| 题号 | 题目 | 类型 | 文件 | 笔记 | 完成 |
|---|---|---|---|---|---|
| 217 | 存在重复元素 | 哈希 · 集合 | `p217_contains_duplicate.go` | `notes/p217-contains-duplicate.md` | ⬜ |
| 128 | 最长连续序列 | 哈希 · 集合 | `p128_longest_consecutive.go` | `notes/p128-longest-consecutive.md` | ⬜ |

### 当前进度快照（2026-09-22 更新）

- 代码文件：**20 / 20**，每个都配了同名 `_test.go`；题解笔记也是 **20 / 20**。
- **已用真编译器验证过**（go1.22.12 初验；go1.27.1 于 2026/9/22 复核）：
  ```text
  go vet ./problems/ .     → 0 报错
  go test ./problems/ .    → ok  27 个测试函数 / 298 个子测试全部通过（覆盖率 100%）
  ```
- 上面的「完成」列故意全部留空：**那是留给你自己勾的**。
  这些代码是「参考答案」，先自己写，写完/做不出来再看；看过的题第二天必须重做一遍。
- 本索引不会自动更新：新增题目时请顺手补一行，以目录里的实际文件为准。

## 三、本周刷题纪律

1. **独立 AC ≥ 15 题**——先自己写，卡住也先想够 40 分钟；看了解法再写出来的不算独立 AC（但算「有题解」）。
2. **每题都要有笔记**——放 `problems/notes/`，模板里「卡点」和「复做日期」两栏**不许省**，这两栏才是笔记存在的理由。
3. **做不出来先看思路，隔天重做**——当天别硬耗；看完思路先关掉页面凭记忆写一遍，第二天不看答案重做，结果记到下面的错题表。
4. **每天至少一次提交**——哪怕只加一道题的题解；这是秋招时「持续学习」的硬证据。

## 四、错题重做记录

> 「结果」建议填：`独立 AC` / `看思路后 AC` / `没做出`。复做日期按笔记里的 `9/26`、`10/3` 两轮填。

| 题目 | 首次日期 | 结果 | 重做日期 | 结果 |
|---|---|---|---|---|
|  |  |  |  |  |
|  |  |  |  |  |
|  |  |  |  |  |
|  |  |  |  |  |
|  |  |  |  |  |
|  |  |  |  |  |
|  |  |  |  |  |
|  |  |  |  |  |

## 五、文件命名规则 & 怎么加第 21 题

**命名规则**（照抄现有文件，别自创）：

| 东西 | 规则 | 例子 |
|---|---|---|
| 实现文件 | `problems/pNNN_snake_slug.go`，题号补足 3 位，slug 用 LeetCode 英文标题的小写下划线缩写 | `p001_two_sum.go`、`p088_merge_sorted_array.go`、`p042_trap.go` |
| 测试文件 | 与实现**同名**加 `_test.go`，测试函数名 `Test` + 驼峰函数名 | `p042_trap_test.go` → `TestTrap` |
| 笔记 | `problems/notes/pNNN-hyphen-slug.md`，题号同样补 3 位，**slug 必须和实现文件一致**（下划线换成连字符） | `p042_trap.go` ↔ `notes/p042-trap.md` |
| 包名 | 一律 `package problems`；`problems/` 下不许有 `main`、`init()`、`panic` | — |
| 依赖 | 只用标准库，`go.mod` 里不允许出现 require | — |

**加第 21 题的步骤**：新建两个 `.go`（实现 + 测试）和一篇笔记 → `go vet ./problems/` → `go test ./problems/ -run TestXxx -v` 全绿 → 在上面的总表里补一行 → 提交。

实现文件骨架（导出函数必须有 doc comment、以函数名开头；注释里写清题目 / 思路 / 复杂度 / 卡点 / Go 注意点）：

```go
package problems

// MaximumSubarray 返回 nums 的最大子数组和（53. 最大子数组和）。
//
// 题目：
//
//	给定整数数组 nums，找出和最大的连续子数组并返回其和（子数组至少含一个元素）。
//
// 思路：
//
//	一遍扫，cur 表示「以当前元素结尾的最大子数组和」：cur = max(v, cur+v)，
//	答案取所有 cur 的最大值。cur+v 比 v 还小就说明前面的累加是负贡献，从 v 重新开始。
//
// 复杂度：
//
//	时间 O(n)，空间 O(1)。
//
// 卡点：
//
//	best 的初值不能取 0，否则全负数数组会错答 0（例如 [-3,-1,-2] 应返回 -1）。
//
// Go 注意点：
//
//	切片遍历用 for _, v := range nums；不要引入包级变量存中间状态。
func MaximumSubarray(nums []int) int {
	best, cur := nums[0], nums[0]
	for _, v := range nums[1:] {
		if cur < 0 {
			cur = v
		} else {
			cur += v
		}
		if cur > best {
			best = cur
		}
	}
	return best
}
```

表驱动测试骨架（示例 + 边界：单元素 / 全负数 / 重复 / 含零与负数；Go 1.22 起 `for range` 每轮是新变量，**不需要** `tt := tt`。注意 53 的输入保证非空，所以这里没有空数组用例，你自己那题若允许空输入就要补上）：

```go
package problems

import "testing"

func TestMaximumSubarray(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{name: "题目示例1", nums: []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, want: 6},
		{name: "单元素", nums: []int{1}, want: 1},
		{name: "全负数", nums: []int{-3, -1, -2}, want: -1},
		{name: "含重复", nums: []int{5, 5, 5}, want: 15},
		{name: "含零与负数", nums: []int{-1, 0, -2}, want: 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaximumSubarray(tt.nums); got != tt.want {
				t.Fatalf("MaximumSubarray(%v) = %d，期望 %d", tt.nums, got, tt.want)
			}
		})
	}
}
```

笔记骨架（`problems/notes/p053-max-subarray.md`）。注意「卡点」和「复做日期」两栏不许省：

````markdown
# 53. 最大子数组和

- 难度 / 类型：中等 · 数组 / 动态规划
- 思路（一句话）：
- 复杂度：时间 O(n)，空间 O(1)
- 卡点：
- Go 注意点：
- 关键代码片段：（3–8 行 go 代码块）
- 复做日期：9/26 ⬜ / 10/3 ⬜
````
