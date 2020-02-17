package main

import (
    "fmt"
    "os"
    "strconv"
    "strings"
)

type stack struct {
    data []interface{}
}

func (this *stack) top() (val interface{}) {
    return this.data[this.size() - 1]
}

func (this *stack) pop() (val interface{}) {
    size := this.size()
    val = this.data[size - 1]
    this.data = this.data[:size - 1]
    return val
}

func (this *stack) push(val ...interface{}) {
    this.data = append(this.data, val...)
}

func (this *stack) empty() bool {
    return len(this.data) == 0
}

func (this *stack) size() int {
    return len(this.data)
}

func getPriority(op string) int {
    switch op {
        case "+", "-" :
            return 0
        case "*", "/" :
            return 1
        default:
            return -1
    }
}

func makeOp(l int, r int, op string) int {
    switch op {
        case "+" :
            return l + r
        case "-" :
            return l - r
        case "*" :
            return l * r
        case "/" :
            return l / r
        default:
            return 0 // TODO
    } 
}

func makeOps(prevOps *stack, nums *stack, currOp string) {
    if !prevOps.empty() && currOp != "(" {
        if currOp == ")" {
            prevOp := prevOps.pop().(string);
            for prevOp != "(" {
                r, l := nums.pop().(int), nums.pop().(int)
                nums.push(makeOp(l, r, prevOp))
                prevOp = prevOps.pop().(string);
            }
            return
        }
        if prevOp := prevOps.top().(string); getPriority(prevOp) >= getPriority(currOp) {
            prevOps.pop()
            r, l := nums.pop().(int), nums.pop().(int)
            nums.push(makeOp(l, r, prevOp))
        }
    }
    prevOps.push(currOp)
}

func calc(expr []string) int {
    var nums, op stack
    for idx := 0; idx < len(expr); idx++ {
        if num, err := strconv.Atoi(expr[idx]); err == nil {
            nums.push(num)
        } else {
            makeOps(&op, &nums, expr[idx])
            if expr[idx] == "(" {
                if expr[idx + 1] == "+" || expr[idx + 1] == "-" {
                    idx++
                    if expr[idx] == "-" {
                        expr[idx + 1] = expr[idx] + expr[idx + 1]
                        expr[idx] = ""
                    }
                }
            }
        }
    }
    return nums.pop().(int)
}

func main() {
    if len(os.Args) == 2 {
        expr := "(" + os.Args[1] + ")"
        fmt.Println(expr)
        slice := strings.Split(expr, "") // string to slice
        ans/*, _*/ := calc(slice)
        fmt.Println(slice, " = ", ans)
    } else {
        fmt.Println("error")
    }
}
