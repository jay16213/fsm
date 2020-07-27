package fsm_test

import (
	"fmt"
	"free5gc/lib/fsm"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	Opened fsm.State = "Opened"
	Closed fsm.State = "Closed"
)

const (
	Open  string = "Open"
	Close string = "Close"
)

func TestFSM(t *testing.T) {
	f, err := fsm.NewFSM(Closed, fsm.Callbacks{
		Opened: func(sm *fsm.FSM, event fsm.Event, args fsm.Args) {
			fmt.Printf("event [%+v] at state [%+v]\n", event, sm.Current())
		},
	})

	assert.Nil(t, err, "NewFSM() failed")

	assert.Equal(t, Closed, f.Current(), "Current() failed")
	assert.True(t, f.IsCurrent(Closed), "IsCurrent() failed")

	assert.Nil(t, f.Transition(Opened, nil), "Transition() failed")
	assert.Nil(t, f.SendEvent(fsm.Event("Test Event"), nil), "SendEvent() failed")

	assert.Equal(t, Opened, f.Current(), "Current() failed")
	assert.True(t, f.IsCurrent(Opened), "IsCurrent() failed")

	f.SetState(Closed)
	assert.Equal(t, Closed, f.Current(), "SetState() failed")
}
