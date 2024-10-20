package i2c

// VEML7700 I2C address
const (
	defaultVEML7700Address = 0x10
)

// Gain settings
const (
	ALS_GAIN_1   = 0x0
	ALS_GAIN_2   = 0x1
	ALS_GAIN_1_4 = 0x3
	ALS_GAIN_1_8 = 0x2
)

// Integration time settings
const (
	ALS_25MS  = 0xC
	ALS_50MS  = 0x8
	ALS_100MS = 0x0
	ALS_200MS = 0x1
	ALS_400MS = 0x2
	ALS_800MS = 0x3
)

// VEML7700Driver is a driver for the VEML7700 ambient light sensor
type VEML7700Driver struct {
	*Driver
	gain            byte
	integrationTime byte
}

// NewVEML7700Driver creates a new driver with the specified I2C interface
func NewVEML7700Driver(c Connector, options ...func(Config)) *VEML7700Driver {
	v := &VEML7700Driver{
		Driver: NewDriver(c, "VEML7700", defaultVEML7700Address),
	}
	v.afterStart = v.initialize

	for _, option := range options {
		option(v)
	}

	return v
}

// initialize configures the VEML7700 upon starting
func (v *VEML7700Driver) initialize() error {
	if err := v.setGain(ALS_GAIN_1_8); err != nil {
		return err
	}
	if err := v.setIntegrationTime(ALS_100MS); err != nil {
		return err
	}
	return v.setShutdown(false)
}

// ReadRawLight reads raw ambient light data from the sensor
func (v *VEML7700Driver) ReadRawLight() (uint16, error) {
	buf := make([]byte, 2)
	if err := v.connection.ReadBlockData(0x04, buf); err != nil {
		return 0, err
	}
	return uint16(buf[0]) | (uint16(buf[1]) << 8), nil
}

// ReadRawWhite reads raw white light data from the sensor
func (v *VEML7700Driver) ReadRawWhite() (uint16, error) {
	buf := make([]byte, 2)
	if err := v.connection.ReadBlockData(0x05, buf); err != nil {
		return 0, err
	}
	return uint16(buf[0]) | (uint16(buf[1]) << 8), nil
}

// setShutdown sets the shutdown state of the sensor
func (v *VEML7700Driver) setShutdown(state bool) error {
	var shutdown byte
	if state {
		shutdown = 1
	} else {
		shutdown = 0
	}
	return v.connection.WriteBlockData(0x00, []byte{shutdown})
}

// setGain sets the gain for the ambient light sensor
func (v *VEML7700Driver) setGain(gain byte) error {
	v.gain = gain
	return v.connection.WriteBlockData(0x00, []byte{gain})
}

// setIntegrationTime sets the integration time for the ambient light sensor
func (v *VEML7700Driver) setIntegrationTime(integrationTime byte) error {
	v.integrationTime = integrationTime
	return v.connection.WriteBlockData(0x00, []byte{integrationTime})
}

// GainValue returns the current gain setting as a float
func (v *VEML7700Driver) GainValue() float64 {
	switch v.gain {
	case ALS_GAIN_2:
		return 2.0
	case ALS_GAIN_1:
		return 1.0
	case ALS_GAIN_1_4:
		return 0.25
	case ALS_GAIN_1_8:
		return 0.125
	default:
		return 1.0
	}
}

// IntegrationTimeValue returns the current integration time in milliseconds
func (v *VEML7700Driver) IntegrationTimeValue() int {
	switch v.integrationTime {
	case ALS_25MS:
		return 25
	case ALS_50MS:
		return 50
	case ALS_100MS:
		return 100
	case ALS_200MS:
		return 200
	case ALS_400MS:
		return 400
	case ALS_800MS:
		return 800
	default:
		return 100
	}
}

// Lux calculates the light value in lux
func (v *VEML7700Driver) Lux() (float64, error) {
	light, err := v.ReadRawLight()
	if err != nil {
		return 0, err
	}
	resolution := v.resolution()
	return resolution * float64(light), nil
}

// resolution calculates the resolution necessary to calculate lux
func (v *VEML7700Driver) resolution() float64 {
	resolutionAtMax := 0.0036
	gainMax := 2.0
	integrationTimeMax := 800.0
	integrationTimeValue := float64(v.IntegrationTimeValue())
	gainValue := v.GainValue()

	if gainValue == gainMax && integrationTimeValue == integrationTimeMax {
		return resolutionAtMax
	}
	return resolutionAtMax * (integrationTimeMax / integrationTimeValue) * (gainMax / gainValue)
}
