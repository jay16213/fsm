package fsm

type Event struct {
	Name string
	Args map[string]interface{}
}

const (
	EntryEvent string = "Entry event"
)
