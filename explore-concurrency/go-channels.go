package goroutines

import (
	"fmt"
	"sync"
	"time"
)

// communication between go routines using channnels

func printnumber(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for num := range ch {
		fmt.Printf("Normal channel %v \n", num)
		time.Sleep(75 * time.Millisecond)
	}
}

// read only channel
func printnumber2(ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for num := range ch {
		fmt.Printf("Read only channel %v \n", num)
		time.Sleep(75 * time.Millisecond)
	}
}

func PassNumberUsingChannel() {
	ch := make(chan int)
	wg := &sync.WaitGroup{}
	fmt.Println("------ starting channel -------")
	wg.Add(2)
	go printnumber(ch, wg)
	go printnumber2(ch, wg)
	for i := 1; i <= 10; i++ {
		ch <- i
	}
	close(ch) // channel has to be closed when using range operator
	fmt.Println("------ ending channel -------")
	wg.Wait()
}
