package fsm

import "sync"

type State struct {
	// current is the current state of the FSM
	current StateType
	// stateMutex ensure that all operation to current is thread-safe
	stateMutex sync.RWMutex
}

// NewState create a State object with current state equal to initState
func NewState(initState StateType) *State {
	s := &State{current: initState}
	return s
}

// Current get the current state of fsm
func (state *State) Current() StateType {
	state.stateMutex.RLock()
	defer state.stateMutex.RUnlock()
	return state.current
}

// Is return true if the current state of fsm is equal to state
func (state *State) Is(target StateType) bool {
	state.stateMutex.RLock()
	defer state.stateMutex.RUnlock()
	return state.current == target
}

// Set current state to next
func (state *State) Set(next StateType) {
	state.stateMutex.Lock()
	defer state.stateMutex.Unlock()
	state.current = next
}
