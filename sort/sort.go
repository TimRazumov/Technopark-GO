package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func Parse(ostream []string) (input, output string, flags map[string]int) {
	flags = make(map[string]int)
	if len(ostream) >= 2 {
		input = ostream[len(ostream)-1]
		for idx := 1; idx < len(ostream)-1; idx++ {
			if ostream[idx] != "-k" && ostream[idx] != "-o" {
				flags[ostream[idx]] = 0
			} else if ostream[idx] == "-o" {
				output = ostream[idx+1]
				idx++
			} else if ostream[idx] == "-k" {
				flags[ostream[idx]], _ = strconv.Atoi(ostream[idx+1])
				idx++
			}
		}
	}
	return input, output, flags
}

func ReadFile(fileName string) (res []string, err error) {
	file, err := os.Open(fileName)
	if err != nil {
		return res, err
	}
	defer file.Close()
	stat, err := file.Stat()
	if err != nil {
		return res, err
	}
	binСontent := make([]byte, stat.Size()-1)
	_, err = file.Read(binСontent)
	if err != nil {
		return res, err
	}
	res = strings.Split(strings.Trim(string(binСontent), "\r"), "\n")
	return res, err
}

func WriteFile(data []string, fileName string) (err error) {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()
	for _, elem := range data {
		fmt.Fprintln(file, elem)
	}
	return err
}

func ApplyFlags(data string, flags map[string]int) (res string) {
	res = data
	if col, has := flags["-k"]; has {
		res = strings.Split(res, " ")[col]
	}
	if _, has := flags["-f"]; has {
		res = strings.ToLower(res)

	}
	return res
}

func RemoveDuplicates(data []string, flags map[string]int) (res []string) {
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

func MySort(data []string, flags map[string]int) (res []string) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
			res = data
		}
	}()
	res = data
	if _, has := flags["-u"]; has {
		res = RemoveDuplicates(res, flags)
	}
	sort.Slice(res, func(i, j int) bool {
		l := ApplyFlags(res[i], flags)
		r := ApplyFlags(res[j], flags)
		_, reverse := flags["-r"]
		if _, has := flags["-n"]; has {
			lNum, _ := strconv.ParseFloat(l, 64)
			rNum, _ := strconv.ParseFloat(r, 64)
			if reverse {
				return lNum > rNum
			}
			return lNum < rNum
		}
		if reverse {
			return l > r
		}
		return l < r
	})
	return res
}

func main() {
	input, output, flags := Parse(os.Args)
	if len(input) > 0 {
		data, err := ReadFile(input)
		res := MySort(data, flags)
		if len(output) != 0 {
			err = WriteFile(res, output)
		} else {
			fmt.Println(res)
		}
		if err != nil {
			fmt.Println(err)
		}
	} else {
		fmt.Println("error")
	}
}
