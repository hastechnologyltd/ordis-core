package protocol

import "bufio"

type MessageProtocol interface {
	ReadMessage(rw *bufio.ReadWriter) error
}
