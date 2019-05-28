package processor

type Schemas struct {
	id     byte
	schema Schema
	access byte
}

func Factory() []Schemas {
	return []Schemas{
		{0xf0, &AdminSchema{}, 4},
		{0x20, &ReadSchema{}, 1},
		{0x10, &WriteSchema{}, 2},
	}
}
