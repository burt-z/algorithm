package main

import (
	"errors"
	"fmt"
)

type Qunue struct {
	MacSize int
	array   [5]int
	front   int
	rear    int
}

type CirecleQunue struct {
	MaxSize int
	array   [4]int
	head    int
	tail    int
}

func (q *Qunue) AddQunue(v int) (err error) {
	if q.rear == q.MacSize-1 {
		fmt.Println("队列已满")
		return errors.New("qunue is full")
	}
	q.rear++
	q.array[q.rear] = v
	return
}

//遍历队列,找到队首和队尾,从队尾遍历到队首
func (q *Qunue) ShowQunue() (err error) {
	for i := q.front + 1; i <= q.rear; i++ {
		fmt.Printf("元素%d", q.array[i])
	}
	fmt.Println()
	return
}

func (q *Qunue) GetQunue() (v int, err error) {
	if q.front == q.rear {
		fmt.Println("qunue is null ")
		return -1, errors.New("qunue is null")
	}
	q.front++
	v = q.array[q.front]
	return v, nil
}

func (c *CirecleQunue) IsFull() bool {
	return (c.tail + +1)%c.MaxSize == c.head
}

func (c *CirecleQunue) IsEmpty() bool {
	return c.head == c.tail
}

func (c *CirecleQunue) Size() int {
	return (c.tail + c.MaxSize - c.head) % c.MaxSize
}

func (c *CirecleQunue) Push(v int) error {
	if c.IsFull() {
		return errors.New("queue is full")
	}
	c.array[c.tail] = v
	c.tail = (c.tail + 1) % c.MaxSize
	return nil
}

// Pop 弹出元素,先判读是否为空
func (c *CirecleQunue) Pop() (v int, err error) {
	if c.IsEmpty() {
		return 0, errors.New("queue is empty")
	}
	v = c.array[c.head]
	c.head = (c.head + 1) % c.MaxSize
	return v, nil
}

func (c *CirecleQunue) ListQueue() error {
	size := c.Size()
	if size == 0 {
		return errors.New("队列为空")
	}
	tempHead := c.head
	for i := 0; i < size; i++ {
		fmt.Printf("元素%d", c.array[tempHead])
		tempHead = (tempHead + 1) % c.MaxSize
	}
	fmt.Println()
	return nil
}

func main() {

	// var q Qunue = Qunue{
	// 	front:   -1,
	// 	rear:    -1,
	// 	MacSize: 5,
	// }

	var c CirecleQunue = CirecleQunue{
		head:    0,
		tail:    0,
		MaxSize: 4,
	}
	var v string
	for {
		fmt.Println("1.输入a添加元素")
		fmt.Println("2.输入s添加元素")
		fmt.Println("3.输入g添加元素")
		fmt.Scanln(&v)
		switch v {
		case "a":
			fmt.Println("s输入参数")
			var x int
			fmt.Scanln(&x)
			// err := q.AddQunue(x)
			err := c.Push(x)
			if err != nil {
				fmt.Println(err)
			}
		case "s":
			// q.ShowQunue()
			c.ListQueue()
		case "g":
			// x, err := q.GetQunue()
			x, err := c.Pop()
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(x)
		}
	}
}
