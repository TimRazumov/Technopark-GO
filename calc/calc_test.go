package main

import (
	"strconv"
	"testing"
	// "errors"
)

type In struct {
	L  int
	R  int
	Op string
}

type Out struct {
	Res int
	Err error
}

func TestBasicOp(t *testing.T) {
	t.Parallel()
	var opTests = []struct {
		in  In
		out Out
	}{
		{In{8, 2, "+"}, Out{10, nil}},
		{In{8, 2, "-"}, Out{6, nil}},
		{In{8, 2, "*"}, Out{16, nil}},
		{In{8, 2, "/"}, Out{4, nil}},
	}

	for idx, tt := range opTests {
		idx := idx
		tt := tt
		t.Run(strconv.Itoa(idx), func(t *testing.T) {
			t.Parallel()
			res, err := BasicOp(tt.in.L, tt.in.R, tt.in.Op)
			if res != tt.out.Res || err != tt.out.Err {
				t.Error("wrong answer")
			}
		})
	}
}

func TestStack(t *testing.T) {
	var s stackInt
	size := 50
	if !s.Empty() {
		t.Error("wrong answer")
	}
	for idx := 0; idx < size; idx++ {
		s.Push(idx)
	}
	if s.Size() != size {
		t.Error("wrong answer")
	}
	for idx := size - 1; idx > -1; idx-- {
		if s.Top() != idx {
			t.Error("wrong answer")
		}
		if s.Pop() != idx {
			t.Error("wrong answer")
		}
	}
	if !s.Empty() {
		t.Error("wrong answer")
	}
}

func TestCalc(t *testing.T) {
	// (3 + 7) * 3 / 10 == 3
	expr := []string{"(", "(", "3", "+", "7", ")", "*", "3", "/", "10", ")"}
	res, err := Calc(expr)
	if res != 3 || err != nil {
		t.Error("wrong answer")
	}
	// (-5 + 21) / 4
	expr = []string{"(", "(", "-", "5", "+", "21", ")", "/", "4", ")"}
	res, err = Calc(expr)
	if res != 4 || err != nil {
		t.Error("wrong answer")
	}
	// -30 * (1 + 2) / 5 == -18
	expr = []string{"(", "-30", "*", "(", "1", "+", "2", ")", "/", "5", ")"}
	res, err = Calc(expr)
	if res != -18 || err != nil {
		t.Error("wrong answer")
	}
	// invalid data
	expr = []string{"(", "3", "kek", "2", ")"}
	_, err = Calc(expr)
	if err == nil {
		t.Error("wrong answer")
	}
}
