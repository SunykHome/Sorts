package main

import "fmt"

func main() {
	arr := [...]int{102, 3, 5, 24, 2}
	fmt.Print("排序前:", arr)
	insert(arr[:])
	fmt.Print("排序后:", arr)
}

func insert(arr []int) { //折半插入排序
	if len(arr) > 1 {
		for i := 1; i <= len(arr)-1; i++ {
			if arr[i-1] <= arr[i] {
				continue
			}
			var target int = arr[i]
			var low int = 0
			var high int = i - 1
			for low < high {
				var center int = low + (high-low)/2
				if arr[center] <= target {
					low = center + 1
				} else {
					high = center - 1
				}
			}
			for j := i; j > low; j-- { // 移动
				arr[j] = arr[j-1]
			}
			arr[low] = target // 插入
		}
	}
}
