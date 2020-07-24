package fsm

import (
	"errors"
	"fmt"
	"sync"
)

type State string
type CallbackFunc func(*Event)

type Transition struct {
	Event string
	From  State
	To    State
}

type transitionKey struct {
	event string
	from  State
}

type FSM struct {
	// state is the current state of the FSM
	state         State
	callbackTable map[State]CallbackFunc

	transitions map[transitionKey]Transition
	stateMutex  sync.RWMutex
	eventMutex  sync.Mutex
}

func NewFSM(initState State, transitions []Transition, callbacks map[State]CallbackFunc) (*FSM, error) {
	fsm := new(FSM)
	fsm.transitions = make(map[transitionKey]Transition)
	fsm.callbackTable = make(map[State]CallbackFunc)

	fsm.state = initState

	states := make(map[State]bool)
	events := make(map[string]bool)

	for _, transition := range transitions {
		key := transitionKey{
			event: transition.Event,
			from:  transition.From,
		}
		if _, ok := fsm.transitions[key]; ok {
			return nil, errors.New("Transitions has duplicate key(same event and same source state)")
		} else {
			fsm.transitions[key] = transition
			states[transition.From] = true
			states[transition.To] = true
			events[transition.Event] = true
		}
	}

	for k, v := range callbacks {
		fsm.callbackTable[k] = v
	}

	return fsm, nil
}

// Current get the current state of fsm
func (fsm *FSM) Current() State {
	fsm.stateMutex.RLock()
	defer fsm.stateMutex.RUnlock()
	return fsm.state
}

// IsCurrent return true if the current state of fsm is equal to state
func (fsm *FSM) IsCurrent(state State) bool {
	fsm.stateMutex.RLock()
	defer fsm.stateMutex.RUnlock()
	return fsm.state == state
}

func (fsm *FSM) SendEvent(event *Event) error {
	current := fsm.Current()

	if callback, ok := fsm.callbackTable[current]; ok {
		callback(event)
		return nil
	} else {
		return fmt.Errorf("State %s does not exists", current)
	}
}

/* args is for ENTRY params*/
func (fsm *FSM) Transition(next State, args map[string]interface{}) error {
	if _, ok := fsm.callbackTable[next]; !ok {
		return fmt.Errorf("State %s does not exists", next)
	}

	event := &Event{
		Name: EntryEvent,
		Args: make(map[string]interface{}),
	}

	for k, v := range args {
		event.Args[k] = v
	}

	if !fsm.IsCurrent(next) {
		return fmt.Errorf("No transition")
	} else {
		return fsm.SendEvent(event)
	}
}

func (fsm *FSM) SetState(state State) {
	fsm.stateMutex.Lock()
	defer fsm.stateMutex.Unlock()
	fsm.state = state
}
