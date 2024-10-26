package main

import (
	"fmt"
	"sync"
)

func main() {
	// 1. Call Add() with the number of required operations to be waited on
	// 2. Call Done() inside each go routine (to be waited on)
	// 3. Call Wait() where you want to wait for the execution of all go routines

	// RULES
	// 1. Done() MUST be called as many times as Add()
	// 2. If calls to Done() are less than calls to Add(), it will result in a deadlock
	// 3. If calls to Done() are more than calls to Add, it will result in panic
	// 4. Calling Wait() without calling Add() will return immediately
	// 5. WaitGroup MUST always be passed by reference (as pointer)
	// 6. Calling another Wait() before the previous one returns results in panic

	var wg sync.WaitGroup

	wg.Add(4) 

	go func() {
     defer wg.Done()
     fmt.Println("Task 1 is completed")
	}()

	go func() {
     defer wg.Done()
     fmt.Println("Task 2 is completed")
	}()

	go func() {
     defer  wg.Done()		
     fmt.Println("Task 3 is completed")
	}()
	go func() {
     defer  wg.Done()		
     fmt.Println("Task 4 is completed")
	}()

   wg.Wait()
   fmt.Println("main program has executed")
}