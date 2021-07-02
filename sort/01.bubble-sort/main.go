package main

import "fmt"

func main() {
	list := []int{5,8,7,9,1,4,2,3,5}
	fmt.Printf("source:%v, addr:%v,%v \r\n", list, &list[0], &list[1])
	res, times := BubbleSort(list)
	fmt.Printf("result:%v, addr:%v,%v, swap times:%d \r\n", res, &res[0], &list[1], times)
}

func BubbleSort(param []int) ([]int, int) {
	length := len(param)
	swapCount := 0
	for i := 0; i <= length -1; i++ {
		needSort := false
		for j := 0; j < length - i -1; j++ {
			if param[j] > param[j+1] {
				if param[j], param[j+1], needSort = swapOnlyLarger(param[j], param[j+1]); needSort {
					swapCount ++
				}
			}
		}
		if !needSort {
			break
		}
	}
	return param, swapCount
}

// if the former larger than latter, swap and return
func swapOnlyLarger(i, j int) (int, int, bool) {
	if i > j {
		return j, i, true
	}
	return i, j, false
}