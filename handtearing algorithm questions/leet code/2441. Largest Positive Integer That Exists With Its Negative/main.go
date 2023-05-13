// 2441. 与对应负数同时存在的最大正整数
// 简单
package main

func main() {
	nums := []int{-1, 2, -3, 3}
	println(findMaxK(nums))
}

func findMaxK(nums []int) int {
	max := -1
	hashMap := make(map[int]int)
	for _, v := range nums {
		if v > 0 && v > max {
			if _, ok := hashMap[-v]; ok {
				max = v
			} else {
				hashMap[v] = 1
			}
		} else if v < 0 && -v > max {
			if _, ok := hashMap[-v]; ok {
				max = -v
			} else {
				hashMap[v] = 1
			}
		}
	}
	return max
}
