package concurrency

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU()) // To use all the cores, the default uses more than one but not always all
}

/*ExecConcurrencyExample is the main function of concurrency package*/
func ExecConcurrencyExample() {
	fmt.Println("|-------Beginning of concurrency package-------|\n")
	wg.Add(2)
	go foo()
	go bar()
	wg.Wait()
	fmt.Println("\n|-------End of concurrency pacakge-------------|\n")
}

func foo() {
	for i := 0; i < 45; i++ {
		fmt.Println("Foo: ", i)
		time.Sleep(time.Duration(3 * time.Millisecond))
	}
	wg.Done()
}

func bar() {
	for i := 0; i < 45; i++ {
		fmt.Println("Bar: ", i)
		time.Sleep(time.Duration(3 * time.Millisecond))
	}
	wg.Done()
}
