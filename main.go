package main

import (
	"github.com/hastechnologyltd/ordis-core/communicate"
	"github.com/pkg/errors"
	"log"
)

func main() {
	//connect := flag.String("connect", "", "IP address of process to join. If empty, go into listen mode.")
	//flag.Parse()
	//if *connect != "" {
	//	err := client(*connect)
	//	if err != nil {
	//		log.Println("Error:", errors.WithStack(err))
	//	}
	//	log.Println("Client done.")
	//	return
	//}
	server := communicate.NewServer(48202)

	err := server.Listen()
	if err != nil {
		log.Println("Error:", errors.WithStack(err))
	}

	log.Println("Server done.")
}

func init() {
	log.SetFlags(log.Lshortfile)
}
