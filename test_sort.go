package main

import "fmt"

// 冒泡排序函数
func bubbleSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			// 如果当前元素比下一个元素大，交换它们
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func main() {
	// 示例数组
	arr := []int{64, 34, 25, 12, 22, 11, 90}

	fmt.Println("排序前:", arr)
	bubbleSort(arr)
	fmt.Println("排序后:", arr)
}
