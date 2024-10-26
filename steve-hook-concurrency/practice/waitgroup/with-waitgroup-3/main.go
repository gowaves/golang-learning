package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	// Go Scheduler will create 16 threads
	// which in theory will be able to execute
	// 16 blocking go routines in parallel
	// totalling ~100ms
	fmt.Println("number of cores:", runtime.NumCPU())
    
	var wg sync.WaitGroup 
	wg.Add(10)
    for i:=0; i<10; i++ {
		go work(&wg, i + 1)
	}
	wg.Wait()
	time.Sleep(100 * time.Millisecond)
	fmt.Println("main is done")
}

func work(wg *sync.WaitGroup, id int) {
   defer wg.Done()	
   time.Sleep(100 * time.Millisecond)
   fmt.Println("task", id, "completed")
}

