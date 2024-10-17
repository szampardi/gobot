//go:build !windows
// +build !windows

package firmata

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/szampardi/gobot"
)

var _ gobot.Adaptor = (*BLEAdaptor)(nil)

func initTestBLEAdaptor() *BLEAdaptor {
	a := NewBLEAdaptor("DEVICE", "123", "456")
	return a
}

func TestFirmataBLEAdaptor(t *testing.T) {
	a := initTestBLEAdaptor()
	assert.True(t, strings.HasPrefix(a.Name(), "BLEFirmata"))
}
