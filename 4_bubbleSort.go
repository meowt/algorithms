package main

func bubbleSort(arr []int) []int {
	for i1 := 0; i1 < len(arr)-1; i1++ {
		for i2 := 0; i2 < len(arr)-1; i2++ {
			if arr[i2+1] < arr[i2] {
				arr[i2], arr[i2+1] = arr[i2+1], arr[i2]
			}
		}
	}
	return arr
}
