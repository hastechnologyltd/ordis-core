package communicate

//import (
//	"bufio"
//	"encoding/gob"
//	"github.com/pkg/errors"
//	"log"
//	"net"
//	"strconv"
//)

//func client(ip string, port string) error {
//	testStruct := complexData{
//		N: 23,
//		S: "string data",
//		M: map[string]int{"one": 1, "two": 2, "three": 3},
//		P: []byte("abc"),
//		C: &complexData{
//			N: 256,
//			S: "Recursive structs? Piece of cake!",
//			M: map[string]int{"01": 1, "10": 2, "11": 3},
//		},
//	}
//	rw, err := Open(ip + port)
//	if err != nil {
//		return errors.Wrap(err, "Client: Failed to open connection to "+ip+port)
//	}
//	log.Println("Send the string request.")
//	n, err := rw.WriteString("STRING\n")
//	if err != nil {
//		return errors.Wrap(err, "Could not send the STRING request ("+strconv.Itoa(n)+" bytes written)")
//	}
//	n, err = rw.WriteString("Additional data.\n")
//	if err != nil {
//		return errors.Wrap(err, "Could not send additional STRING data ("+strconv.Itoa(n)+" bytes written)")
//	}
//	log.Println("Flush the buffer.")
//	err = rw.Flush()
//	if err != nil {
//		return errors.Wrap(err, "Flush failed.")
//	}
//	log.Println("Read the reply.")
//	response, err := rw.ReadString('\n')
//	if err != nil {
//		return errors.Wrap(err, "Client: Failed to read the reply: '"+response+"'")
//	}
//
//	log.Println("STRING request: got a response:", response)
//	log.Println("Send a struct as GOB:")
//	log.Printf("Outer complexData struct: \n%#v\n", testStruct)
//	log.Printf("Inner complexData struct: \n%#v\n", testStruct.C)
//	enc := gob.NewEncoder(rw)
//	n, err = rw.WriteString("GOB\n")
//	if err != nil {
//		return errors.Wrap(err, "Could not write GOB data ("+strconv.Itoa(n)+" bytes written)")
//	}
//	err = enc.Encode(testStruct)
//	if err != nil {
//		return errors.Wrapf(err, "Encode failed for struct: %#v", testStruct)
//	}
//	err = rw.Flush()
//	if err != nil {
//		return errors.Wrap(err, "Flush failed.")
//	}
//	return nil
//}
//
//func Open(addr string) (*bufio.ReadWriter, error) {
//	log.Println("Dial " + addr)
//	conn, err := net.Dial("tcp", addr)
//	if err != nil {
//		return nil, errors.Wrap(err, "Dialing "+addr+" failed")
//	}
//	return bufio.NewReadWriter(bufio.NewReader(conn), bufio.NewWriter(conn)), nil
//}
