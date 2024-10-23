package main

import (
	"fmt"
	"sync"
	"time"
)
var wg sync.WaitGroup

//Example of direct call function 
func main() {
	now := time.Now()
	wg.Add(1)
	go workMain() 
	wg.Wait()
	fmt.Println("elapsed:", time.Since(now))
	fmt.Println("done waiting, main exits")
}

func workMain() {
	time.Sleep(500 * time.Millisecond)
	fmt.Println("printing some stuff")
	wg.Done()
}




//Example by anonymous function 
/*func main() {
	now := time.Now()
	wg.Add(1)
	go func() {
      defer wg.Done()
	  workAnonymous()
	}()
    wg.Wait()
	fmt.Println("elapsed:", time.Since(now))
	fmt.Println("done waiting, main exits")
}

func workAnonymous() {
	time.Sleep(500 * time.Millisecond)
	fmt.Println("printing some stuff")
}*/



