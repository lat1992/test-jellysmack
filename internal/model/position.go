package model

type PositionStore interface {
	// GetOrientation will return orientation
	GetOrientation() string

	// SetOrientation will set orientation to o
	SetOrientation(o string)

	// GetX will return X
	GetX() uint

	// IncreaseX will increase X
	IncreaseX()

	// DecreaseX will decrease X
	DecreaseX()

	// GetY will return Y
	GetY() uint

	// IncreaseY will increase Y
	IncreaseY()

	// DecreaseY will decrease Y
	DecreaseY()
}

type Position struct {
	x uint
	y uint
	o string
}

func NewPositionStore(x, y uint, o string) *Position {
	return &Position{
		x: x,
		y: y,
		o: o,
	}
}

func (s *Position) GetOrientation() string {
	return s.o
}

func (s *Position) SetOrientation(o string) {
	if o != "N" && o != "E" && o != "W" && o != "S" {
		return
	}
	s.o = o
}

func (s *Position) GetX() uint {
	return s.x
}

func (s *Position) IncreaseX() {
	s.x++
}

func (s *Position) DecreaseX() {
	s.x--
}

func (s *Position) GetY() uint {
	return s.y
}

func (s *Position) IncreaseY() {
	s.y++
}

func (s *Position) DecreaseY() {
	s.y--
}
