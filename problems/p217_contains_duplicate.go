package problems

// ContainsDuplicate 判断 nums 中是否存在重复元素（217. 存在重复元素）。
//
// 题目：
//
//	给定整数数组 nums，只要任一元素出现两次及以上就返回 true；
//	每个元素互不相同（含空输入）返回 false。
//
// 思路：
//
//	一遍扫数组，用 map[int]struct{} 当集合：先查后存。遍历到 v 时先看 v 在不在集合里，
//	在就说明之前出现过，直接返回 true；不在就把 v 存进去。扫完没命中就是没有重复。
//
//	对比「先排序」的写法与取舍：
//	  排序法 —— 先 sort.Ints(nums) 就地排序，再比较相邻元素是否相等：时间 O(n log n)，
//	  空间只有排序递归栈的 O(log n)（原地排序）。n 小、本来就要求有序、或者不许用
//	  额外 O(n) 空间时更划算；代价是慢一档，而且 sort.Ints 会**改掉调用者切片里的元素**
//	  （切片共享底层数组，这个坑在 notes/ 里单独记了一篇）。
//	  本题只问「有没有重复」、不关心顺序，所以选哈希法：一遍扫完，均摊 O(n)。
//
// 复杂度：
//
//	时间 O(n)：每个元素进表一次、查表一次，均摊 O(1)。
//	空间 O(n)：最坏情况无重复，集合要存下全部 n 个数。
//
// 卡点：
//
//  1. 第一反应是「先排序再比相邻」，能过但多付 O(n log n)，还要小心排序改了入参；
//     要能说清这里为什么不用它（见上面的取舍）。
//  2. map 的 value 写成 bool 也能过，但 bool 占一个字节、语义也不对；用零大小的
//     struct{} 表达「只关心 key 在不在」，空 value 字面量是 struct{}{}。
//  3. 总想给空切片 / 单元素加特判，其实不需要：空切片循环不执行、单元素循环只走一次。
//
// Go 注意点：
//
//   - 判断 key 是否存在用 `if _, ok := seen[v]; ok`，不要写成 `seen[v] == true`：
//     一旦以后 value 类型变了就会误判，ok 才是标准写法。
//   - make(map[int]struct{}, len(nums)) 预分配容量，省掉扩容时的 rehash。
//   - for _, v := range nums 不需要下标就用 _ 占位，Go 不允许只取下标却不用值。
func ContainsDuplicate(nums []int) bool {
	seen := make(map[int]struct{}, len(nums))
	for _, v := range nums {
		if _, ok := seen[v]; ok {
			return true
		}
		seen[v] = struct{}{}
	}
	return false
}
