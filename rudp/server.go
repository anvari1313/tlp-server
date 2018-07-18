package rudp

import (
	"net"
	"fmt"
)

var messageFragmentationMap map[uint64] []Message


func StartRUDPServer(port string) {
	fmt.Println("Server is started")
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
	//conn.WriteTo(buffer[0:n], addr)
	fmt.Println(addr)
	if err != nil {
		fmt.Println("Error Reading")
		return
	} else {
		m := ParseDatagramMessage(buffer[0:n])
		ack := Message{SequenceNumber:m.SequenceNumber, IsAck:true, IsFragmented:false, IsLastFragment:false, FragmentationId:0, FragmentationOffset:0}
		conn.WriteTo(SerializeMessage(ack), addr)
		receiveMessage(m)
		//fmt.Println(buffer[0:n])
		fmt.Println("Package Done")
	}

}

func receivedData(data []byte)  {
	fmt.Print("Received Data : ")
	fmt.Println(data)
}

func receiveMessage(message Message)  {
	if messageFragmentationMap == nil {
		messageFragmentationMap = make(map[uint64] []Message)
	}

	if message.IsFragmented == false {
		receivedData(message.Data)
	} else {
		if messageFragmentationMap[message.FragmentationId] == nil {
			buffer := make([]Message, 0)
			buffer = append(buffer, message)
			messageFragmentationMap[message.FragmentationId] = buffer
		} else {
			me := messageFragmentationMap[message.FragmentationId]
			me = append(me, message)
			messageFragmentationMap[message.FragmentationId] = me
			if message.IsLastFragment == true {
				cumulators := make([]byte, 0)
				for i := 0; i < len(me); i++ {
					fmt.Printf("Message[%d].Data = ", i)
					fmt.Println(me[i].Data)
					cumulators = append(cumulators, me[i].Data ...)
				}
				receivedData(cumulators)
			}
		}
	}

}