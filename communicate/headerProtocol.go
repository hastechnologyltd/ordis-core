package communicate

import (
	"bufio"
	"bytes"
	"encoding/binary"
)

type HeaderProtocol struct {
	CorrelationId []byte
	Who           []byte
	Created       [5]byte
	What          []byte
	How           []byte
}

func (headerProtocol *HeaderProtocol) ReadFrom(rw *bufio.ReadWriter) error {
	var err error

	if headerProtocol.CorrelationId, err = rw.ReadBytes(0); err != nil {
		return err
	}

	if headerProtocol.Who, err = rw.ReadBytes(0); err != nil {
		return err
	}

	createdBytes := make([]byte, 5)
	if _, err := rw.Read(createdBytes); err != nil {
		return err
	}
	createdReader := bytes.NewReader(createdBytes)
	if err = binary.Read(createdReader, binary.LittleEndian, &headerProtocol.Created); err != nil {
		return err
	}

	if headerProtocol.What, err = rw.ReadBytes(0); err != nil {
		return err
	}

	howSize, err := rw.ReadByte()

	if _, err := rw.Read(headerProtocol.How); err == nil {
		howBytes := make([]byte, howSize)
		if _, err := rw.Read(howBytes); err != nil {
			return err
		}
		headerProtocol.How = howBytes
	}

	return err
}
