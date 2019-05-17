package audit

import (
	"log"
	"os"
	"strconv"
)

func (audits *Audits) Backup(datastoreName string) {

	currentNode := audits.tail
	data := strconv.FormatUint(currentNode.id, 10) + " - " + currentNode.data + "\n"

	f, err := os.OpenFile(datastoreName,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()
	if _, err := f.WriteString(data); err != nil {
		log.Println(err)
	}
}
