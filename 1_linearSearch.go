package main

func linearSearch(arr []int, target int) int {
	for i, v := range arr {
		if v == target {
			return i
		}
	}
	return 1000000
}
