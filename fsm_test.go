package fsm_test

import (
	"fmt"
	"free5gc/lib/fsm"
	"testing"
)

const (
	Opened fsm.State = "Opened"
	Closed fsm.State = "Closed"
)

const (
	Open  string = "Open"
	Close string = "Close"
)

func TestInitFSM(t *testing.T) {
	f, err := fsm.NewFSM(Closed, fsm.Callbacks{
		Opened: func(sm *fsm.FSM, event fsm.Event, args fsm.Args) {
			fmt.Printf("event [%+v] at state [%+v]\n", event, sm.Current())
		},
	})

	if err != nil {
		t.Error(err)
	}

	if err = f.Transition(Opened, nil); err != nil {
		t.Error(err)
	}
}
