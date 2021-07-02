package main

import "fmt"

func main() {
	list := []int{5, 8, 7, 9, 1, 4, 2, 3, 5}
	//list := []int{1, 2, 3, 4, 5, 6}
	fmt.Printf("source:%v, addr:%v,%v \r\n", list, &list[0], &list[1])
	res, times := InsertionSort(list)
	fmt.Printf("result:%v, addr:%v,%v, swap times:%d \r\n", res, &res[0], &list[1], times)
}

func InsertionSort(param []int) ([]int, int) {
	length := len(param)
	swapTime := 0
	for back := 1; back < length; back++ {
		needSwap := false
		for front := 0; front < back; front++ {
			if param[back], param[front], needSwap = swapOnlyLess(param[back], param[front]); needSwap {
				swapTime++
			}
		}
	}
	return param, swapTime
}

// if the former larger than latter, swap and return
func swapOnlyLess(i, j int) (int, int, bool) {
	if i < j {
		return j, i, true
	}
	return i, j, false
}
