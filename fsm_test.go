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

func open(event *fsm.Event) {
	fmt.Printf("open event callback: %+v\n", event)
}

func TestInitFSM(t *testing.T) {
	fsm.NewFSM(Closed,
		[]fsm.Transition{
			{
				Event: Open,
				From:  Closed,
				To:    Opened,
			},
			{
				Event: Close,
				From:  Opened,
				To:    Closed,
			},
		}, map[fsm.State]fsm.CallbackFunc{
			Opened: open,
		})
}
