package bebop

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/szampardi/gobot"
)

var _ gobot.Driver = (*Driver)(nil)

func TestBebopDriverName(t *testing.T) {
	a := initTestBebopAdaptor()
	d := NewDriver(a)
	assert.True(t, strings.HasPrefix(d.Name(), "Bebop"))
	d.SetName("NewName")
	assert.Equal(t, "NewName", d.Name())
}
