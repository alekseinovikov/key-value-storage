package log

type TransactionLogger interface {
	WriteDelete(key string)
	WritePut(key, value string)
}

type Event struct {
	Sequence  uint64
	EventType EventType
	Key       string
	Value     string
}

type EventType byte

const (
	_           = iota
	EventDelete = iota
	EventPut    = iota
)
