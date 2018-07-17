package rudp

import (
	"encoding/binary"
	"math"
)

type Message struct {
	SequenceNumber      uint32
	IsFragmented        bool
	IsAck               bool
	IsLastFragment      bool
	FragmentationOffset uint64
	FragmentationId     uint64
	Data                []byte
}

const SEQUENCE_NUMBER_START_INDEX = 0
const SEQUENCE_NUMBER_LENGHT = 4

const FLAGS_BYTE_INDEX = 4

const IS_FRAGMENTED_BIT_FLAG = 7
const IS_ACK_BIT_FLAG = 6
const IS_LAST_FRAGMENT_BIT_FLAG = 5


const FRAGMENT_OFFSET_START_INDEX = 5
const FRAGMENT_OFFSET_LENGHT = 8

const FRAGMENTATION_ID_START_INDEX = 13
const FRAGMENTATION_ID_LENGHT = 8

const PAYLOAD_START_INDEX = 21


func ParseDatagramMessage(datagram []byte) Message {

	message := Message{
		SequenceNumber:      binary.BigEndian.Uint32(datagram[SEQUENCE_NUMBER_START_INDEX: SEQUENCE_NUMBER_LENGHT]),
		IsFragmented:        datagram[FLAGS_BYTE_INDEX] & byte(math.Pow(2, IS_FRAGMENTED_BIT_FLAG)) == byte(math.Pow(2, IS_FRAGMENTED_BIT_FLAG)),
		IsAck:               datagram[FLAGS_BYTE_INDEX] & byte(math.Pow(2, IS_ACK_BIT_FLAG)) == byte(math.Pow(2, IS_ACK_BIT_FLAG)),
		IsLastFragment:      datagram[FLAGS_BYTE_INDEX] & byte(math.Pow(2, IS_LAST_FRAGMENT_BIT_FLAG)) == byte(math.Pow(2, IS_LAST_FRAGMENT_BIT_FLAG)),
		FragmentationOffset: binary.BigEndian.Uint64(datagram[FRAGMENT_OFFSET_START_INDEX: FRAGMENT_OFFSET_START_INDEX + FRAGMENT_OFFSET_LENGHT]),
		FragmentationId:     binary.BigEndian.Uint64(datagram[FRAGMENTATION_ID_START_INDEX: FRAGMENTATION_ID_START_INDEX + FRAGMENTATION_ID_LENGHT]),
		Data:                datagram[PAYLOAD_START_INDEX:],
	}

	return message
}

func SerializeMessage(message Message) []byte {

	datagram := make([]byte, 21 + len(message.Data))

	// Making sequence number
	for i := SEQUENCE_NUMBER_START_INDEX; i < SEQUENCE_NUMBER_START_INDEX + SEQUENCE_NUMBER_LENGHT; i++ {
		datagram[i] = byte(message.SequenceNumber >> uint((SEQUENCE_NUMBER_LENGHT - i - 1) * 8))
	}

	// Making flags
	datagram[4] = 0

	if message.IsFragmented {
		datagram[4] += byte(math.Pow(2, IS_LAST_FRAGMENT_BIT_FLAG))
	}

	if message.IsAck {
		datagram[4] += byte(math.Pow(2, IS_ACK_BIT_FLAG))
	}

	if message.IsLastFragment {
		datagram[4] += byte(math.Pow(2, IS_LAST_FRAGMENT_BIT_FLAG))
	}

	// Making offset
	for i := FRAGMENT_OFFSET_START_INDEX; i < FRAGMENT_OFFSET_START_INDEX + FRAGMENT_OFFSET_LENGHT; i++ {
		datagram[i] = byte(message.FragmentationOffset >> uint((FRAGMENT_OFFSET_START_INDEX + FRAGMENT_OFFSET_LENGHT - i - 1) * 8))
	}

	// Making fragmentation id
	for i := FRAGMENTATION_ID_START_INDEX; i < FRAGMENTATION_ID_START_INDEX + FRAGMENTATION_ID_LENGHT; i++ {
		datagram[i] = byte(message.FragmentationId >> uint((FRAGMENTATION_ID_START_INDEX + FRAGMENTATION_ID_LENGHT - i - 1) * 8))
	}

	for i := 0; i < len(message.Data); i++ {
		datagram[21 + i] = message.Data[i]
	}

	return datagram
}