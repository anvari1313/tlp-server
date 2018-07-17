package rudp

import (
	"net"
	"fmt"
	"encoding/hex"
)

func StartRUDPServer(port string) {
	//Basic variables
	protocol := "udp"

	//Build the address
	udpAddr, err := net.ResolveUDPAddr(protocol, port)
	if err != nil {
		fmt.Println("Wrong Address")
		return
	}

	//Create the connection
	udpConn, err := net.ListenUDP(protocol, udpAddr)
	if err != nil {
		fmt.Println(err)
	}

	for {
		display(udpConn)
	}
}

func display(conn *net.UDPConn) {

	buffer := make([]byte, 1024)
	n, addr, err := conn.ReadFromUDP(buffer)
	conn.WriteTo(buffer[0:n], addr)
	fmt.Println(addr)
	if err != nil {
		fmt.Println("Error Reading")
		return
	} else {
		fmt.Print(n, "  ")
		fmt.Println(hex.EncodeToString(buffer[0:n]))
		fmt.Println(buffer[0:n])
		fmt.Println("Package Done")
	}

}