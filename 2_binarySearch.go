package main

import (
	"fmt"
)

func binarySearch(arr []int, target int) int {
	arr = selectionSort(arr)
	fmt.Println("Отсортированный массив:", arr)
	fmt.Println("Искомый элемент:", target)
	start := 0
	end := len(arr)
	found := false
	for found == false && start <= end {
		middle := (start + end) / 2
		if arr[middle] == target {
			found = true
			position := middle
			return position
		}
		if target < arr[middle] {
			end = middle - 1
		} else {
			start = middle + 1
		}
	}
	return 1000000
}
