package microbit

import (
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/szampardi/gobot"
)

var _ gobot.Driver = (*ButtonDriver)(nil)

func initTestButtonDriver() *ButtonDriver {
	d := NewButtonDriver(NewBleTestAdaptor())
	return d
}

func TestButtonDriver(t *testing.T) {
	d := initTestButtonDriver()
	assert.True(t, strings.HasPrefix(d.Name(), "Microbit Button"))
	d.SetName("NewName")
	assert.Equal(t, "NewName", d.Name())
}

func TestButtonDriverStartAndHalt(t *testing.T) {
	d := initTestButtonDriver()
	require.NoError(t, d.Start())
	require.NoError(t, d.Halt())
}

func TestButtonDriverReadData(t *testing.T) {
	sem := make(chan bool)
	a := NewBleTestAdaptor()
	d := NewButtonDriver(a)
	_ = d.Start()
	_ = d.On(ButtonB, func(data interface{}) {
		sem <- true
	})

	a.TestReceiveNotification([]byte{1}, nil)

	select {
	case <-sem:
	case <-time.After(100 * time.Millisecond):
		t.Errorf("Microbit Event \"ButtonB\" was not published")
	}
}
