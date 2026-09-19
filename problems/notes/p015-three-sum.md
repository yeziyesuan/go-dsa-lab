# 15. 三数之和

- 难度 / 类型：中等 · 排序 + 双指针 + 去重
- 思路（一句话）：排序后枚举第一个数 `nums[i]`，在 `nums[i+1:]` 上用左右指针夹逼凑 `-nums[i]`，三层去重（首元素回看去重 + 命中后 `l`/`r` 跳重复）。
- 复杂度：时间 O(n²)（含排序 O(n log n)，可忽略），空间 O(1) 不计返回值
- 卡点：去重是最容易翻车的地方。命中之后**必须**用 while 把 `nums[l]`、`nums[r]` 的所有重复值跳完，只移动一步的话下一轮又会凑出同一个三元组。首元素去重要用「回看」`nums[i] == nums[i-1]` 而不是「预判」`nums[i] == nums[i+1]`，预判写法在边界和漏跳上都更容易错。第三个坑是返回结果时误用切片共享底层数组（`append(res, nums[l:r+1])`），结果会随后续变化。
- Go 注意点：`sort.Ints(nums)` 是原地排序，会改调用方传进来的切片——题目允许，但测试里要先拷贝再调用。结果 `[][]int` 里每个三元组都要新建 `[]int{a,b,c}`，不能 append 子切片，否则共享底层数组。无解返回 `nil` 没问题，`len(nil) == 0`，调用方不会 panic；测试里把 `nil` 和空切片归一化后比较。
- 关键代码片段：

```go
l, r := i+1, n-1
for l < r {
	sum := nums[i] + nums[l] + nums[r]
	if sum < 0 {
		l++
	} else if sum > 0 {
		r--
	} else {
		res = append(res, []int{nums[i], nums[l], nums[r]})
		for l < r && nums[l] == nums[l+1] { l++ }
		for l < r && nums[r] == nums[r-1] { r-- }
		l++
		r--
	}
}
```

- 复做日期：9/26 ⬜ / 10/3 ⬜
