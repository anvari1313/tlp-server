package main

import (
	"encoding/hex"
	"fmt"
	"net"
)

func main() {
	fmt.Println("Hello World!")

	//Basic variables
	port := ":8080"
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

	//Keep calling this function
	for {
		display(udpConn)
	}

}

func display(conn *net.UDPConn) {

	buffer := make([]byte, 1024)
	n, addr, err := conn.ReadFromUDP(buffer)
	fmt.Println(addr)
	if err != nil {
		fmt.Println("Error Reading")
		return
	} else {
		fmt.Println(hex.EncodeToString(buffer[0:n]))
		fmt.Println(buffer[0:n])
		fmt.Println("Package Done")
	}

}