package model

import (
	"testing"

	"github.com/go-playground/assert"
)

func Test_GetAll(t *testing.T) {
	tests := []struct {
		name  string
		x     uint
		y     uint
		o     string
		wantX uint
		wantY uint
		wantO string
	}{
		{
			name:  "normal use case",
			x:     3,
			y:     3,
			o:     "N",
			wantX: 3,
			wantY: 3,
			wantO: "N",
		},
		{
			name: "should return nothing",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewPositionStore(tt.x, tt.y, tt.o)
			x := s.GetX()
			y := s.GetY()
			o := s.GetOrientation()

			assert.Equal(t, tt.wantX, x)
			assert.Equal(t, tt.wantY, y)
			assert.Equal(t, tt.wantO, o)
		})
	}
}

func Test_SetOrientation(t *testing.T) {
	tests := []struct {
		name  string
		o     string
		wantO string
	}{
		{
			name:  "normal use case",
			o:     "N",
			wantO: "N",
		},
		{
			name:  "another character than N S W E",
			o:     "A",
			wantO: "W",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewPositionStore(3, 3, "W")
			s.SetOrientation(tt.o)

			assert.Equal(t, tt.wantO, s.o)
		})
	}
}
