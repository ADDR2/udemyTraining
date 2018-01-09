package tcpServer

import "fmt"
import "net"
import "log"
import "io"

/*ExecTcpServerExample is the main function of tcpServer package*/
func ExecTcpServerExample() {
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
