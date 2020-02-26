package main

import (
	"fmt"
)

type CatNode struct {
	num  int
	name string
	next *CatNode
}

func InsertNode(head *CatNode, newNode *CatNode) { //环形链表不再有空的节点,如果是第一个节点则进行赋值

	if head.next == nil {
		head.num = newNode.num
		head.name = newNode.name
		head.next = head //指向自己形成一个环
		return
	}
	temp := head
	for {
		if temp.next == head {
			break
		}
		temp = temp.next
	}

	temp.next = newNode
	newNode.next = head
}

func ListCicleLink(head *CatNode) {
	temp := head

	if temp.next == nil {
		fmt.Println("链表是空的")
		return
	}

	for {
		fmt.Printf("元素%d,%s,%p", temp.num, temp.name, temp.next) //因为第一个元素就有值
		// 所以不能temp.next再取值,不然会漏掉第一个
		if temp.next == head { //找到环形列表的最后一个
			break
		}
		temp = temp.next
	}
}

func DelNode(head *CatNode, id int) *CatNode {

	temp := head
	helper := head

	if temp.next == nil {
		fmt.Println("环形链表是空的")
		return head
	}

	//如果只有一个节点
	if temp.next == head {
		if temp.num == id {
			temp.next = nil
		}
		return head
	}

	for { //将helper定位到最后一个节点
		if helper.next == head {
			break
		}
		helper = helper.next
	}

	flag := true
	for {
		if temp.next == head { //最后一个节点 ,结束条件,不能整合到下面的代码里面temp.num == id
			if temp.num == id {
				helper.next = temp.next
			}
			break
		}
		if temp.num == id {
			if temp.next == head {
				helper.next = temp.next
				break
			}
			if temp == head { //删除头结点
				head = temp.next
			}
			helper.next = temp.next
			break
		}

		temp = temp.next
		helper = helper.next
	}

	return head
}

func main() {

	head := &CatNode{}
	ListCicleLink(head)
	cat1 := &CatNode{
		num:  1,
		name: "tom",
	}
	cat2 := &CatNode{
		num:  2,
		name: "json",
	}
	// cat3 := &CatNode{
	// 	num:  3,
	// 	name: "paton",
	// }
	InsertNode(head, cat1)
	InsertNode(head, cat2)
	// InsertNode(head, cat3)
	ListCicleLink(head)

	head = DelNode(head, 10)
	fmt.Println()
	fmt.Println()
	ListCicleLink(head)
}
