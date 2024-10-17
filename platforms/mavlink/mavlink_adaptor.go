package mavlink

import (
	"io"

	"go.bug.st/serial"

	"github.com/szampardi/gobot"
	common "github.com/szampardi/gobot/platforms/mavlink/common"
)

// Adaptor is a Mavlink transport adaptor.
type BaseAdaptor interface {
	gobot.Connection

	io.Writer
	ReadMAVLinkPacket() (*common.MAVLinkPacket, error)
}

// Adaptor is a Mavlink-over-serial adaptor.
type Adaptor struct {
	name    string
	port    string
	sp      io.ReadWriteCloser
	connect func(string) (io.ReadWriteCloser, error)
}

// NewAdaptor creates a new mavlink adaptor with specified port
func NewAdaptor(port string) *Adaptor {
	return &Adaptor{
		name: "Mavlink",
		port: port,
		connect: func(port string) (io.ReadWriteCloser, error) {
			return serial.Open(port, &serial.Mode{BaudRate: 57600})
		},
	}
}

func (m *Adaptor) Name() string     { return m.name }
func (m *Adaptor) SetName(n string) { m.name = n }
func (m *Adaptor) Port() string     { return m.port }

// Connect returns true if connection to device is successful
func (m *Adaptor) Connect() error {
	sp, err := m.connect(m.Port())
	if err != nil {
		return err
	}
	m.sp = sp

	return nil
}

// Finalize returns true if connection to devices is closed successfully
func (m *Adaptor) Finalize() error {
	return m.sp.Close()
}

func (m *Adaptor) ReadMAVLinkPacket() (*common.MAVLinkPacket, error) {
	return common.ReadMAVLinkPacket(m.sp)
}

func (m *Adaptor) Write(b []byte) (int, error) {
	return m.sp.Write(b)
}
