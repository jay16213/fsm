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
	state         State
	callbackTable map[State]CallbackFunc

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

func (fsm *FSM) SendEvent(event Event, args Args) error {
	current := fsm.Current()

	if callback, ok := fsm.callbackTable[current]; ok {
		callback(fsm, event, args)
		return nil
	} else {
		return fmt.Errorf("State %s does not exists", current)
	}
}

/* args is for ENTRY params*/
func (fsm *FSM) Transition(next State, args Args) error {
	if _, ok := fsm.callbackTable[next]; !ok {
		return fmt.Errorf("State %s does not exists", next)
	}

	if fsm.IsCurrent(next) {
		return fmt.Errorf("No transition")
	} else {
		fsm.SetState(next)
		return fsm.SendEvent(EntryEvent, args)
	}
}

func (fsm *FSM) SetState(state State) {
	fsm.stateMutex.Lock()
	defer fsm.stateMutex.Unlock()
	fsm.state = state
}
