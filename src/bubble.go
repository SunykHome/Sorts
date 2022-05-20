package main

import "fmt"

func main() {
	arr := [...]int{102, 3, 5, 24, 2}
	fmt.Print("排序前:", arr)
	bubbles(arr[:])
	fmt.Print("排序后:", arr)
}

func bubbles(arr []int) { //冒泡排序(提前结束+冒泡界优化)
	if len(arr) > 1 {
		var startBool bool = true //交换标志
		var lastSwappedIdx int = len(arr) - 1
		var swappedIdx int = -1

		for startBool != false {
			startBool = false
			for n := 0; n < lastSwappedIdx; n++ {
				if arr[n] > arr[n+1] {
					arr[n+1] = arr[n] ^ arr[n+1]
					arr[n] = arr[n] ^ arr[n+1]
					arr[n+1] = arr[n] ^ arr[n+1]
					startBool = true
					swappedIdx = n
				}
			}
			lastSwappedIdx = swappedIdx
		}
	}
}
