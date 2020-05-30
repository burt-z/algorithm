package main

import (
	"fmt"
)

func main(){
	arr := []int{3, 7, 9, 8, 38, 93, 12, 222, 45, 93, 23, 84, 65, 2}
	quickSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}

func quickSort(arr []int, left, right int) {
	if left < right {
		l, r := left, right
		mid := (left + right) / 2

		for l <= r {
			for arr[l] < arr[mid] {
				l++
			}

			for arr[mid] < arr[r] {
				r--
			}
			if l <= r {
				arr[l], arr[r] = arr[r], arr[l]
				l++
				r--
			}
		}

		if left < r {
			quickSort(arr, left, r)
		}
		if right > l {
			quickSort(arr, l, right)
		}

	}

}