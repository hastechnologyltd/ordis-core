package communicate

import (
	"bufio"
	"time"
)

type Header struct {
	CorrelationId string
	Who           string
	Created       time.Time
	What          string
}

func NewHeader(rw *bufio.ReadWriter) Header {
	return Header{}
}
