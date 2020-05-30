package main

import (
	"fmt"
	"os"
)

// 散列表

// EmpNode 用户信息
type EmpNode struct {
	ID   int
	Name string
	Next *EmpNode
}

// EmpLink 链表,用来存储emp用户信息
type EmpLink struct {
	Nodes *EmpNode
}

//HashTable 散列表,用来存储链表的数组
type HashTable struct {
	LinkArr [7]EmpLink
}

// Showlink 显示链表
func (empLink *EmpLink) Showlink(num int) {
	temp := empLink.Nodes
	if temp == nil {
		fmt.Printf("链表%d为空", num)
		fmt.Println()
		return
	}

	for {
		if temp != nil {
			fmt.Printf("链表%d 雇员id%d 名称%s", num, temp.ID, temp.Name)
			temp = temp.Next
		} else {
			break
		}
	}
	fmt.Println()
}

//InsertEmp 往链表中添加节点
func (empLink *EmpLink) InsertEmp(emp *EmpNode) {
	cur := empLink.Nodes   //辅助指针,指向当前节点
	var pre *EmpNode = nil //辅助指针,指向当前节点的前一个节点
	// 为空直接赋值退出
	if cur == nil {
		fmt.Println("1====")
		empLink.Nodes = emp
		fmt.Println("1====", empLink.Nodes)
		return
	}

	// 如果链表不为空
	for {
		if cur != nil {
			fmt.Println("2====")
			if cur.ID > emp.ID {
				fmt.Println("3====")
				break
			}
			pre = cur
			cur = cur.Next
		} else {
			break
		}
	}
	fmt.Println("4====")
	if cur == nil { //遍历到了最后一个
		pre.Next = emp
	} else {
		// 分为两种情况,一种是需要放到第一个head,一种是不放到第一个
		if pre == nil {
			emp.Next = cur
			empLink.Nodes = emp
		} else {
			pre.Next = emp
			emp.Next = cur
		}
	}
}

//Insert 散列表添加元素
func (h *HashTable) Insert(e *EmpNode) {
	linkNum := h.HashFunc(e.ID)
	empLink := &(h.LinkArr[linkNum])
	fmt.Println(e)
	empLink.InsertEmp(e)
}

//HashFunc 散列方法.计算存储到那个链表中,采用用户id/7余数的方式
func (h *HashTable) HashFunc(id int) int {
	fmt.Println("arr id ", id%7)
	return id % 7
}

// ShowAll 显示散列
func (h HashTable) ShowAll() {
	for i := 0; i < 7; i++ {
		h.LinkArr[i].Showlink(i)
	}
}

func main() {
	key := ""
	id := 0
	name := ""
	var hashTable HashTable
	for {
		fmt.Println("================菜单==============")
		fmt.Println("输入1,添加")
		fmt.Println("输入2,显示")
		fmt.Println("输入3,退出")
		fmt.Scanln(&key)
		switch key {
		case "1":
			fmt.Println("输入id")
			fmt.Scanln(&id)
			fmt.Println("输入名称")
			fmt.Scanln(&name)
			emp := &EmpNode{
				ID:   id,
				Name: name,
			}
			hashTable.Insert(emp)
		case "2":
			hashTable.ShowAll()
		case "3":
			os.Exit(0)
		}
	}

}
