package main

import (
	"fmt"
	"sync"
)

// Example based on slice
func main() {
	requests := make([]int, 100)
	for i:=0; i<100; i++ {
		requests[i] = i + 1
	}
	var wg  sync.WaitGroup
	max := 10
	for i:=0; i< len(requests); i+=max  {
       for j := i; j< i+max; j++ {
		fmt.Println(j)
		 wg.Add(1)
		 go func(r int) {
			defer wg.Done()
			fmt.Println("Processing request", r)
		 }(requests[j])
	   }
	  wg.Wait()
	  fmt.Println(max, "requests processed") 
	}
}


// Video example based on anonymous function 
/*type request func()

func main() {
	//make it a map to simulate randomness of requests
	// When ranging over them
	requests := map[int]request{}

	for i := 1; i <= 100; i++ {
		f := func(n int) request {
			return func() {
				fmt.Println("request", n)
			}
		}
		requests[i] = f(i)
	}
    fmt.Println(len(requests))
	var wg  sync.WaitGroup
	max := 10
	for i:=0; i< len(requests); i+=max {
		for j:=i; j < i + max; j++ {
			wg.Add(1) 
			go func(r request) {
				defer wg.Done()
				r()
			}(requests[j+1])
		}
		wg.Wait()
		fmt.Println(max, "requests processed")
	}
}*/


