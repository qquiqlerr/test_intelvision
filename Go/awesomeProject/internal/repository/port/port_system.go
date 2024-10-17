package port

import (
	"errors"
	"math/rand"
)

var errPortNotFound = errors.New("port not found")

type PortSystemImpl struct {
	IN  []int
	OUT []int
}

func (p *PortSystemImpl) Read(id int) (int, error) {
	if id >= len(p.IN) || id < 0 {
		return 0, errPortNotFound
	}
	return p.IN[id], nil
}

func (p *PortSystemImpl) Write(id, value int) error {
	if id >= len(p.OUT) || id < 0 {
		return errPortNotFound
	}
	p.OUT[id] = value
	return nil
}

func NewPortSystem(INSize, OUTSize int) *PortSystemImpl {
	IN := make([]int, INSize)
	OUT := make([]int, OUTSize)
	for i := 0; i < INSize; i++ {
		IN[i] = rand.Intn(2)
	}
	return &PortSystemImpl{
		IN:  IN,
		OUT: OUT,
	}
}
