package main

import (
    "fmt"
    "sync"
    "sync/atomic"
	"time"
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

func main() {
	var ok = true
	var recieved uint32
	freeFlowJobs := []job{
		job(func(in, out chan interface{}) {
			out <- 1
			time.Sleep(10 * time.Millisecond)
			currRecieved := atomic.LoadUint32(&recieved)
			fmt.Println("currRecieved", currRecieved)
			if currRecieved == 0 {
				ok = false
			}
		}),
		job(func(in, out chan interface{}) {
			for _ = range in {
				atomic.AddUint32(&recieved, 1)
				fmt.Println("recieved", recieved)
			}
		}),
	}
	ExecutePipeline(freeFlowJobs...)
	if !ok || recieved == 0 {
		fmt.Println("no value free flow - dont collect them")
	} else {
	    fmt.Println("ok")
	}
}