package audit

type Element struct {
	id   uint64
	data string
	next *Element
}

type Audits struct {
	head *Element
	tail *Element
}
