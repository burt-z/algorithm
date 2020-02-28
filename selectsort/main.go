package main

import (
	"fmt"
)

func SelectSort(arr *[5]int) {
	// 取其中一个依次和后面的进行比较,符合条件则交换
	for j := 0; j < len(arr)-1; j++ {
		// 假设第0个最大
		max := arr[j]
		// 最大的数的index 
		maxIndex := j
		for i := j + 1; i < len(arr); i++ {
			if max < arr[i] {
				// 修改最大值
				max = arr[i]
				// 修改最大值的index
				maxIndex = i
			}

		}
		// 交换条件
		if maxIndex != j {
			arr[j], arr[maxIndex] = arr[maxIndex], arr[j]
		}
		fmt.Println(*arr)
	}

}
func main() {
	x := [5]int{10, 34, 19, 100, 80}
	SelectSort(&x)
	fmt.Println(x)
}
