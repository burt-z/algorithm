package main

import (
	"fmt"
)

// 使用递归的方式来模拟小老鼠出迷宫的问题.采用二维数组来模拟
// 1:表示墙,2表示通路,3:表示走过但是不通
// 最终位置为arr[6][5]==2

func SetWay(arr *[8][7]int, i, j int) bool {
	if arr[6][5] == 2 {
		return true
	} else {
		// 需要继续找,
		if arr[i][j] == 0 { //可探测,测试下一个位置是否可以通
			// 先将当前的设置成2,表示通路
			arr[i][j] = 2
			// 这里采用下右上左的顺序,如果采用上下左右会导致遍历所以的点
			if SetWay(arr, i+1, j) {
				return true
			} else if SetWay(arr, i, j+1) {
				return true
			} else if SetWay(arr, i-1, j) {
				return true
			} else if SetWay(arr, i, j-1) {
				return true
			} else {
				// 说明当前点的上下左右都不能走
				arr[i][j] = 3
				return false
			}
		} else {
			return false
		}
	}
}

func main() {
	var arr [8][7]int
	// 设置四周的边界
	for i := 0; i < 7; i++ {
		arr[0][i] = 1
		arr[7][i] = 1
	}
	for i := 0; i < 8; i++ {
		arr[i][0] = 1
		arr[i][6] = 1
	}

	arr[3][1] = 1
	arr[3][2] = 1

	SetWay(&arr, 1, 1)
	for i := 0; i < 8; i++ {
		for j := 0; j < 7; j++ {
			fmt.Print(arr[i][j], " ")
		}
		fmt.Println()
	}
}
