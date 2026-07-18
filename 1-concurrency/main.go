package main

import (
	"fmt"
	"math/rand"
)

func main() {
	numCh := make(chan int, 10)
	readCh := make(chan int, 10)

	go func() {
		getSlice(numCh)
	}()

	go func() {
		getQuadNumbers(numCh, readCh)
	}()

	for n := range readCh {
		fmt.Println(n)
	}

}

func getQuadNumbers(num chan int, readCh chan int) {
	quadSl := make([]int, 10)
	for i := range 10 {
		n := <-num
		quadSl[i] = n * n
	}

	for _, v := range quadSl {
		readCh <- v
	}
	close(readCh)

}

func getSlice(ch chan int) {
	sl := make([]int, 10)
	for i := range 10 {
		num := rand.Intn(101)
		sl[i] = num
	}

	for _, n := range sl {
		ch <- n
	}
	close(ch)
}
