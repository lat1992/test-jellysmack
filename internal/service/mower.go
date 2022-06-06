package service

import (
	"fmt"
	"jellysmack-test/internal/model"
	"time"
)

type MowerService interface {
	// Run will run the mower
	Run() string
}

type Mower struct {
	position    model.PositionStore
	yard        YardService
	instruction string
}

func NewMowerService(posX, posY int, posO, instruction string, yard YardService) (*Mower, error) {
	if posX <= 0 || posY <= 0 {
		return nil, fmt.Errorf("the position's number are wrong")
	}
	if posO != "N" && posO != "E" && posO != "W" && posO != "S" {
		return nil, fmt.Errorf("the orientation is wrong")
	}
	for _, r := range instruction {
		if r != 'L' && r != 'R' && r != 'F' {
			return nil, fmt.Errorf("there are instruction other than 'L', 'R' and 'F'")
		}
	}
	position := model.NewPositionStore(uint(posX), uint(posY), posO)
	yard.TakePlace(uint(posX), uint(posY))

	return &Mower{
		position:    position,
		yard:        yard,
		instruction: instruction,
	}, nil
}

func (s *Mower) Run() string {
	for _, move := range s.instruction {
		if move == 'L' {
			s.turnLeft()
		} else if move == 'R' {
			s.turnRight()
		} else if move == 'F' {
			s.goForward()
		}
		time.Sleep(time.Millisecond * 100)
	}
	return fmt.Sprintf("%d %d %s", s.position.GetX(), s.position.GetY(), s.position.GetOrientation())
}

func (s *Mower) turnLeft() {
	if s.position.GetOrientation() == "N" {
		s.position.SetOrientation("W")
	} else if s.position.GetOrientation() == "E" {
		s.position.SetOrientation("N")
	} else if s.position.GetOrientation() == "W" {
		s.position.SetOrientation("S")
	} else if s.position.GetOrientation() == "S" {
		s.position.SetOrientation("E")
	}
}

func (s *Mower) turnRight() {
	if s.position.GetOrientation() == "N" {
		s.position.SetOrientation("E")
	} else if s.position.GetOrientation() == "E" {
		s.position.SetOrientation("S")
	} else if s.position.GetOrientation() == "W" {
		s.position.SetOrientation("N")
	} else if s.position.GetOrientation() == "S" {
		s.position.SetOrientation("W")
	}
}

func (s *Mower) goForward() {
	if !s.yard.CheckForwardPlace(s.position.GetX(), s.position.GetY(), s.position.GetOrientation()) {
		s.yard.ClearPlace(s.position.GetX(), s.position.GetY())
		if s.position.GetOrientation() == "N" {
			s.position.IncreaseY()
		} else if s.position.GetOrientation() == "E" {
			s.position.IncreaseX()
		} else if s.position.GetOrientation() == "W" {
			s.position.DecreaseX()
		} else if s.position.GetOrientation() == "S" {
			s.position.DecreaseY()
		}
		s.yard.TakePlace(s.position.GetX(), s.position.GetY())
	}
}
