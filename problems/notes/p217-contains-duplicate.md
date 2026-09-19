# 217. 存在重复元素

- 难度 / 类型：简单 · 哈希
- 思路（一句话）：一遍扫数组，用 `map[int]struct{}` 当集合，先查后存，查到即说明有重复。
- 复杂度：时间 O(n)，空间 O(n)
- 卡点：第一反应是「先排序再比相邻」，虽然能过，但要多付 O(n log n)，而且 `sort.Ints` 是**就地排序**，会把调用者切片里的元素改掉（切片共享底层数组）；另外总想给空数组/单元素加特判，其实循环自然就返回 `false`，不用写。判断 key 是否存在必须用 `_, ok :=`，不要用 `seen[v] == true` 反推。
- Go 注意点：map 的 value 用零大小的 `struct{}`（只关心 key 在不在），空 value 字面量写作 `struct{}{}`；`make(map[int]struct{}, len(nums))` 预分配容量，省掉扩容 rehash；`for _, v := range nums` 不需要下标就用 `_` 占位。
- 关键代码片段：

```go
seen := make(map[int]struct{}, len(nums))
for _, v := range nums {
	if _, ok := seen[v]; ok {
		return true
	}
	seen[v] = struct{}{}
}
```

- 复做日期：9/26 ⬜ / 10/3 ⬜
