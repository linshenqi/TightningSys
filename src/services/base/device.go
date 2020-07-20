package base

const (
	StatusOnline  = "online"
	StatusOffline = "offline"
)

type IDevice interface {
	Start() error
	Stop() error
	Status() string
	DeviceType() string
	Children() map[string]IDevice
	Config() interface{}
	SerialNumber() string
}

type BaseDevice struct {
	IDevice
}

func (s *BaseDevice) Start() error {
	return nil
}

func (s *BaseDevice) Stop() error {
	return nil
}

func (s *BaseDevice) Status() string {
	return StatusOffline
}

func (s *BaseDevice) DeviceType() string {
	return ""
}

func (s *BaseDevice) Children() map[string]IDevice {
	return nil
}

func (s *BaseDevice) Config() interface{} {
	return nil
}

func (s *BaseDevice) SerialNumber() string {
	return ""
}
