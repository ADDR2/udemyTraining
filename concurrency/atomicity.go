package concurrency

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

var counter64 int64

/*ExecAtomicityExample is the main function of concurrency package*/
func ExecAtomicityExample() {
	fmt.Println("|-------Beginning of concurrency package-------|\n")
	wg.Add(2)
	go atomicIncrementor("Foo: ")
	go atomicIncrementor("Bar: ")
	wg.Wait()
	fmt.Println("Final counter: ", counter64)
	fmt.Println("\n|-------End of concurrency pacakge-------------|\n")
}

func atomicIncrementor(s string) {
	for i := 0; i < 20; i++ {
		time.Sleep(time.Duration(rand.Intn(3)) * time.Millisecond)
		atomic.AddInt64(&counter64, 1)
		fmt.Println(s, i, "Counter: ", counter64)
	}
	wg.Done()
}

//go run -race main.go
