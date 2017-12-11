package errors

import (
	"fmt"
	"log"
	"os"
)

type MathError struct {
	lat, long string
	err       error
}

func (n *MathError) Error() string {
	return fmt.Sprintf("Math error: %v %v %v", n.lat, n.long, n.err)
}

func init() {
	newFile, err := os.Create("log.txt")
	if err != nil {
		fmt.Println(err)
	}
	log.SetOutput(newFile)
}

/*ExecErrorsExample is the main function of errors package*/
func ExecErrorsExample() {
	fmt.Println("|-------Beginning of errors package-------|\n")
	_, err := sqrt(-10)
	if err != nil {
		//log.Println("Error happened", err)
		log.Fatalln(err)
		//log.Panic(err)
		//panic(err)
	}
	fmt.Println("\n|-------End of errors pacakge-------------|\n")
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		mathError := fmt.Errorf("square root of negative number: %v", f)
		return 0, &MathError{"50.2289 N", "99.4656 W", mathError}
	}
	return 42, nil
}
