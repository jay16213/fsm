package fsm

// State is a thread-safe structure that represents the state of a object
// and it can be used for FSM
type State int

// NewState create a State object with current state set to initState
func NewState(initState int) *State {
	s := State(initState)
	return &s
}

// Current get the current state
func (s *State) Current() int {
	return int(*s)
}

// Is return true if the current state is equal to target
func (s *State) Is(target int) bool {
	return s.Current() == target
}

// Set current state to next
func (s *State) Set(next int) {
	*s = State(next)
}
