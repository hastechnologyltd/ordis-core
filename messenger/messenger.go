package messenger

import (
	"container/list"
	"fmt"
	"time"
)

type Message struct {
	Val string
}

var queue *list.List
var messages chan Message

func NewMessenger() {
	messages = make(chan Message)
	queue = list.New()
}

func Messenger() {
	for {
		if queue.Len() > 0 {
			e := queue.Front()
			queue.Remove(e)
			messages <- e.Value.(Message)
		}
	}
}

func SendMessage(m Message) {
	queue.PushBack(m)
}

func RetrieveMessage() {
	for {
		msg := <-messages
		fmt.Println(msg.Val)
		time.Sleep(60)
	}
}
