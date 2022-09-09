package main

import "fmt"

func main() {
	selectAlgorithm(5)
}

func selectAlgorithm(algNum int) {
	switch algNum {
	case 1:
		arr := []int{1, -1, 5, 2, 16, -3, 51}
		fmt.Println("Идентификатор элемента = ", linearSearch(arr, -3))
	case 2:
		arr := []int{1, -1, 5, 2, 16, -3, 51}
		fmt.Println("Идентификатор элемента = ", binarySearch(arr, 5))
	case 3:
		arr := []int{1, -1, 5, 2, 16, -3, 51}
		fmt.Println("Отсортированный массив:", selectionSort(arr))
	case 4:
		arr := []int{1, -1, 5, 2, 16, -3, 51}
		fmt.Println("Отсортированный массив:", bubbleSort(arr))
	case 5:
		arr := []int{1, -1, 5, 2, 16, -3, 5}
		fmt.Println("Отсортированный массив:", quickSort(arr, 0, len(arr)-1))
	case -1:
		arr := []int{1, -1, 5, 2, 16, -3, 51}
		fmt.Println("Отсортированный массив:", test(arr))
	default:
		fmt.Println("Введен неправильный номер алгоритма")
	}
}
