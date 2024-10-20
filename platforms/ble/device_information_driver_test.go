package ble

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/szampardi/gobot"
)

var _ gobot.Driver = (*DeviceInformationDriver)(nil)

func initTestDeviceInformationDriver() *DeviceInformationDriver {
	d := NewDeviceInformationDriver(NewBleTestAdaptor())
	return d
}

func TestDeviceInformationDriver(t *testing.T) {
	d := initTestDeviceInformationDriver()
	assert.True(t, strings.HasPrefix(d.Name(), "DeviceInformation"))
	d.SetName("NewName")
	assert.Equal(t, "NewName", d.Name())
}

func TestDeviceInformationDriverStartAndHalt(t *testing.T) {
	d := initTestDeviceInformationDriver()
	require.NoError(t, d.Start())
	require.NoError(t, d.Halt())
}

func TestDeviceInformationDriverGetModelNumber(t *testing.T) {
	a := NewBleTestAdaptor()
	d := NewDeviceInformationDriver(a)
	a.TestReadCharacteristic(func(cUUID string) ([]byte, error) {
		return []byte("TestDevice"), nil
	})

	assert.Equal(t, "TestDevice", d.GetModelNumber())
}

func TestDeviceInformationDriverGetFirmwareRevision(t *testing.T) {
	a := NewBleTestAdaptor()
	d := NewDeviceInformationDriver(a)
	a.TestReadCharacteristic(func(cUUID string) ([]byte, error) {
		return []byte("TestDevice"), nil
	})

	assert.Equal(t, "TestDevice", d.GetFirmwareRevision())
}

func TestDeviceInformationDriverGetHardwareRevision(t *testing.T) {
	a := NewBleTestAdaptor()
	d := NewDeviceInformationDriver(a)
	a.TestReadCharacteristic(func(cUUID string) ([]byte, error) {
		return []byte("TestDevice"), nil
	})

	assert.Equal(t, "TestDevice", d.GetHardwareRevision())
}

func TestDeviceInformationDriverGetManufacturerName(t *testing.T) {
	a := NewBleTestAdaptor()
	d := NewDeviceInformationDriver(a)
	a.TestReadCharacteristic(func(cUUID string) ([]byte, error) {
		return []byte("TestDevice"), nil
	})

	assert.Equal(t, "TestDevice", d.GetManufacturerName())
}

func TestDeviceInformationDriverGetPnPId(t *testing.T) {
	a := NewBleTestAdaptor()
	d := NewDeviceInformationDriver(a)
	a.TestReadCharacteristic(func(cUUID string) ([]byte, error) {
		return []byte("TestDevice"), nil
	})

	assert.Equal(t, "TestDevice", d.GetPnPId())
}
