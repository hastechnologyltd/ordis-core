package main

import (
	"fmt"
	"log"
	"net"
	"strconv"
)

//var audits *audit.Audits
//
//const (
//	ConnHost = "localhost"
//	ConnPort = "28028"
//	ConnType = "tcp"
//)
//
func handleRequest(conn *net.TCPConn) {
	// Make a buffer to hold incoming data.
	buf := make([]byte, 1024)
	// Read the incoming connection into the buffer.
	_, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Error reading:", err.Error())
	}
	fmt.Println(buf)
	// Send a response back to person contacting us.
	conn.Write([]byte("Message received."))
	// Close the connection when you're done with it.
	conn.Close()
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func newServer(addressPort int) {

	tcpAddr, err := net.ResolveTCPAddr("tcp4", ":"+strconv.Itoa(addressPort))
	checkError(err)
	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	go listen(listener)
}

func listen(listener *net.TCPListener) {
	defer listener.Close()
	for {
		tcpConn, _ := listener.AcceptTCP()
		handleRequest(tcpConn)
	}
}

func main() {

	newServer(48201)

	newServer(48202)

	newServer(48203)

	select {}

	//l, err := net.Listen(ConnType, ConnHost+":"+ConnPort)
	//if err != nil {
	//	fmt.Println("Error listening:", err.Error())
	//	os.Exit(1)
	//}
	//// Close the listener when the application closes.
	//defer l.Close()
	//fmt.Println("Listening on " + ConnHost + ":" + ConnPort)
	//for {
	//	// Listen for an incoming connection.
	//	conn, err := l.Accept()
	//	if err != nil {
	//		fmt.Println("Error accepting: ", err.Error())
	//		os.Exit(1)
	//	}
	//	// Handle connections in a new goroutine.
	//	go handleRequest(conn)
	//}
	//
	//audits = audit.CreateAudit()
	//
	//router := mux.NewRouter().StrictSlash(true)
	//
	//router.HandleFunc("/", Index)
	//router.HandleFunc("/audit/{auditData}", AddAudit)
	//
	//log.Fatal(http.ListenAndServe(":8080", router))
}

//func Index(w http.ResponseWriter, r *http.Request) {
//	fmt.Fprintln(w, "Welcome!")
//}
//
//func AddAudit(w http.ResponseWriter, r *http.Request) {
//	vars := mux.Vars(r)
//	auditData := vars["auditData"]
//	audits.AddAudit(auditData)
//	audits.Backup("data.txt")
//	audits.Display()
//}
