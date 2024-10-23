package main

import (
	"fmt"
	"time"
)

// Example with normal funciton

func main() {
	now := time.Now()
	ch := make(chan struct{})
	go workNormal(ch)
	<-ch
	fmt.Println("elapsed:", time.Since(now))
	fmt.Println("done waiting, main exits")
}

func workNormal(ch chan struct{}) {
	time.Sleep(500 * time.Millisecond)
	fmt.Println("printing some stuff")
	ch<-struct{}{}
}

//Example with anonymous function
/*func main() {
	now := time.Now()
    ch := make(chan struct{})
	go func() {
		workAnonymous() 
		ch <- struct{}{}
	}()
	<-ch
	fmt.Println("elapsed:", time.Since(now))
	fmt.Println("done waiting, main exits")

}

func workAnonymous() {
	time.Sleep(500 * time.Millisecond)
	fmt.Println("printing some stuff")
}*/