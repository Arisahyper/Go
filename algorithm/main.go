package main

// O(n)
func lenearSearch(key int, nums []int) int {
	for i, v := range nums {
		if v == key {
			return i
		}
	}
	return -1
}

func main() {
	// 配列
	numberList := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	lenearSearch(1, numberList)
}
