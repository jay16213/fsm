package fsm_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/free5gc/fsm"
)

const (
	Opened int = 0
	Closed int = 1
)

const (
	Open  fsm.EventType = "Open"
	Close fsm.EventType = "Close"
)

func TestState(t *testing.T) {
	s := fsm.NewState(Closed)

	assert.Equal(t, Closed, s.Current(), "Current() failed")
	assert.True(t, s.Is(Closed), "Is() failed")

	s.Set(Opened)

	assert.Equal(t, Opened, s.Current(), "Current() failed")
	assert.True(t, s.Is(Opened), "Is() failed")
}
