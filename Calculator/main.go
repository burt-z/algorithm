package main

import (
	"errors"
	"fmt"
	"strconv"
)

// 用栈模拟计算器

// 数字栈
type NumStack struct {
	Top    int     //栈顶
	arr    [10]int //存入数字
	MaxTop int
}

type CharStack struct {
	Top    int
	arr    [10]string
	MaxTop int
}

func (s *NumStack) PushNum(x int) error {
	if s.MaxTop-1 == s.Top {
		return errors.New("栈满")
	}
	s.Top++
	s.arr[s.Top] = x
	fmt.Println("NumStack加人元素", x)
	return nil
}

func (s *NumStack) PopNum() (int, error) {
	if s.Top == -1 {
		fmt.Println("空栈")
		return 0, errors.New("栈空")
	}
	item := s.arr[s.Top]
	s.Top--
	fmt.Println("NumStack弹出元素", item)
	return item, nil
}

func (c *CharStack) PushChar(x string) error {
	if c.MaxTop-1 == c.Top {
		return errors.New("栈满")
	}
	c.Top++
	c.arr[c.Top] = x
	fmt.Println("CharStack加人元素", x)
	return nil
}

func (c *CharStack) PopChar() (string, error) {
	if c.Top == -1 {
		fmt.Println("空栈")
		return "", errors.New("栈空")
	}
	item := c.arr[c.Top]
	c.Top--
	fmt.Println("CharStack弹出元素", item)
	return item, nil
}

// IsChar 判断是否是字符串
func IsChar(x string) bool {
	if x == "+" || x == "-" || x == "/" || x == "*" {
		return true
	}
	return false
}

// Prority 返回字符串的优先级
func Prority(c string) int {
	if c == "*" || c == "/" {
		return 1
	}
	return 0
}

func Cal(num1, num2 int, c string) (res int) {
	switch c {
	case "*":
		res = num2 * num1
	case "/":
		res = num2 / num1
	case "+":
		res = num2 + num1
	case "-":
		res = num2 - num1
	}
	return res
}

func main() {
	str := "111+2*3-11"
	numStack := &NumStack{
		MaxTop: 10,
		Top:    -1,
	}

	charStack := &CharStack{
		MaxTop: 10,
		Top:    -1,
	}

	index := 0
	strKeep := ""
	// 数字和字符入栈
	for {
		c := str[index : index+1]
		if IsChar(c) {
			//如果是字符串判断栈是否是空,如果是空,可以字节入栈,不为空,判断栈顶的符号的优先级
			if charStack.Top == -1 {
				charStack.PushChar(c)
			} else {
				// 比较符号的优先级,如果栈顶的符号的优先级大于将要入栈的符号,则符号先出栈,并从数字栈中pop两个数
				if Prority(charStack.arr[charStack.Top]) > Prority(c) {
					s, _ := charStack.PopChar()
					num1, _ := numStack.PopNum()
					num2, _ := numStack.PopNum()
					res := Cal(num1, num2, s)
					//将计算的结果放到数字栈里面
					numStack.PushNum(res)
					//并且将当前的符号放到字符站里面
					charStack.PushChar(c)

				} else { //如果栈顶的优先级不大于当前元素的优先级,则符号入栈
					charStack.PushChar(c)
				}
			}

		} else {
			//数字直接入栈
			strKeep = strKeep + c
			// 单个数字可以直接入栈,多个数字的话需要拼接,
			if index == len(str)-1 { //说明是最后一个,不需要判断当前的下一个是不是字符le
				val, _ := strconv.ParseInt(strKeep, 10, 64)
				numStack.PushNum(int(val))
			} else {
				if IsChar(str[index+1 : index+2]) { //当前的数子先一个是字符
					val, _ := strconv.ParseInt(strKeep, 10, 64)
					numStack.PushNum(int(val))
					strKeep = ""
				}
			}

		}
		// 说明遍历到最后了
		if index+1 == len(str) {
			break
		}
		index++
	}
	//字符和数字出栈
	for {
		if charStack.Top == -1 {
			break
		}
		n1, _ := numStack.PopNum()
		n2, _ := numStack.PopNum()
		c, _ := charStack.PopChar()

		res := Cal(n1, n2, c)
		fmt.Println("计算结果", n1, n2, res, c)
		// 计算结果入栈
		numStack.PushNum(res)
	}
	// 最后的结果
	r, _ := numStack.PopNum()
	fmt.Println("最终结果", r)
}

func getDefault() {
	num := 1000 % 3
	fmt.Println("数字", num)
}
