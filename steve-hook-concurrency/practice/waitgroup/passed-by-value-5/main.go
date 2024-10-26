package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go work(wg)
	wg.Wait()
	fmt.Println("main is executed!")
}

func work(wg sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Work has been done")
}
