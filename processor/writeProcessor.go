package processor

import "time"

type WriteSchema struct {
	CorrelationId string
	Who           string
	Created       time.Time
	What          string
	How           []byte
}

func (schema *WriteSchema) Process() {

}
