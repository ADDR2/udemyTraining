package concurrency

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var mutex sync.Mutex

/*ExecMutexExample is the main function of concurrency package*/
func ExecMutexExample() {
	fmt.Println("|-------Beginning of concurrency package-------|\n")
	wg.Add(2)
	go mutexIncrementor("Foo: ")
	go mutexIncrementor("Bar: ")
	wg.Wait()
	fmt.Println("Final counter: ", counter)
	fmt.Println("\n|-------End of concurrency pacakge-------------|\n")
}

func mutexIncrementor(s string) {
	for i := 0; i < 20; i++ {
		time.Sleep(time.Duration(rand.Intn(3)) * time.Millisecond)
		mutex.Lock()
		counter++
		fmt.Println(s, i, "Counter: ", counter)
		mutex.Unlock()
	}
	wg.Done()
}

//go run -race main.go
