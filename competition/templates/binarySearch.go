package main

import "fmt"

func binarySearch(target int, numbers []int) (int, int) {
	var count int
	for {
		count += 1
		if target == numbers[len(numbers)/2] {
			return target, count
		} else if target < numbers[len(numbers)/2] {
			numbers = numbers[:len(numbers)/2]
		} else if target > numbers[len(numbers)/2] {
			numbers = numbers[len(numbers)/2:]
		}
	}
}

func main() {
	target := 10
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println(numbers[:len(numbers)/2])
	fmt.Println(numbers[len(numbers)/2:])

	fmt.Println(numbers, target)
	fmt.Println(binarySearch(target, numbers))
}
