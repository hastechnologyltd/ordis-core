package protocol

import (
	"bufio"
	"bytes"
	"encoding/binary"
)

type WriteMessageProtocol struct {
	CorrelationId []byte
	Who           []byte
	Created       [5]byte
	What          []byte
	How           []byte
}

func (messageProtocol *WriteMessageProtocol) ReadMessage(rw *bufio.ReadWriter) error {
	var err error

	if messageProtocol.CorrelationId, err = rw.ReadBytes(0); err != nil {
		return err
	}

	if messageProtocol.Who, err = rw.ReadBytes(0); err != nil {
		return err
	}

	createdBytes := make([]byte, 5)
	if _, err := rw.Read(createdBytes); err != nil {
		return err
	}
	createdReader := bytes.NewReader(createdBytes)
	if err = binary.Read(createdReader, binary.LittleEndian, &messageProtocol.Created); err != nil {
		return err
	}

	if messageProtocol.What, err = rw.ReadBytes(0); err != nil {
		return err
	}

	howSize, _ := rw.ReadByte()
	howBytes := make([]byte, howSize)
	if _, err := rw.Read(howBytes); err != nil {
		return err
	}
	messageProtocol.How = howBytes

	return err
}
