package main

import "fmt"

func main() {
	arr := [...]int{102, 3, 5, 24, 2}
	fmt.Print("排序前:", arr)
	choose(arr[:])
	fmt.Print("排序后:", arr)
}

func choose(arr []int) {
	if len(arr) > 1 {
		for n := 0; n < len(arr)-1-n; n++ {
			var minIndex int = n //最小位
			var maxIndex int = n //最大位
			for j := n + 1; j < len(arr)-n; j++ {
				if arr[j] < arr[minIndex] {
					minIndex = j
				}
				if arr[j] > arr[maxIndex] {
					maxIndex = j
				}
			}
			if minIndex == maxIndex {
				break
			}
			if minIndex != n { //选择到最小不是此次n 进行调换
				arr[minIndex] = arr[n] ^ arr[minIndex]
				arr[n] = arr[n] ^ arr[minIndex]
				arr[minIndex] = arr[n] ^ arr[minIndex]
			}
			if maxIndex == n {
				maxIndex = minIndex
			}
			if maxIndex != n { //选择到最大不是此次n 进行调换
				arr[maxIndex] = arr[n] ^ arr[maxIndex]
				arr[n] = arr[n] ^ arr[maxIndex]
				arr[maxIndex] = arr[n] ^ arr[maxIndex]
			}
		}
	}
}
