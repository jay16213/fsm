package fsm

import (
	"fmt"
	"sync"
)

type State string
type Event string
type Args map[string]interface{}

type CallbackFunc func(*FSM, Event, Args)
type Callbacks map[State]CallbackFunc

const (
	EntryEvent Event = "Entry event"
)

type FSM struct {
	// state is the current state of the FSM
	state State

	// callbackTable store one callback function for one state
	callbackTable map[State]CallbackFunc

	// stateMutex ensure that all operation to state is thread-safe
	stateMutex sync.RWMutex
}

func NewFSM(initState State, callbacks Callbacks) (*FSM, error) {
	fsm := new(FSM)

	fsm.callbackTable = make(map[State]CallbackFunc)
	fsm.state = initState

	for state, callback := range callbacks {
		fsm.callbackTable[state] = callback
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

// SendEvent triggers a callback for current state with event and optional args
func (fsm *FSM) SendEvent(event Event, args Args) error {
	current := fsm.Current()

	if callback, ok := fsm.callbackTable[current]; ok {
		callback(fsm, event, args)
		return nil
	} else {
		return fmt.Errorf("State %s does not exists", current)
	}
}

// Transition will transfer fsm current state to next, then trigger a EntryEvent callback for next state
func (fsm *FSM) Transition(next State, args Args) error {
	if _, ok := fsm.callbackTable[next]; !ok {
		return fmt.Errorf("State %s does not exists", next)
	}

	fsm.SetState(next)
	return fsm.SendEvent(EntryEvent, args)
}

// SetState set fsm current state to next
func (fsm *FSM) SetState(next State) {
	fsm.stateMutex.Lock()
	defer fsm.stateMutex.Unlock()
	fsm.state = next
}
