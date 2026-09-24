package main

import "fmt"

// 两数之和
/*func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}*/
//哈希
func twoSum(nums []int, target int) []int {
	hashtable := make (map[int]int)
	for i := 0; i < len(nums); i++ {
		v := nums[i]
		if _, ok :=hashtable[target-v];ok{
			return []int{hashtable[target-v],i}
		}else{
			hashtable[v] = i}
		
	}
	return nil
}

func main() {
	fmt.Println("示例1:", twoSum([]int{2, 7, 11, 15}, 9), "期望 [0 1]")
	fmt.Println("示例2:", twoSum([]int{3, 2, 4}, 6), "期望 [1 2]")
	fmt.Println("示例3:", twoSum([]int{3, 3}, 6), "期望 [0 1]")
}
