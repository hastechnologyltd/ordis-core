package communicate

import (
	"github.com/gansidui/gotcp"
	"github.com/gansidui/gotcp/examples/echo"
	"log"
	"net"
	"strconv"
)

func server(addressPort int, protocol gotcp.Protocol) {

	tcpAddr, err := net.ResolveTCPAddr("tcp4", ":"+strconv.Itoa(addressPort))
	checkError(err)
	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	config := &gotcp.Config{
		PacketSendChanLimit:    20,
		PacketReceiveChanLimit: 20,
	}
	srv := gotcp.NewServer(config, &Callback{}, protocol)
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
