package channels

import "fmt"
import "time"
import "sync"

/*ExecChannelsExample is the main function of channels package*/
func ExecChannelsExample() {
	fmt.Println("|-------Beginning of channels package-------|\n")
	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
	}()

	go func() {
		for {
			fmt.Println(<-c)
		}
	}()

	time.Sleep(time.Second)
	fmt.Println("\n|-------End of channels pacakge-------------|\n")
}

/*ExecRangeExample is the main function of channels package*/
func ExecRangeExample() {
	fmt.Println("|-------Beginning of channels package-------|\n")
	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		close(c)
	}()

	for i := range c {
		fmt.Println(i)
	}
	fmt.Println("\n|-------End of channels pacakge-------------|\n")
}

/*ExecNToOneExample is the main function of channels package*/
func ExecNToOneExample() {
	fmt.Println("|-------Beginning of channels package-------|\n")
	c := make(chan int)

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		wg.Done()
	}()

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		wg.Done()
	}()

	go func() {
		wg.Wait()
		close(c)
	}()

	for i := range c {
		fmt.Println(i)
	}
	fmt.Println("\n|-------End of channels pacakge-------------|\n")
}

/*ExecSemaphoreExample is the main function of channels package*/
func ExecSemaphoreExample() {
	fmt.Println("|-------Beginning of channels package-------|\n")
	c := make(chan int)
	done := make(chan bool)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		done <- true
	}()

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		done <- true
	}()

	go func() {
		<-done
		<-done
		close(c)
	}()

	for i := range c {
		fmt.Println(i)
	}
	fmt.Println("\n|-------End of channels pacakge-------------|\n")
}

/*ExecMultipleSemaphoreExample is the main function of channels package*/
func ExecMultipleSemaphoreExample() {
	fmt.Println("|-------Beginning of channels package-------|\n")
	n := 10
	c := make(chan int)
	done := make(chan bool)

	for i := 0; i < n; i++ {
		go func() {
			for i := 0; i < 10; i++ {
				c <- i
			}
			done <- true
		}()
	}

	go func() {
		for i := 0; i < n; i++ {
			<-done
		}
		close(c)
	}()

	for i := range c {
		fmt.Println(i)
	}
	fmt.Println("\n|-------End of channels pacakge-------------|\n")
}

/*ExecOneToNExample is the main function of channels package*/
func ExecOneToNExample() {
	fmt.Println("|-------Beginning of channels package-------|\n")
	c := make(chan int)
	done := make(chan bool)

	go func() {
		for i := 0; i < 100000; i++ {
			c <- i
		}
		close(c)
	}()

	go func() {
		for i := range c {
			fmt.Println(i)
		}
		done <- true
	}()

	go func() {
		for i := range c {
			fmt.Println(i)
		}
		done <- true
	}()

	<-done
	<-done
	fmt.Println("\n|-------End of channels pacakge-------------|\n")
}

/*ExecMultipleOneToNExample is the main function of channels package*/
func ExecMultipleOneToNExample() {
	fmt.Println("|-------Beginning of channels package-------|\n")
	n := 10
	c := make(chan int)
	done := make(chan bool)

	go func() {
		for i := 0; i < 100000; i++ {
			c <- i
		}
		close(c)
	}()

	for i := 0; i < n; i++ {
		go func() {
			for i := range c {
				fmt.Println(i)
			}
			done <- true
		}()
	}

	for i := 0; i < n; i++ {
		<-done
	}
	fmt.Println("\n|-------End of channels pacakge-------------|\n")
}

/*ExecPassingExample is the main function of channels package*/
func ExecPassingExample() {
	fmt.Println("|-------Beginning of channels package-------|\n")
	c := incrementor()
	cSum := puller(c)
	for i := range cSum {
		fmt.Println(i)
	}
	fmt.Println("\n|-------End of channels pacakge-------------|\n")
}

func incrementor() chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			out <- i
		}
		close(out)
	}()
	return out
}

func puller(c chan int) chan int {
	out := make(chan int)
	go func() {
		var sum int
		for i := range c {
			sum += i
		}
		out <- sum
		close(out)
	}()
	return out
}
