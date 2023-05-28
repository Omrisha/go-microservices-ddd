package es

type StreamID struct {
	value string
}

func NewStreamID(from string) StreamID {
	if from == "" {
		panic("newStreamID: empty input value")
	}

	return StreamID{value: from}
}

func (streamId StreamID) String() string {
	return streamId.value
}
