package processor

type Schema interface {
	Process()
}

// messenger -> mapper -> processor
