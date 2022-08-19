package channel_algorithm

import (
	"bytes"
	"fmt"
	"net/http"
	"runtime"
	"strconv"
	"sync"
	"sync/atomic"
)

//题目一：实现10个链接访问百度1000次
//实现一：使用sync包
func connectRequestWeb(connectNum, totalVisitNum int, url string) {
	wg := sync.WaitGroup{}
	requestChan := make(chan struct{}, connectNum)
	for i := 0; i < totalVisitNum; i++ {
		requestChan <- struct{}{}
		wg.Add(1)
		go func() {
			defer wg.Done()
			visitWeb(url)
			<-requestChan
		}()
	}
	wg.Wait()
}

func visitWeb(url string) {
	_, err := http.Get(url)
	if err != nil {
		fmt.Println("connect is successful: ")
	}
}

func getGID() uint64 {
	b := make([]byte, 64)
	index := runtime.Stack(b, false)
	b = b[:index]
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	n, _ := strconv.ParseUint(string(b), 10, 64)
	return n
}

//题目二：3 个函数分别打印 cat、dog、fish，要求每个函数都要起一个 goroutine，
//按照 cat、dog、fish 顺序打印在屏幕上 100 次。
func printCat(wg *sync.WaitGroup, chCat, chDog chan struct{}, count uint64, totalNum int) {
	for {
		if count > uint64(totalNum) {
			wg.Done()
			return
		}
		<-chCat
		fmt.Println("this is cat...")
		atomic.AddUint64(&count, 1)
		chDog <- struct{}{}
	}

}

func printDog(wg *sync.WaitGroup, chDog, chFish chan struct{}, count uint64, totalNum int) {
	for {
		if count > uint64(totalNum) {
			wg.Done()
			return
		}
		<-chDog
		fmt.Println("this is dog...")
		atomic.AddUint64(&count, 1)
		chFish <- struct{}{}
	}
}

func printFish(wg *sync.WaitGroup, chFish, chCat chan struct{}, count uint64, totalNum int) {
	for {
		if count > uint64(totalNum) {
			wg.Done()
			return
		}
		<-chFish
		fmt.Println("this is fish...")
		atomic.AddUint64(&count, 1)
		chCat <- struct{}{}
	}

}

func printTotal(totalNum int) {
	wg := sync.WaitGroup{}
	wg.Add(3)
	var count uint64 = 0
	chCat := make(chan struct{}, 1)
	chDog := make(chan struct{}, 1)
	chFish := make(chan struct{}, 1)
	chCat <- struct{}{}
	go printCat(&wg, chCat, chDog, count, totalNum)
	go printDog(&wg, chDog, chFish, count, totalNum)
	go printFish(&wg, chFish, chCat, count, totalNum)
	wg.Wait()
	close(chCat)
	close(chDog)
	close(chFish)
}
