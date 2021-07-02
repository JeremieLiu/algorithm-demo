package main

import "fmt"

func main() {
	list := []int{5, 8, 7, 9, 1, 4, 2, 3, 5}
	//list := []int{1, 2, 3, 4, 5, 6}
	fmt.Printf("source:%v, addr:%v,%v \r\n", list, &list[0], &list[1])
	res, times := SelectionSort(list)
	fmt.Printf("result:%v, addr:%v,%v, swap times:%d \r\n", res, &res[0], &list[1], times)
}

func SelectionSort(param []int) ([]int, int) {
	length := len(param)
	swapCount := 0
	for i := 0; i < length; i++ {
		min := param[i]
		index := i
		hasSwap := false
		for j := i + 1; j < length; j++ {
			if param[j] < min {
				min = param[j]
				index = j
			}
		}

		if index != i {
			if param[i], param[index], hasSwap = swap(param[i], param[index]); hasSwap {
				swapCount++
			}
		}
	}

	return param, swapCount
}

func swap(i, j int) (int, int, bool) {
	if i > j {
		return j, i, true
	}
	return i, j, false
}
