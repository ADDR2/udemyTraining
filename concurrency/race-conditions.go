package concurrency

import (
	"fmt"
	"math/rand"
	"time"
)

var counter int

/*ExecRaceConditionsExample is the main function of concurrency package*/
func ExecRaceConditionsExample() {
	fmt.Println("|-------Beginning of concurrency package-------|\n")
	wg.Add(2)
	go incrementor("Foo: ")
	go incrementor("Bar: ")
	wg.Wait()
	fmt.Println("Final counter: ", counter)
	fmt.Println("\n|-------End of concurrency pacakge-------------|\n")
}

func incrementor(s string) {
	for i := 0; i < 20; i++ {
		x := counter
		x++
		time.Sleep(time.Duration(rand.Intn(3)) * time.Millisecond)
		counter = x
		fmt.Println(s, i, "Counter: ", counter)
	}
	wg.Done()
}

//go run -race main.go
