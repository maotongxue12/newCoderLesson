package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main()  {
	var wg sync.WaitGroup
	var dogCounter uint64
	var fishCounter uint64
	var catCounter uint64
	dogCh := make(chan struct{}, 1)
	fishCh := make(chan struct{}, 1)
	catCh := make(chan struct{}, 1)

	wg.Add(3)
	go dog(&wg, dogCounter, dogCh, fishCh)
	go fish(&wg, fishCounter, fishCh, catCh)
	go cat(&wg, catCounter, catCh, dogCh)
	dogCh <- struct{}{}
	wg.Wait()
}

func dog(wg *sync.WaitGroup, counter uint64, dogCh, fishCh chan struct{})  {
	for {
		if counter >= uint64(100) {
			wg.Done()
			return
		}
		<- dogCh
		fmt.Println("dog")
		atomic.AddUint64(&counter, 1)
		fishCh <- struct{}{}
	}
}

func fish(wg *sync.WaitGroup, counter uint64, fishCh, catCh chan struct{})  {
	for {
		if counter >= uint64(100) {
			wg.Done()
			return
		}
		<- fishCh
		fmt.Println("fish")
		atomic.AddUint64(&counter, 1)
		catCh <- struct{}{}
	}
}

func cat(wg *sync.WaitGroup, counter uint64, catCh, dogCh chan struct{})  {
	for {
		if counter >= uint64(100) {
			wg.Done()
			return
		}
		<- catCh
		fmt.Println("cat")
		atomic.AddUint64(&counter, 1)
		dogCh <- struct{}{}
	}
}