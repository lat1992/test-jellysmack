package model

type MapStore interface {
	// UsePlace will set the boolean to true for the target coordinate
	UsePlace(x, y uint)

	// RemovePlace will set the boolean to false for the target coordinate
	RemovePlace(x, y uint)

	// GetPlace will get the boolean to true for the target coordinate
	GetPlace(x, y uint) bool

	// GetLimitX will return the limit of X
	GetLimitX() uint

	// GetLimitY will return the limit of Y
	GetLimitY() uint
}

type Map struct {
	places [][]bool
	x      uint
	y      uint
}

func NewMapStore(x, y uint) *Map {
	places := make([][]bool, y+1)
	for i := uint(0); i <= y; i++ {
		places[i] = make([]bool, x+1)
	}
	return &Map{
		places: places,
		x:      x,
		y:      y,
	}
}

func (s *Map) UsePlace(x, y uint) {
	if x > s.x || y > s.y {
		return
	}
	s.places[y][x] = true
}

func (s *Map) RemovePlace(x, y uint) {
	if x > s.x || y > s.y {
		return
	}
	s.places[y][x] = false
}

func (s *Map) GetPlace(x, y uint) bool {
	if x > s.x || y > s.y {
		return true
	}
	return s.places[y][x]
}

func (s *Map) GetLimitX() uint {
	return s.x
}

func (s *Map) GetLimitY() uint {
	return s.y
}
