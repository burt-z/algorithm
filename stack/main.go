package main

import (
	"errors"

	"fmt"
)

type Stack struct {
	MaxTop int //栈的最大存储量
	Top    int //栈顶,因为栈底固定,所以可以不用设置,直接使用栈顶
	arr    [6]int
}

func (s *Stack) Push(x int) error {
	if s.MaxTop-1 == s.Top {
		return errors.New("栈满")
	}
	s.Top++
	s.arr[s.Top] = x
	return nil
}

func (s *Stack) Pop() (int, error) {
	if s.Top == -1 {
		fmt.Println("空栈")
		return 0, errors.New("栈空")
	}
	item := s.arr[s.Top]
	s.Top--
	return item, nil
}

func (s *Stack) Show() {
	if s.Top == -1 {
		fmt.Println("栈空")
		return
	}
	for i := s.Top; i >= 0; i-- {
		fmt.Println(s.arr[i])
	}
}

func main() {
	s := &Stack{
		MaxTop: 6,
		Top:    -1,
	}
	s.Show()
	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Push(4)
	s.Push(5)
	s.Show()
	s.Pop()
	s.Show()
}
