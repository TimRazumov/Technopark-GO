package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type stack struct {
	data []interface{}
}

func (this *stack) Top() (val interface{}) {
	return this.data[this.Size()-1]
}

func (this *stack) Pop() (val interface{}) {
	size := this.Size()
	val = this.data[size-1]
	this.data = this.data[:size-1]
	return val
}

func (this *stack) Push(val ...interface{}) {
	this.data = append(this.data, val...)
}

func (this *stack) Empty() bool {
	return len(this.data) == 0
}

func (this *stack) Size() int {
	return len(this.data)
}

func GetPriority(op string) int {
	switch op {
	case "(", ")":
		return 0
	case "+", "-":
		return 1
	case "*", "/":
		return 2
	default:
		return -1
	}
}

func BasicOp(l int, r int, op string) (int, error) {
	switch op {
	case "+":
		return l + r, nil
	case "-":
		return l - r, nil
	case "*":
		return l * r, nil
	case "/":
		return l / r, nil
	default:
		return 0, errors.New("undef op")
	}
}

func MakeOps(prevOps *stack, nums *stack, currOp string) error {
	if GetPriority(currOp) == -1 {
		return errors.New("invalid expr")
	}
	if !prevOps.Empty() && currOp != "(" {
		if currOp == ")" {
			prevOp := prevOps.Pop().(string)
			for prevOp != "(" {
				r, l := nums.Pop().(int), nums.Pop().(int)
				if res, err := BasicOp(l, r, prevOp); err == nil {
					nums.Push(res)
				} else {
					return errors.New("invalid expr")
				}
				prevOp = prevOps.Pop().(string)
			}
			return nil
		}
		if prevOp := prevOps.Top().(string); GetPriority(prevOp) >= GetPriority(currOp) {
			prevOps.Pop()
			r, l := nums.Pop().(int), nums.Pop().(int)
			if res, err := BasicOp(l, r, prevOp); err == nil {
				nums.Push(res)
			} else {
				return errors.New("invalid expr")
			}
		}
	}
	prevOps.Push(currOp)
	return nil
}

func Calc(expr []string) (int, error) {
	var nums, op stack
	for idx := 0; idx < len(expr); idx++ {
		if num, err := strconv.Atoi(expr[idx]); err == nil {
			nums.Push(num)
		} else {
			if err := MakeOps(&op, &nums, expr[idx]); err != nil {
				return 0, err
			}
			if expr[idx] == "(" {
				if expr[idx+1] == "+" || expr[idx+1] == "-" {
					idx++
					if expr[idx] == "-" {
						expr[idx+1] = expr[idx] + expr[idx+1]
						expr[idx] = ""
					}
				}
			}
		}
	}
	return nums.Pop().(int), nil
}

func main() {
	if len(os.Args) == 2 {
		expr := "(" + os.Args[1] + ")"
		fmt.Println(expr)
		slice := strings.Split(expr, "") // string to slice
		if ans, err := Calc(slice); err == nil {
			fmt.Println(slice, " = ", ans)
		} else {
			fmt.Println(err)
		}
	} else {
		fmt.Println("print expr")
	}
}
