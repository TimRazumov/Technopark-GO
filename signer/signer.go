package main

import (
	"sort"
	"strconv"
	"strings"
	"sync"
)

func StartWorker(worker job, in, out chan interface{}, wg *sync.WaitGroup) {
	defer func() {
		close(out)
		wg.Done()
	}()
	worker(in, out)
}

func ExecutePipeline(workers ...job) {
	wg := &sync.WaitGroup{}
	wg.Add(len(workers))
	in := make(chan interface{}, MaxInputDataLen)
	for _, worker := range workers {
		out := make(chan interface{}, MaxInputDataLen)
		go StartWorker(worker, in, out, wg)
		in = out
	}
	wg.Wait()
}

// crc32(data)+"~"+crc32(md5(data))
func SingleHasher(data string, out chan interface{}, mu *sync.Mutex, pWg *sync.WaitGroup) {
	defer pWg.Done()
	wg := &sync.WaitGroup{}
	wg.Add(2)
	var first, second string
	go func() {
		first = DataSignerCrc32(data)
		wg.Done()
	}()
	go func() {
		mu.Lock()
		second = DataSignerMd5(data)
		mu.Unlock()
		second = DataSignerCrc32(second)
		wg.Done()
	}()
	wg.Wait()
	out <- first + "~" + second
}

func SingleHash(in, out chan interface{}) {
	mu := &sync.Mutex{}
	parentWg := &sync.WaitGroup{}
	for val := range in {
		parentWg.Add(1)
		go SingleHasher(strconv.Itoa(val.(int)), out, mu, parentWg)
	}
	parentWg.Wait()
}

// string res += crc32(th+data)) (th=0..5)
func MultiHasher(data string, out chan interface{}, pWg *sync.WaitGroup) {
	defer pWg.Done()
	const numTh = 6
	hash := make([]string, numTh)
	wg := &sync.WaitGroup{}
	wg.Add(numTh)
	for th := 0; th < numTh; th++ {
		go func(idx int) {
			hash[idx] = DataSignerCrc32(strconv.Itoa(idx) + data)
			wg.Done()
		}(th)
	}
	wg.Wait()
	out <- strings.Join(hash, "")
}

func MultiHash(in, out chan interface{}) {
	parentWg := &sync.WaitGroup{}
	for val := range in {
		parentWg.Add(1)
		go MultiHasher(val.(string), out, parentWg)
	}
	parentWg.Wait()
}

func CombineResults(in, out chan interface{}) {
	var res []string
	for val := range in {
		res = append(res, val.(string))
	}
	sort.Slice(res, func(i, j int) bool {
		return res[i] < res[j]
	})
	out <- strings.Join(res, "_")
}
