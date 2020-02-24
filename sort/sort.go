package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Flags struct {
	F bool
	U bool
	R bool
	N bool
	K int
}

func Parse(ostream []string) (input, output string, flags Flags) {
	flag.BoolVar(&flags.F, "f", false, "Case insensitive")
	flag.BoolVar(&flags.U, "u", false, "Remove duplicates")
	flag.BoolVar(&flags.R, "r", false, "Reverse result")
	flag.BoolVar(&flags.N, "n", false, "Sort as numbers")
	flag.IntVar(&flags.K, "k", -1, "Sort by column")
	flag.StringVar(&output, "o", "", "Output filename")
	flag.Parse()
	input = flag.Arg(0)
	return input, output, flags
}

func ReadFile(fileName string) (res []string, err error) {
	file, err := os.Open(fileName)
	if err != nil {
		return res, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		res = append(res, scanner.Text())
	}
	return res, err
}

func WriteFile(data []string, fileName string) (err error) {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()
	for _, elem := range data {
		_, err = fmt.Fprintln(file, elem)
		if err != nil {
			break
		}
	}
	return err
}

func ApplyFlags(data string, flags Flags) (res string) {
	res = data
	if flags.K != -1 {
		res = strings.Split(res, " ")[flags.K]
	}
	if flags.F {
		res = strings.ToLower(res)

	}
	return res
}

func RemoveDuplicates(data []string, flags Flags) (res []string) {
	hasElem := make(map[string]bool)
	for _, elem := range data {
		hash := ApplyFlags(elem, flags)
		if _, has := hasElem[hash]; !has {
			hasElem[hash] = true
			res = append(res, elem)
		}
	}
	return res
}

func CompInt(cond bool, l, r float64) bool {
	if cond {
		return l > r
	}
	return l < r
}

func CompString(cond bool, l, r string) bool {
	if cond {
		return l > r
	}
	return l < r
}

func MySort(data []string, flags Flags) (res []string) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
			res = data
		}
	}()
	res = data
	if flags.U {
		res = RemoveDuplicates(res, flags)
	}
	sort.Slice(res, func(i, j int) bool {
		l := ApplyFlags(res[i], flags)
		r := ApplyFlags(res[j], flags)
		ans := false
		if flags.N {
			lNum, _ := strconv.ParseFloat(l, 64)
			rNum, _ := strconv.ParseFloat(r, 64)
			ans = CompInt(flags.R, lNum, rNum)
		} else {
			ans = CompString(flags.R, l, r)
		}
		return ans
	})
	return res
}

func main() {
	input, output, flags := Parse(os.Args)
	if input == "" {
		log.Fatal("not enough arguments")
	}
	data, err := ReadFile(input)
	if err != nil {
		log.Fatal("invalid file")
	}
	res := MySort(data, flags)
	if output != "" {
		if wErr := WriteFile(res, output); wErr != nil {
			log.Fatal(wErr)
		}
	} else {
		fmt.Println(res)
	}

}
