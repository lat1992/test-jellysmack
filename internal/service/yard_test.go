package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_CheckForwardPlace(t *testing.T) {
	tests := []struct {
		name string
		x    uint
		y    uint
		o    string
		want bool
	}{
		{
			name: "normal use case",
			x:    1,
			y:    2,
			o:    "N",
			want: false,
		},
		{
			name: "cannot move",
			x:    0,
			y:    0,
			o:    "S",
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s, err := NewYardService(5, 5)
			assert.Nil(t, err)
			result := s.CheckForwardPlace(tt.x, tt.y, tt.o)
			assert.Equal(t, tt.want, result)
		})
	}
}

func Test_TakePlace(t *testing.T) {
	tests := []struct {
		name   string
		x      uint
		y      uint
		o      string
		placeX uint
		placeY uint
		want   bool
	}{
		{
			name:   "normal use case",
			x:      0,
			y:      1,
			o:      "S",
			placeX: 0,
			placeY: 0,
			want:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s, err := NewYardService(5, 5)
			assert.Nil(t, err)
			s.TakePlace(tt.placeX, tt.placeY)
			result := s.CheckForwardPlace(tt.x, tt.y, tt.o)
			assert.Equal(t, tt.want, result)
		})
	}
}

func Test_ClearPlace(t *testing.T) {
	tests := []struct {
		name   string
		x      uint
		y      uint
		o      string
		placeX uint
		placeY uint
		want   bool
	}{
		{
			name:   "normal use case",
			x:      0,
			y:      1,
			o:      "S",
			placeX: 0,
			placeY: 0,
			want:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s, err := NewYardService(5, 5)
			assert.Nil(t, err)
			s.TakePlace(tt.placeX, tt.placeY)
			s.ClearPlace(tt.placeX, tt.placeY)
			result := s.CheckForwardPlace(tt.x, tt.y, tt.o)
			assert.Equal(t, tt.want, result)
		})
	}
}
