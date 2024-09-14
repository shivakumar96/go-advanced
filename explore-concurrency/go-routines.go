package goroutines

import (
	"fmt"
	"sync"
)

func subroutine(num int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("The square of %v is: %v\n", num, num*num)
}

func Mainroutine() {
	n := 10
	wg := &sync.WaitGroup{}
	wg.Add(1) // adding  thread
	go subroutine(n, wg)
	wg.Wait()
}

// Anonymous routine

func Anonymousroutine() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	n := 10
	// Anonymous function
	go func(num int) {
		defer wg.Done()
		fmt.Printf("Anonymous function: The square of %v is: %v\n", num, num*num)
	}(n)
	wg.Wait()
}
