package communicate

import "net"

type Packet interface {
	Serialise() []byte
}

type Protocol interface {
	ReadPacket(conn *net.TCPConn) (Packet, error)
	WritePacket()
}
