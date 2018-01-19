package tcpServer

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"time"
)

/*ExecTCPServerExample is the main function of tcpServer package*/
func ExecTCPServerExample() {
	fmt.Println("|-------Beginning of tcpServer package-------|\n")
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Panic(err)
	}

	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
		}

		io.WriteString(conn, "\nHello from TCP server\n")
		fmt.Fprintln(conn, "How is your day?")
		fmt.Fprintf(conn, "%v", "Well, I hope!")

		conn.Close()
	}
	fmt.Println("\n|-------End of tcpServer pacakge-------------|\n")
}

/*ExecTCPReaderExample*/
func ExecTCPReaderExample(ack bool) {
	fmt.Println("|-------Beginning of tcpServer package-------|\n")
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Panic(err)
	}

	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
		}

		if ack {
			go handleWithResponse(conn)
		} else {
			go handle(conn)
		}
	}
	fmt.Println("\n|-------End of tcpServer pacakge-------------|\n")
}

func handle(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
	}

	defer conn.Close()
	fmt.Println("Code got here.")
}

func handleWithResponse(conn net.Conn) {
	err := conn.SetDeadline(time.Now().Add(10 * time.Second))
	if err != nil {
		log.Println("CONN TIMEOUT")
	}
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		fmt.Fprintf(conn, "I heard you say: %s\n", ln)
	}

	defer conn.Close()
	fmt.Println("Code got here.")
}

/*ExecTCPDialExample*/
func ExecTCPDialExample() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	bs, err := ioutil.ReadAll(conn)
	if err != nil {
		log.Println(err)
	}

	fmt.Println(string(bs))
}

/*ExecTCPWriteDialExample*/
func ExecTCPWriteDialExample() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	fmt.Fprintln(conn, "I dialed you.")
}
