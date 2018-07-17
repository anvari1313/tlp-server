package dnstcp

import (
	"net"
	"fmt"
	"time"
	"log"
)

var protocol = "tcp"

func StartDNSTCPServer(port string) {
	// making the address
	tcpAddress, err := net.ResolveTCPAddr(protocol, port)

	if err != nil {
		fmt.Println("TCP failed on Address")
		return
	}
	log.Println("TCP address is OK")
	// opening the connection
	tcpListener, err := net.ListenTCP(protocol, tcpAddress)
	if err != nil {
		fmt.Println("TCP failed on opening connection:")
		fmt.Println(err)
		return
	}
	log.Println("TCP connection is Open")

	for {
		tcpConnection, err := tcpListener.Accept()
		if err != nil {
			// error causes rejecting the connection, so we ignore it
			continue
		}
		log.Println("New TCP connection has accepted")
		go handleClient(tcpConnection)
	}

}

func handleClient(conn net.Conn) {
	// set read deadline for 1 minute
	conn.SetReadDeadline(time.Now().Add(1 * time.Minute))
	buffer := make([]byte, 2048)

	for {
		readLen, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("TCP read message failed:")
			fmt.Println(err)
			break;
		}

		if readLen == 0 {
			// connection closed
			break
		} else {
			fmt.Println(string(buffer[:readLen]))
		}
	}
}