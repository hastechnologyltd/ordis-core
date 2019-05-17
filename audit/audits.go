package audit

import (
	"fmt"
	"github.com/sony/sonyflake"
)

var sonyFlake *sonyflake.Sonyflake

func CreateAudit() *Audits {
	var st sonyflake.Settings
	sonyFlake = sonyflake.NewSonyflake(st)
	return &Audits{}
}

func (audits *Audits) AddAudit(data string) {
	var id, _ = sonyFlake.NextID()
	element := &Element{
		id:   id,
		data: data,
	}
	if audits.head == nil {
		audits.head = element
		audits.tail = element
	} else {
		currentNode := audits.tail
		currentNode.next = element
		audits.tail = element
	}

	fmt.Printf("%+v\n", audits.tail)
}

func (audits *Audits) Display() {
	currentNode := audits.head
	if currentNode == nil {
		fmt.Println("Playlist is empty.")
	}
	fmt.Printf("%+v\n", *currentNode)
	for currentNode.next != nil {
		currentNode = currentNode.next
		fmt.Printf("%+v\n", *currentNode)
	}
}
