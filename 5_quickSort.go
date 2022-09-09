package main

import "fmt"

func quickSort(arr []int, low, high int) []int {
	if low < high {
		var p int
		arr, p = partition(arr, low, high)
		arr = quickSort(arr, low, p-1)
		arr = quickSort(arr, p+1, high)
	}
	return arr
}

func partition(arr []int, low, high int) ([]int, int) {
	pivot := arr[high]
	fmt.Println(arr)
	for j := low; j < high; j++ {
		if arr[j] < pivot {
			fmt.Println(arr[j], pivot)
			arr[low], arr[j] = arr[j], arr[low]
			low++
		}
	}
	arr[low], arr[high] = arr[high], arr[low]
	fmt.Println("Low:", low)
	return arr, low
}
