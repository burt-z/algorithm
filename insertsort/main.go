package main

import (
	"fmt"
)

func InsortSort(arr *[6]int) {

	for i := 1; i < len(arr); i++ {
		val := arr[i]
		index := i - 1
		// 从大到小排列
		for index >= 0 && arr[index] < val {
			arr[index+1] = arr[index] // 如果前面的数比当前小,则将前面的数后移,直到index<0
			index--
		}
		if index+1 != i {
			arr[index+1] = val //将val赋值给不能比较的那个值,因为上面index--,所以这里需要加1
		}
	}
	fmt.Println(*arr)
}
func main() {
	arr := [6]int{12, 23, 10, 45, 56, 2}
	InsortSort(&arr)
}
