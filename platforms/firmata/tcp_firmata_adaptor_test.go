//go:build !windows
// +build !windows

package firmata

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/szampardi/gobot"
)

var _ gobot.Adaptor = (*TCPAdaptor)(nil)

func initTestTCPAdaptor() *TCPAdaptor {
	a := NewTCPAdaptor("localhost:4567")
	return a
}

func TestFirmataTCPAdaptor(t *testing.T) {
	a := initTestTCPAdaptor()
	assert.True(t, strings.HasPrefix(a.Name(), "TCPFirmata"))
}
