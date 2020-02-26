package main

import (
	"fmt"
)

type HeroNode struct {
	num  int
	name string
	next *HeroNode
}

func InsertNode(head *HeroNode, newnode *HeroNode) {
	temp := head
	for {
		if temp.next == nil {
			break
		}
		temp = temp.next
	}
	temp.next = newnode
}

func ListNode(head *HeroNode) {
	temp := head

	if temp.next == nil {
		fmt.Println("listnode is nil")
		return
	}

	for {
		fmt.Printf("[%d %s %p]", temp.next.num, temp.next.name, temp.next.next)
		temp = temp.next
		if temp.next == nil {
			fmt.Println("遍历完成")
			break
		}
	}

}

// InsertNodeSort 按照num的从小到大的顺序插入
func InsertNodeSort(head *HeroNode, newNode *HeroNode) {
	temp := head
	flag := true
	for {
		if temp.next == nil { //shu说明当前的是最后一个节点
			break
		} else if temp.next.num > newNode.num { //说明当前节点是最合适的节点,因为下一个节点的num大于新的num
			break
		} else if temp.next.num == newNode.num { //当前设定num相同不合法,如何不介意可以将上面的修改成>=.
			flag = false
			break
		}
		temp = temp.next
	}
	if !flag {
		fmt.Println("已经存在num为", newNode.num)
		return
	} else {
		newNode.next = temp.next
		temp.next = newNode
	}
}
func DelNodeByID(head *HeroNode, id int) {
	temp := head
	flag := false
	for {
		if temp.next == nil {
			break
		} else if temp.next.num == id {
			flag = true
			break
		}
		temp = temp.next
	}
	if flag {
		temp.next = temp.next.next
	} else {
		fmt.Println("未找到")
	}
}

func main() {
	head := &HeroNode{}

	hero1 := HeroNode{
		num:  1,
		name: "及时雨",
	}
	hero2 := HeroNode{
		num:  2,
		name: "玉麒麟",
	}
	hero4 := HeroNode{
		num:  4,
		name: "豹子头",
	}
	hero3 := HeroNode{
		num:  3,
		name: "公孙策",
	}
	// InsertNode(head, &hero1)
	// InsertNode(head, &hero2)
	InsertNodeSort(head, &hero1)
	InsertNodeSort(head, &hero2)
	InsertNodeSort(head, &hero4)
	InsertNodeSort(head, &hero3)
	ListNode(head)
	DelNodeByID(head, 2)
	ListNode(head)
	DelNodeByID(head, 2)
}
