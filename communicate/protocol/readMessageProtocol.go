package protocol

import "bufio"

type ReadMessageProtocol struct {
	CorrelationId []byte
}

func (messageProtocol *ReadMessageProtocol) ReadMessage(rw *bufio.ReadWriter) error {
	var err error

	messageProtocol.CorrelationId, err = rw.ReadBytes(0)

	return err
}
