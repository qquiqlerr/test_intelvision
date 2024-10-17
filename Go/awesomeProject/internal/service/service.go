package service

import "errors"

var ErrBadValue = errors.New("value must be 0 or 1")

type PortSystem interface {
	Read(id int) (int, error)
	Write(id, value int) error
}

type PortServiceImpl struct {
	Port PortSystem
}

func (p *PortServiceImpl) Read(id int) (int, error) {
	return p.Port.Read(id)
}

func (p *PortServiceImpl) Write(id, value int) error {
	if value < 0 || value > 1 {
		return ErrBadValue
	}
	return p.Port.Write(id, value)
}

func NewPortService(portSystem PortSystem) *PortServiceImpl {
	return &PortServiceImpl{
		Port: portSystem,
	}
}
