package main

import (
	"fmt"
	"time"
)

func main() {
  task1() 
  task2()
  task3()
  task4()
}

func task1() {
	time.Sleep(100* time.Millisecond)
	fmt.Println("task1 is completed")
}

func task2() {
	time.Sleep(200* time.Millisecond)
	fmt.Println("task12 is completed")
}

func task3() {
   fmt.Println("task3 is completed")
}

func task4() {
	time.Sleep(50* time.Millisecond)
	fmt.Println("task4 is completed")
}