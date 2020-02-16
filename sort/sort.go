package main

import (
	"fmt"
	"os"
	"strings"
	"sort"
	"strconv"
)

func parse(ostream []string) (input, output string, flags map[string]int) {
	flags = make(map[string]int)
	if len(ostream) >= 2 {
		input = ostream[len(ostream) - 1]
		for idx := 1; idx <  len(ostream) - 1; idx++ {
		    if ostream[idx] != "-k" && ostream[idx] != "-o" {
				flags[ostream[idx]] = 0
			} else if ostream[idx] == "-o" {
				output = ostream[idx + 1]
				idx++
		    } else if ostream[idx] == "-k" {
				flags[ostream[idx]], _ = strconv.Atoi(ostream[idx + 1])
				idx++
			}
		}
	}
	return input, output, flags
}

func readFile(fileName string) (res []string, err error) {
	file, err := os.Open(fileName)
    if err != nil {
        return res, err
    }
    defer file.Close()
    stat, err := file.Stat()
    if err != nil {
        return res, err
    }
    binСontent := make([]byte, stat.Size() - 1)
    _, err = file.Read(binСontent)
    if err != nil {
        return res, err
    }
    res = strings.Split(strings.Trim(string(binСontent), "\r"), "\n")
    // res = res[0 : len(res) - 1]
    return res, err
}

func writeFile(data []string, fileName string) (err error) {
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

func applyFlags(data string, flags map[string]int) (res string) {
	res = data
    if col, has := flags["-k"]; has {
    	res = strings.Split(res, " ")[col]
    }
    if _, has := flags["-f"]; has {
    	res = strings.ToLower(res)

    }
    return res
}

// с флагом -n и числами 1.10 и 1.1 не сработает
func removeDuplicates(data []string, flags map[string]int) (res []string) {
    hasElem := make(map[string]bool)
	for _, elem := range data {
		hash := applyFlags(elem, flags)
		if _, has := hasElem[hash]; !has {
			hasElem[hash] = true
			res = append(res, elem)
		}
	}
	return res
}

func mySort(data []string, flags map[string]int) (res []string) {
	defer func() {
        if r := recover(); r != nil {
            fmt.Println(r)
            res = data
        }
    }()
	res = data
    if _, has := flags["-u"]; has {
    	res = removeDuplicates(res, flags)
    }
    sort.Slice(res, func(i, j int) bool {
    	l := applyFlags(res[i], flags)
    	r := applyFlags(res[j], flags)
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
	input, output, flags := parse(os.Args)
    if len(input) > 0 {
    	data, err := readFile(input)
    	res := mySort(data, flags)
    	if len(output) != 0 {
    		err = writeFile(res, output)
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
