package model

import (
	"testing"

	"github.com/go-playground/assert"
)

func Test_UsePlace(t *testing.T) {
	tests := []struct {
		name string
		limX uint
		limY uint
		x    uint
		y    uint
		want [][]bool
	}{
		{
			name: "normal use case",
			limY: 2,
			limX: 2,
			x:    1,
			y:    2,
			want: [][]bool{
				{false, false, false},
				{false, false, false},
				{false, true, false},
			},
		},
		{
			name: "should not be place",
			limY: 2,
			limX: 2,
			x:    1,
			y:    5,
			want: [][]bool{
				{false, false, false},
				{false, false, false},
				{false, false, false},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewMapStore(tt.limX, tt.limY)
			s.UsePlace(tt.x, tt.y)
			assert.Equal(t, tt.want, s.places)
		})
	}
}

func Test_RemovePlace(t *testing.T) {
	tests := []struct {
		name string
		limX uint
		limY uint
		x    uint
		y    uint
		want [][]bool
	}{
		{
			name: "normal use case",
			limY: 2,
			limX: 2,
			x:    1,
			y:    2,
			want: [][]bool{
				{false, false, false},
				{false, false, false},
				{false, false, false},
			},
		},
		{
			name: "should not remove any place",
			limY: 2,
			limX: 2,
			x:    1,
			y:    5,
			want: [][]bool{
				{false, false, false},
				{false, false, false},
				{false, true, false},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewMapStore(tt.limX, tt.limY)
			s.places[2][1] = true
			s.RemovePlace(tt.x, tt.y)
			assert.Equal(t, tt.want, s.places)
		})
	}
}

func Test_GetPlace(t *testing.T) {
	tests := []struct {
		name string
		limX uint
		limY uint
		x    uint
		y    uint
		want bool
	}{
		{
			name: "normal use case",
			limY: 2,
			limX: 2,
			x:    1,
			y:    2,
			want: true,
		},
		{
			name: "should return false",
			limY: 2,
			limX: 2,
			x:    1,
			y:    5,
			want: true,
		},
		{
			name: "out of limit",
			limY: 2,
			limX: 2,
			x:    1,
			y:    5,
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewMapStore(tt.limX, tt.limY)
			s.places[2][1] = true
			result := s.GetPlace(tt.x, tt.y)
			assert.Equal(t, tt.want, result)
		})
	}
}

func Test_GetLimit(t *testing.T) {
	tests := []struct {
		name  string
		limX  uint
		limY  uint
		wantX uint
		wantY uint
	}{
		{
			name:  "normal use case",
			limY:  2,
			limX:  2,
			wantY: 2,
			wantX: 2,
		},
		{
			name: "nothing is set",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewMapStore(tt.limX, tt.limY)
			y := s.GetLimitY()
			x := s.GetLimitX()
			assert.Equal(t, tt.wantX, x)
			assert.Equal(t, tt.wantY, y)
		})
	}
}
