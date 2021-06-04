package fsm

import "sync/atomic"

// State is a thread-safe structure that represents the state of a object
// and it can be used for FSM
type State int32

// NewState create a State object with current state set to initState
func NewState(initState int) *State {
	s := State(initState)
	return &s
}

// Current get the current state
func (s *State) Current() int {
	return int(atomic.LoadInt32((*int32)(s)))
}

// Is return true if the current state is equal to target
func (s *State) Is(target int) bool {
	return s.Current() == target
}

// Set current state to next
func (s *State) Set(next int) {
	atomic.StoreInt32((*int32)(s), int32(next))
}
