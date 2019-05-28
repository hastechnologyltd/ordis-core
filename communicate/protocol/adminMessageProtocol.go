package protocol

import "bufio"

type AdminMessageProtocol struct {
	CorrelationId []byte
}

func (messageProtocol *AdminMessageProtocol) ReadMessage(rw *bufio.ReadWriter) error {
	var err error

	messageProtocol.CorrelationId, err = rw.ReadBytes(0)

	return err
}
