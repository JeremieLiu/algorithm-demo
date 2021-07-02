package main

import "fmt"

func main() {
	list := []int{5, 8, 7, 9, 1, 4, 2, 3, 5}
	fmt.Printf("source:%v, addr:%v,%v \r\n", list, &list[0], &list[1])
	res := MergeSort(list)
	fmt.Printf("result:%v, addr:%v,%v \r\n", res, &res[0], &list[1])
}

func MergeSort(param []int) []int {
	length := len(param)
	s1 := 0
	e1 := length/2 - 1
	s2 := e1 + 1
	e2 := length
	tempIndex := 0
	tempArray := make([]int, length)

	for {
		if s1 < e1 && s2 < e2 { // 没有结束
			if param[s1] <= param[s2] {
				tempArray[tempIndex] = param[s1]
				s1++
				tempIndex++
			} else {
				tempArray[tempIndex] = param[s2]
				s2++
				tempIndex++
			}
		} else {
			break
		}
	}

	restStart := s1
	restEnd := e1
	if s2 < e2 {
		restStart = s2
		restEnd = e2
	}

	for _, v := range param[restStart : restEnd+1] {
		tempArray[tempIndex] = v
		tempIndex++
	}

	return tempArray
}
