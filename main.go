// Copyright 2019 Has Technology Ltd.
// Use of this source code is governed by a GNU GPLv3
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"github.com/hastechnologyltd/ordis-core/messenger"
	"log"
	"time"
)

func main() {
	//server := communicate.NewServer(48202)
	//
	//err := server.Listen()
	//if err != nil {
	//	log.Println("Error:", errors.WithStack(err))
	//}
	//
	//log.Println("Server done.")

	messenger.NewMessenger()
	go messenger.RetrieveMessage()

	messenger.SendMessage(messenger.Message{"Hello"})
	messenger.SendMessage(messenger.Message{"Jeff"})
	messenger.SendMessage(messenger.Message{"Here is"})
	messenger.SendMessage(messenger.Message{"a message"})

	go messenger.Messenger()

	for {
		fmt.Println("Waiting...")
		time.Sleep(1000)
	}

}

func init() {
	log.SetFlags(log.Lshortfile)
}
