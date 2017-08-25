package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	c := make(chan int)

	goNum := 5

	go Go(c)
	for i:=0; i<goNum; i++ {
		c <- 100 - i
		fmt.Println("fa", <- c)
	}


}

func Go(c chan int) {
	i := 0
	for {
		fmt.Println("pa", <- c)
		i ++
		c <- i
	}

}

