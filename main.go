package main

import (
	"github.com/anvari1313/tlp-server/rudp"
	"fmt"
)



func main() {
	t := []byte{1, 2, 29, 22}
	//me := rudp.Message{1256, true, true, false,5612, 5, t}
	//fmt.Println(me)
	//serailized := rudp.SerializeMessage(me)
	//newMessage := rudp.ParseDatagramMessage(serailized)
	//fmt.Println(newMessage)

	m := map[uint64][]rudp.Message{}
	if m[12] == nil {
		fmt.Println("Hash is null")
		buffer := make([]rudp.Message, 0)
		m[12] = buffer
		s := m[12]
		message := rudp.Message{
			1256,
			true,
			true,
			false,
			5612,
			5,
			t}
		a := append(s, message)
		m[12] = a
	}

	if m[12] == nil {
		fmt.Println("This is also null")
	} else {
		fmt.Println(m[12])
	}
	port := ":8080"
	rudp.StartRUDPServer(port)
}