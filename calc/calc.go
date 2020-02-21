package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type stackInt struct {
	data []int
}

func (this *stackInt) Top() (val int) {
	return this.data[this.Size()-1]
}

func (this *stackInt) Pop() (val int) {
	size := this.Size()
	val = this.data[size-1]
	this.data = this.data[:size-1]
	return val
}

func (this *stackInt) Push(val ...int) {
	this.data = append(this.data, val...)
}

func (this *stackInt) Empty() bool {
	return len(this.data) == 0
}

func (this *stackInt) Size() int {
	return len(this.data)
}

type stackString struct {
	data []string
}

func (this *stackString) Top() (val string) {
	return this.data[this.Size()-1]
}

func (this *stackString) Pop() (val string) {
	size := this.Size()
	val = this.data[size-1]
	this.data = this.data[:size-1]
	return val
}

func (this *stackString) Push(val ...string) {
	this.data = append(this.data, val...)
}

func (this *stackString) Empty() bool {
	return len(this.data) == 0
}

func (this *stackString) Size() int {
	return len(this.data)
}

var opPriority = map[string]int{
	"(": 0,
	")": 0,
	"+": 1,
	"-": 1,
	"*": 2,
	"/": 2,
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

func MakeOps(prevOps *stackString, nums *stackInt, currOp string) error {
	if _, has := opPriority[currOp]; !has {
		return errors.New("invalid expr")
	}
	if !prevOps.Empty() && currOp != "(" {
		if currOp == ")" {
			prevOp := prevOps.Pop()
			for prevOp != "(" {
				r, l := nums.Pop(), nums.Pop()
				if res, err := BasicOp(l, r, prevOp); err == nil {
					nums.Push(res)
				} else {
					return errors.New("invalid expr")
				}
				prevOp = prevOps.Pop()
			}
			return nil
		}
		if prevOp := prevOps.Top(); opPriority[prevOp] >= opPriority[currOp] {
			prevOps.Pop()
			r, l := nums.Pop(), nums.Pop()
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
	var nums stackInt
	var op stackString
	for idx := 0; idx < len(expr); idx++ {
		if num, err := strconv.Atoi(expr[idx]); err == nil {
			nums.Push(num)
		} else {
			if err := MakeOps(&op, &nums, expr[idx]); err != nil {
				return 0, err
			}
			if expr[idx] == "(" && (expr[idx+1] == "+" || expr[idx+1] == "-") {
				idx++
				if expr[idx] == "-" {
					expr[idx+1] = expr[idx] + expr[idx+1]
					expr[idx] = ""
				}
			}
		}
	}
	return nums.Pop(), nil
}

func Parse(expr string) []string {
	tmp := "(" + strings.ReplaceAll(os.Args[1], " ", "") + ")"
	return strings.Split(tmp, "")
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("not enough arguments")
	}
	expr := Parse(os.Args[1])
	if ans, err := Calc(expr); err == nil {
		fmt.Println(expr, " = ", ans)
	} else {
		fmt.Println(err)
	}
}
