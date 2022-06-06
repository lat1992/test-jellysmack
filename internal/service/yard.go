package service

import (
	"fmt"
	"jellysmack-test/internal/model"
)

type YardService interface {
	// TakePlace will set true into coordinates
	TakePlace(x, y uint)

	// ClearPlace will set false into coordinates
	ClearPlace(x, y uint)

	// CheckForwardPlace will return boolean after check if there are free
	CheckForwardPlace(x, y uint, o string) bool
}

type Yard struct {
	coordinates model.MapStore
}

func NewYardService(x, y int) (*Yard, error) {
	if x <= 0 || y <= 0 {
		return nil, fmt.Errorf("the limit's number are wrong")
	}

	m := model.NewMapStore(uint(x), uint(y))
	return &Yard{
		coordinates: m,
	}, nil
}

func (s *Yard) TakePlace(x, y uint) {
	s.coordinates.UsePlace(x, y)
}

func (s *Yard) ClearPlace(x, y uint) {
	s.coordinates.RemovePlace(x, y)
}

func (s *Yard) CheckForwardPlace(x, y uint, o string) bool {
	if o == "N" && y < s.coordinates.GetLimitY() {
		return s.coordinates.GetPlace(x, y+1)
	} else if o == "E" && x < s.coordinates.GetLimitX() {
		return s.coordinates.GetPlace(x+1, y)
	} else if o == "W" && x > 0 {
		return s.coordinates.GetPlace(x-1, y)
	} else if o == "S" && y > 0 {
		return s.coordinates.GetPlace(x, y-1)
	}
	return true
}
