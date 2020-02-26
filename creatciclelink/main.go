package main

import (
	"fmt"
)

type Boy struct {
	num  int
	next *Boy
}

func AddBoy(num int) *Boy {
	frist := &Boy{}
	curBoy := &Boy{}

	if num < 1 {
		fmt.Println("链表空的...")
		return frist
	}

	for i := 1; i <= num; i++ {
		boy := &Boy{
			num: i,
		}

		if i == 1 {
			frist = boy
			curBoy = boy
			curBoy.next = frist //成环
		} else {
			curBoy.next = boy   //添加下一个
			curBoy = boy        //curboy指针后移
			curBoy.next = frist //指向头,成环
		}
	}
	return frist
}

func ShowLink(frist *Boy) {
	if frist.next == nil {
		fmt.Println("空的")
		return
	}

	temp := frist

	for {
		fmt.Println("元素", temp.num)
		if temp.next == frist {
			break
		}
		temp = temp.next
	}

}

func PlayGame(frist *Boy, startNum, count int) {
	if frist.next == nil {
		fmt.Println("链表为空")
		return
	}

	helper := frist
	for { //找到最后的节点.
		if helper.next == frist {
			break
		}
		helper = helper.next
	}

	for i := 1; i <= startNum-1; i++ { //指针移动需要开始的节点
		frist = frist.next
		helper = helper.next
	}

	for {
		for i := 1; i <= count-1; i++ { //移动的个数
			frist = frist.next
			helper = helper.next
		}
		fmt.Printf("出圈的元素id%d\n", frist.num)
		frist = frist.next  //frist指针往前移动,跳过当前的frist的节点
		helper.next = frist //后面的辅助指针指向移动后的frist,成环
		if frist == helper {
			break
		}
	}
	fmt.Println("最后一个元素", frist.num)

}

func main() {
	frist := AddBoy(5)
	ShowLink(frist)
	fmt.Println()
	PlayGame(frist, 2, 3)
}
