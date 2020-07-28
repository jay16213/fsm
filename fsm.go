package fsm

import (
	"fmt"
)

type StateType string
type EventType string
type ArgsType map[string]interface{}

type Callback func(*State, EventType, ArgsType)
type Callbacks map[StateType]Callback

// Transition defines a transition
// that a Event is triggered at From state,
// and transfer to To state after the Event
type Transition struct {
	Event EventType
	From  StateType
	To    StateType
}

type Transitions []Transition

type eventKey struct {
	Event EventType
	From  StateType
}

// Entry/Exit event are defined by fsm package
const (
	EntryEvent EventType = "Entry event"
	ExitEvent  EventType = "Exit event"
)

type FSM struct {
	// transitions stores one transition for each event
	transitions map[eventKey]Transition
	// callbacks stores one callback function for one state
	callbacks map[StateType]Callback
}

// NewFSM create a new FSM object then registers transitions and callbacks to it
func NewFSM(transitions Transitions, callbacks Callbacks) (*FSM, error) {
	fsm := &FSM{
		transitions: make(map[eventKey]Transition),
		callbacks:   make(map[StateType]Callback),
	}

	allStates := make(map[StateType]bool)

	for _, transition := range transitions {
		key := eventKey{
			Event: transition.Event,
			From:  transition.From,
		}
		if _, ok := fsm.transitions[key]; ok {
			return nil, fmt.Errorf("Duplicate transition: %+v", transition)
		} else {
			fsm.transitions[key] = transition
			allStates[transition.From] = true
			allStates[transition.To] = true
		}
	}

	for state, callback := range callbacks {
		if _, ok := allStates[state]; !ok {
			return nil, fmt.Errorf("Unknown state: %+v", state)
		} else {
			fsm.callbacks[state] = callback
		}
	}
	return fsm, nil
}

// SendEvent triggers a callback with an event, and do transition after callback if need
// There are 3 types of callback:
//  - on exit callback: call when fsm leave one state, with ExitEvent event
//  - event callback: call when user trigger a user-defined event
//  - on entry callback: call when fsm enter one state, with EntryEvent event
func (fsm *FSM) SendEvent(state *State, event EventType, args ArgsType) error {
	key := eventKey{
		From:  state.Current(),
		Event: event,
	}

	if trans, ok := fsm.transitions[key]; ok {
		// exit callback
		if trans.From != trans.To {
			fsm.callbacks[trans.From](state, ExitEvent, args)
		}

		// event callback
		fsm.callbacks[trans.From](state, event, args)

		// entry callback
		if trans.From != trans.To {
			state.Set(trans.To)
			fsm.callbacks[trans.To](state, EntryEvent, args)
		}
		return nil
	} else {
		return fmt.Errorf("Unknown transition[From: %s, Event: %s]", state.Current(), event)
	}
}
