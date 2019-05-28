package communicate

import (
	"bufio"
	"github.com/hastechnologyltd/ordis-core/communicate/protocol"
	"github.com/hastechnologyltd/ordis-core/security"
	"github.com/pkg/errors"
	"io"
	"log"
	"net"
	"strconv"
	"sync"
)

//  great tutorial https://appliedgo.net/networking/

type server struct {
	endpoint *Endpoint
}

const (
	AdminCommand = 0xf0
	ReadCommand  = 0x10
	WriteCommand = 0x20
)

func NewServer(port int) server {
	server := server{
		endpoint: NewEndpoint(port),
	}

	server.endpoint.AddHandleFunc(AdminCommand, handleAdminCommand)
	server.endpoint.AddHandleFunc(ReadCommand, handleReadCommand)
	server.endpoint.AddHandleFunc(WriteCommand, handleWriteCommand)
	return server
}

func (s *server) Listen() error {
	var err error
	s.endpoint.listener, err = net.Listen("tcp", s.endpoint.portAddress)
	if err != nil {
		return errors.Wrapf(err, "Unable to listen on port %s\n", s.endpoint.portAddress)
	}
	log.Println("Listen on", s.endpoint.listener.Addr().String())
	for {
		log.Println("Accept a connection request.")
		conn, err := s.endpoint.listener.Accept()
		if err != nil {
			log.Println("Failed accepting a connection request:", err)
			continue
		}
		log.Println("Handle incoming messages.")
		go s.endpoint.handleMessages(conn)
	}
}

type HandleFunc func(*bufio.ReadWriter)

type Endpoint struct {
	listener    net.Listener
	handler     map[byte]HandleFunc
	m           sync.RWMutex
	portAddress string
}

func NewEndpoint(port int) *Endpoint {
	return &Endpoint{
		handler:     map[byte]HandleFunc{},
		portAddress: ":" + strconv.Itoa(port),
	}
}

func (e *Endpoint) AddHandleFunc(command byte, f HandleFunc) {
	e.m.Lock()
	e.handler[command] = f
	e.m.Unlock()
}

func (e *Endpoint) handleMessages(conn net.Conn) {
	rw := bufio.NewReadWriter(bufio.NewReader(conn), bufio.NewWriter(conn))
	defer conn.Close()

	for {
		log.Print("Receive command '")

		command, err := rw.ReadByte()
		log.Println(command)
		log.Println(command & 240)
		log.Println(command & 15)
		username, err := rw.ReadBytes(0)
		log.Println(string(username[:len(username)-1]))
		password, err := rw.ReadBytes(0)
		log.Println(string(password[:len(password)-1]))

		credentials := security.NewCredentials(username, password)
		isAuthenticated := credentials.Authenticate()
		if !isAuthenticated {
			errors.New("Invalid Credentials")
			return
		}

		writeMessageProtocol := protocol.WriteMessageProtocol{}
		err = writeMessageProtocol.ReadMessage(rw)

		additionalData, _, err := rw.ReadLine()
		log.Println(additionalData)
		//errors.New("Invalid protocol given")

		//cmd, err := rw.ReadString('\n')
		switch {
		case err == io.EOF:
			log.Println("Reached EOF - close this connection.\n   ---")
			return
		case err != nil:
			log.Println("\nError reading command. Got: '"+string(command)+"'\n", err)
			return
		}

		e.m.RLock()
		handleCommand, ok := e.handler[command]
		e.m.RUnlock()
		if !ok {
			log.Println("Command '" + string(command) + "' is not registered.")
			return
		}
		handleCommand(rw)
	}
}

func handleAdminCommand(rw *bufio.ReadWriter) {

}

func handleReadCommand(rw *bufio.ReadWriter) {

}

func handleWriteCommand(rw *bufio.ReadWriter) {

}

//type complexData struct {
//	N int
//	S string
//	M map[string]int
//	P []byte
//	C *complexData
//}

//func funcName(rw *bufio.ReadWriter) {
//	dataHeaderBytes := make([]byte, 3)
//	rw.Read(dataHeaderBytes)
//	dataHeaderReader := bytes.NewReader(dataHeaderBytes)
//	var dataHeader struct {
//		CorrelationSize byte
//		WhoSize         byte
//		WhatSize        byte
//	}
//	if err := binary.Read(dataHeaderReader, binary.LittleEndian, &dataHeader); err != nil {
//		fmt.Println("binary.Read failed:", err)
//	}
//}

//func handleStrings(rw *bufio.ReadWriter) {
//	log.Print("Receive STRING message:")
//	s, err := rw.ReadString('\n')
//	if err != nil {
//		log.Println("Cannot read from connection.\n", err)
//	}
//	s = strings.Trim(s, "\n ")
//	log.Println(s)
//	_, err = rw.WriteString("Thank you.\n")
//	if err != nil {
//		log.Println("Cannot write to connection.\n", err)
//	}
//	err = rw.Flush()
//	if err != nil {
//		log.Println("Flush failed.", err)
//	}
//}

//func handleGob(rw *bufio.ReadWriter) {
//	log.Print("Receive GOB data:")
//	var data complexData
//	dec := gob.NewDecoder(rw)
//	err := dec.Decode(&data)
//	if err != nil {
//		log.Println("Error decoding GOB data:", err)
//		return
//	}
//	log.Printf("Outer complexData struct: \n%#v\n", data)
//	log.Printf("Inner complexData struct: \n%#v\n", data.C)
//}
